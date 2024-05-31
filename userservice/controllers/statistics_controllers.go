package controllers

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"userservice/db"

	"github.com/gin-gonic/gin"

	"userservice/grpc"
	"userservice/message_broker"
	"userservice/protos"
	"userservice/schemas"
)

func getTaskById(c *gin.Context) (taskId int64) {
	taskId_, err := strconv.Atoi(c.Query("task_id"))
	if err != nil || taskId_ <= 0 {
		log.Println("ERROR: invalid task id")
		c.Status(http.StatusBadRequest)
		return 0
	}

	taskId = int64(taskId_)
	taskCreds := protos.TaskCreds{Id: taskId}
	_, err = grpc.GRPCTaskServiceClient.GetTask(context.Background(), &taskCreds)
	if err != nil {
		log.Println("ERROR: invalid task id")
		c.Status(http.StatusBadRequest)
		return 0
	}

	return taskId
}

func ViewOrLikeTask(c *gin.Context) {
	taskId := getTaskById(c)
	if taskId == 0 {
		return
	}

	currentUser, _ := c.Get("user")
	var dbUser schemas.UserData
	db.DB.First(&dbUser, currentUser.(uint))

	event := schemas.Event{
		TaskID:   taskId,
		Username: dbUser.Login,
	}
	eventType := c.Query("event_type")
	if eventType == "view" {
		event.EventType = schemas.View
	} else if eventType == "like" {
		event.EventType = schemas.Like
	} else {
		log.Println("ERROR: invalid event_type")
		c.Status(http.StatusBadRequest)
		return
	}

	err := message_broker.SendEventToBroker(event)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	} else {
		c.Status(http.StatusOK)
	}
}

func GetSpecificTaskStatistics(c *gin.Context) {
	taskId := getTaskById(c)
	if taskId == 0 {
		return
	}
	specificTaskReq := protos.SpecificTaskRequest{TaskID: taskId}

	statisticsServiceResp, err := grpc.GRPCStatisticsServiceClient.GetSpecificTaskStatistics(context.Background(), &specificTaskReq)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusOK, statisticsServiceResp)
	}
}
