package main

import (
	"github.com/gin-gonic/gin"
	"user_service/controllers"
	"user_service/db"
	"user_service/middlewares"
)

func main() {
	db.ConnectToDb()
	r := gin.Default()

	r.POST("/user", controllers.CreateUser)

	r.POST("/user/login", gin.BasicAuth(gin.Accounts{
		"admin": "secret",
	}), controllers.LoginUser)

	r.PUT("/user/:login", middlewares.IsAuthorized, controllers.UpdateUser)

	r.Run()
}
