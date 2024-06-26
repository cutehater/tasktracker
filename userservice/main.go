package main

import (
	"github.com/gin-gonic/gin"

	"userservice/controllers"
	"userservice/db"
	"userservice/grpc"
	"userservice/message_broker"
	"userservice/middlewares"
)

func main() {
	db.ConnectToDb()
	grpc.CreateGRPCClient()
	message_broker.InitMessageProducer()

	r := gin.Default()

	// User methods
	r.POST("/user", controllers.CreateUser)
	r.GET("/user/login", controllers.LoginUser)
	r.PUT("/user/:login", middlewares.IsAuthorized, controllers.UpdateUser)

	// Task methods
	r.POST("/task", middlewares.IsAuthorized, controllers.CreateTask)
	r.PUT("/task/:id", middlewares.IsAuthorized, controllers.UpdateTask)
	r.DELETE("/task/:id", middlewares.IsAuthorized, controllers.DeleteTask)
	r.GET("/task/:id", middlewares.IsAuthorized, controllers.GetTask)
	r.GET("/task/page", middlewares.IsAuthorized, controllers.GetTasksByPage)

	// Like/View methods
	r.POST("/like", middlewares.IsAuthorized, controllers.LikeTask)
	r.POST("/view", middlewares.IsAuthorized, controllers.ViewTask)

	r.Run()
}
