package controllers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"userservice/grpc"
	"userservice/protos"
)

func returnResponse(c *gin.Context, resp any, err error, okStatus int) {
	if err != nil {
		if status.Code(err) == codes.PermissionDenied {
			log.Println("ERROR: owner mismatch")
			c.Status(http.StatusForbidden)
		} else if status.Code(err) == codes.NotFound {
			log.Println("ERROR: task not found")
			c.Status(http.StatusBadRequest)
		} else {
			log.Println("ERROR: internal database error")
			c.Status(http.StatusInternalServerError)
		}
	} else {
		c.JSON(okStatus, resp)
	}
}

func CreateTask(c *gin.Context) {
	var task protos.Task
	if err := c.BindJSON(&task); err != nil {
		log.Println("ERROR: invalid request body")
		c.Status(http.StatusBadRequest)
		return
	}

	taskOwner, _ := c.Get("user")
	task.Owner = taskOwner.(string)
	resp, err := grpc.GRPCClient.CreateTask(context.Background(), &task)
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
	task.Owner = taskOwner.(string)
	id, _ := strconv.Atoi(c.Param("id"))
	task.Id = int64(id)
	resp, err := grpc.GRPCClient.UpdateTask(context.Background(), &task)
	returnResponse(c, resp, err, http.StatusOK)
}

func DeleteTask(c *gin.Context) {
	var taskCreds protos.TaskCreds
	taskOwner, _ := c.Get("user")
	taskCreds.Owner = taskOwner.(string)
	id, _ := strconv.Atoi(c.Param("id"))
	taskCreds.Id = int64(id)
	resp, err := grpc.GRPCClient.DeleteTask(context.Background(), &taskCreds)
	returnResponse(c, resp, err, http.StatusOK)
}

func GetTask(c *gin.Context) {
	var taskCreds protos.TaskCreds
	taskOwner, _ := c.Get("user")
	taskCreds.Owner = taskOwner.(string)
	id, _ := strconv.Atoi(c.Param("id"))
	taskCreds.Id = int64(id)
	resp, err := grpc.GRPCClient.GetTask(context.Background(), &taskCreds)
	returnResponse(c, resp, err, http.StatusOK)
}

func GetTasksByPage(c *gin.Context) {
	var pageReq protos.PageRequest

	pageReq.Owner = c.Query("user")

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

	resp, err := grpc.GRPCClient.GetTasksByPage(context.Background(), &pageReq)
	returnResponse(c, resp, err, http.StatusOK)
}
