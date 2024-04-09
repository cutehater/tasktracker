package controllers

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"

	"userservice/db"
	"userservice/schemas"
)

func CreateUser(c *gin.Context) {
	var loginReq schemas.LoginRequest

	if err := c.BindJSON(&loginReq); err != nil {
		log.Println("ERROR: invalid request body")
		c.Status(http.StatusBadRequest)
		return
	}

	var user schemas.UserData
	db.DB.First(&user, "login = ?", loginReq.Login)

	if user.ID != 0 {
		log.Println("ERROR: user already exists")
		c.Status(http.StatusConflict)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(loginReq.Password), 10)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	newUser := schemas.UserData{Login: loginReq.Login, PasswordHash: string(hash)}
	res := db.DB.Create(&newUser)
	if res.Error != nil {
		log.Println(res.Error)
		c.Status(http.StatusInternalServerError)
	} else {
		c.Status(http.StatusCreated)
	}
}

func LoginUser(c *gin.Context) {
	var loginReq schemas.LoginRequest

	if err := c.BindJSON(&loginReq); err != nil {
		log.Println("ERROR: invalid request body")
		c.Status(http.StatusBadRequest)
		return
	}

	var user schemas.UserData
	db.DB.First(&user, "login = ?", loginReq.Login)

	if user.ID == 0 {
		log.Println("ERROR: user does not exist")
		c.Status(http.StatusBadRequest)
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(loginReq.Password))
	if err != nil {
		log.Println(err)
		c.Status(http.StatusBadRequest)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"login": loginReq.Login,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
	} else {
		c.SetCookie("JSESSIONID", tokenString, 3600*24, "", "", false, true)
	}
}

func UpdateUser(c *gin.Context) {
	login := c.Param("login")
	if user, _ := c.Get("user"); user.(string) != login {
		log.Println("ERROR: forbidden to change other users data")
		c.Status(http.StatusForbidden)
		return
	}

	var user schemas.UserData
	db.DB.First(&user, "login = ?", login)

	if user.ID == 0 {
		log.Println("ERROR: user does not exist")
		c.Status(http.StatusInternalServerError)
		return
	}

	var updateUser schemas.UserData
	if err := c.BindJSON(&updateUser); err != nil {
		log.Println("ERROR: invalid request body")
		c.Status(http.StatusBadRequest)
		return
	}

	updateUser.Login = user.Login
	updateUser.PasswordHash = user.PasswordHash

	res := db.DB.Model(&user).Updates(updateUser)
	if res.Error != nil {
		log.Println(res.Error)
		c.Status(http.StatusInternalServerError)
	} else {
		c.Status(http.StatusOK)
	}
}
