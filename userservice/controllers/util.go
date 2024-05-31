package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func returnResponse(c *gin.Context, resp any, err error, okStatus int) {
	if err != nil {
		if status.Code(err) == codes.PermissionDenied {
			log.Println("ERROR: owner mismatch")
			c.Status(http.StatusForbidden)
		} else if status.Code(err) == codes.NotFound {
			log.Println("ERROR: task or user not found")
			c.Status(http.StatusBadRequest)
		} else {
			log.Println("ERROR: internal database error")
			c.Status(http.StatusInternalServerError)
		}
	} else {
		c.JSON(okStatus, resp)
	}
}
