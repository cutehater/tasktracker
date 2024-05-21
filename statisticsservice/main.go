package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"statisticsservice/message_broker"
)

func main() {
	message_broker.RunConsumer()
	r := gin.Default()

	r.GET("/", func(c *gin.Context) { c.Status(http.StatusOK) })

	r.Run()
}
