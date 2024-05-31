package controllers

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"userservice/db"
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
	req := protos.SpecificTaskRequest{TaskID: taskId}

	resp, err := grpc.GRPCStatisticsServiceClient.GetSpecificTaskStatistics(context.Background(), &req)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

func getOrderBy(c *gin.Context) (protos.StatisticsType, error) {
	orderBy := c.Query("")
	if len(orderBy) == 0 {
		orderBy = "view"
	}
	orderBy = strings.ToLower(orderBy)
	if orderBy == "view" {
		return protos.StatisticsType_View, nil
	} else if orderBy == "like" {
		return protos.StatisticsType_Like, nil
	} else {
		log.Println("ERROR: invalid order by type")
		c.Status(http.StatusBadRequest)
		return protos.StatisticsType_View, errors.New("invalid order by type")
	}
}

func GetTopTasks(c *gin.Context) {
	orderBy, err := getOrderBy(c)
	if err != nil {
		return
	}
	req := protos.TopRequest{Type: orderBy}

	resp, err := grpc.GRPCStatisticsServiceClient.GetTopTasks(context.Background(), &req)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

func GetTopUsers(c *gin.Context) {
	orderBy, err := getOrderBy(c)
	if err != nil {
		return
	}
	req := protos.TopRequest{Type: orderBy}

	resp, err := grpc.GRPCStatisticsServiceClient.GetTopUsers(context.Background(), &req)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}
