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

func getTaskByID(c *gin.Context) (taskID int64, ownerID int64, err error) {
	taskId_, err := strconv.Atoi(c.Param("id"))
	if err != nil || taskId_ <= 0 {
		return 0, 0, errors.New("invalid task id")
	}

	taskID = int64(taskId_)
	taskCreds := protos.TaskCreds{Id: taskID}
	resp, err := grpc.GRPCTaskServiceClient.GetTask(context.Background(), &taskCreds)
	if err != nil {
		return 0, 0, errors.New("invalid task creds")
	}
	ownerID = resp.OwnerId

	return
}

func getUsernameByUserID(userID int64) (username string, err error) {
	var dbUser schemas.UserData
	db.DB.First(&dbUser, uint(userID))
	if dbUser.ID == 0 {
		return "", errors.New("user not found")
	} else {
		return dbUser.Login, nil
	}
}

func ViewOrLikeTask(c *gin.Context) {
	taskId, ownerID, err := getTaskByID(c)
	if err != nil {
		c.Status(http.StatusBadRequest)
		log.Println(err.Error())
		return
	}

	userID, _ := c.Get("user")

	ownerUsername, err := getUsernameByUserID(ownerID)
	if err != nil {
		log.Println(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	event := schemas.Event{
		TaskID:        taskId,
		UserID:        int64(userID.(uint)),
		OwnerUsername: ownerUsername,
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

	err = message_broker.SendEventToBroker(event)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	} else {
		c.Status(http.StatusOK)
	}
}

func GetSpecificTaskStatistics(c *gin.Context) {
	taskId, _, err := getTaskByID(c)
	if taskId == 0 {
		if err != nil {
			c.Status(http.StatusBadRequest)
			log.Println(err.Error())
			return
		}
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
	orderBy := c.Query("order_by")
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
		log.Println(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func GetTopUsers(c *gin.Context) {
	orderBy, err := getOrderBy(c)
	if err != nil {
		return
	}
	req := protos.TopRequest{Type: orderBy}

	resp, err := grpc.GRPCStatisticsServiceClient.GetTopUsers(context.Background(), &req)
	if err != nil {
		log.Println(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, resp)
}
