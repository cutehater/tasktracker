package controllers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"userservice/grpc"
	"userservice/message_broker"
	"userservice/protos"
	"userservice/schemas"
)

func getTaskAndUser(c *gin.Context) (taskId int64, userId int64) {
	taskId_, err := strconv.Atoi(c.Query("task_id"))
	if err != nil || taskId_ <= 0 {
		log.Println("ERROR: invalid task id")
		c.Status(http.StatusBadRequest)
		return 0, 0
	}

	taskId = int64(taskId_)
	taskCreds := protos.TaskCreds{Id: taskId}
	_, err = grpc.GRPCClient.GetTask(context.Background(), &taskCreds)
	if err != nil {
		log.Println("ERROR: invalid task id")
		c.Status(http.StatusBadRequest)
		return 0, 0
	}

	userId_, _ := c.Get("user")
	userId = int64(userId_.(uint))

	return
}

func ViewTask(c *gin.Context) {
	taskId, userId := getTaskAndUser(c)
	if taskId == 0 {
		return
	}

	event := schemas.Event{
		TaskID:    taskId,
		UserID:    userId,
		EventType: schemas.View,
	}

	err := message_broker.SendEventToBroker(event)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	} else {
		c.Status(http.StatusOK)
	}
}

func LikeTask(c *gin.Context) {
	taskId, userId := getTaskAndUser(c)
	if taskId == 0 {
		return
	}

	event := schemas.Event{
		TaskID:    taskId,
		UserID:    userId,
		EventType: schemas.Like,
	}

	err := message_broker.SendEventToBroker(event)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	} else {
		c.Status(http.StatusOK)
	}
}
