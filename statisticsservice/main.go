package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"statisticsservice/db"
	"statisticsservice/message_broker"
)

func main() {
	db.ConnectToDb()
	message_broker.RunConsumer()
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		db.GetAllEvents()
		c.Status(http.StatusOK)
	})

	r.Run()
}
