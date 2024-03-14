package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"os"
	"time"
)

func IsAuthorized(c *gin.Context) {
	cookie, err := c.Cookie("JSESSIONID")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		return os.Getenv("SECRET"), nil
	})
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if float64(time.Now().Unix()) > claims["exp"].(float64) || c.Param("login") != claims["login"] {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
