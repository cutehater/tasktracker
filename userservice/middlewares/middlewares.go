package middlewares

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func IsAuthorized(c *gin.Context) {
	cookie, err := c.Cookie("JSESSIONID")
	if err != nil {
		log.Println("ERROR: invalid cookie header")
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		log.Println("ERROR: error parsing cookie token")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			log.Println("ERROR: Cookie expired")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("user", claims["login"])
		c.Next()
	} else {
		log.Println("ERROR: Cookie not found")
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
