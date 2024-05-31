package controllers

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"userservice/db"
	"userservice/schemas"

	"github.com/gin-gonic/gin"
	"userservice/grpc"
	"userservice/protos"
)

func CreateTask(c *gin.Context) {
	var task protos.Task
	if err := c.BindJSON(&task); err != nil {
		log.Println("ERROR: invalid request body")
		c.Status(http.StatusBadRequest)
		return
	}

	taskOwner, _ := c.Get("user")
	task.OwnerId = int64(taskOwner.(uint))
	resp, err := grpc.GRPCTaskServiceClient.CreateTask(context.Background(), &task)
	returnResponse(c, resp, err, http.StatusCreated)
}

func UpdateTask(c *gin.Context) {
	var task protos.Task
	if err := c.BindJSON(&task); err != nil {
		log.Println("ERROR: invalid request body")
		c.Status(http.StatusBadRequest)
		return
	}

	taskOwner, _ := c.Get("user")
	task.OwnerId = int64(taskOwner.(uint))
	id, _ := strconv.Atoi(c.Param("id"))
	task.Id = int64(id)
	resp, err := grpc.GRPCTaskServiceClient.UpdateTask(context.Background(), &task)
	returnResponse(c, resp, err, http.StatusOK)
}

func DeleteTask(c *gin.Context) {
	var taskCreds protos.TaskCreds
	taskOwner, _ := c.Get("user")
	taskCreds.OwnerId = int64(taskOwner.(uint))
	id, _ := strconv.Atoi(c.Param("id"))
	taskCreds.Id = int64(id)
	resp, err := grpc.GRPCTaskServiceClient.DeleteTask(context.Background(), &taskCreds)
	returnResponse(c, resp, err, http.StatusOK)
}

func GetTask(c *gin.Context) {
	var taskCreds protos.TaskCreds
	taskOwner, _ := c.Get("user")
	taskCreds.OwnerId = int64(taskOwner.(uint))
	id, _ := strconv.Atoi(c.Param("id"))
	taskCreds.Id = int64(id)
	resp, err := grpc.GRPCTaskServiceClient.GetTask(context.Background(), &taskCreds)
	returnResponse(c, resp, err, http.StatusOK)
}

func GetTasksByPage(c *gin.Context) {
	var pageReq protos.PageRequest

	var dbUser schemas.UserData
	db.DB.First(&dbUser, "login = ?", c.Query("user"))
	if dbUser.ID == 0 {
		log.Println("ERROR: user not found")
		c.Status(http.StatusBadRequest)
		return
	}
	pageReq.OwnerId = int64(dbUser.ID)

	size, err := strconv.Atoi(c.Query("size"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	pageReq.Size = int64(size)

	number, err := strconv.Atoi(c.Query("number"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	pageReq.Number = int64(number)

	resp, err := grpc.GRPCTaskServiceClient.GetTasksByPage(context.Background(), &pageReq)
	returnResponse(c, resp, err, http.StatusOK)
}
