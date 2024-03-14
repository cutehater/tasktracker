package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
	"user_service/db"
	"user_service/schemas"
)

var jwtKey = []byte("my_secret_key")
var tokens []string

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func CreateUser(c *gin.Context) {
	var loginReq schemas.LoginRequest

	if err := c.BindJSON(&loginReq); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	var user schemas.UserData
	db.DB.First(&user, "login = ?", loginReq.Login)

	if user.ID != 0 {
		c.Status(http.StatusConflict) // User already exists
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(loginReq.Password), 10)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	newUser := schemas.UserData{Login: loginReq.Login, PasswordHash: string(hash)}
	res := db.DB.Create(&newUser)
	if res.Error != nil {
		c.Status(http.StatusInternalServerError)
	} else {
		c.Status(http.StatusCreated)
	}
}

func LoginUser(c *gin.Context) {
	var loginReq schemas.LoginRequest

	if err := c.BindJSON(&loginReq); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	var user schemas.UserData
	db.DB.First(&user, "login = ?", loginReq.Login)

	if user.ID == 0 {
		c.Status(http.StatusBadRequest) // No such user
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(loginReq.Password))
	if err != nil {
		c.Status(http.StatusBadRequest) // Login or password incorrect
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"login": loginReq.Login,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.Status(http.StatusInternalServerError)
	} else {
		c.SetCookie("JSESSIONID", tokenString, 3600*24, "", "", false, true)
	}
}

func UpdateUser(c *gin.Context) {
	login := c.Param("login")
	var user schemas.UserData
	db.DB.First(&user, "login = ?", login)

	if user.ID == 0 {
		c.Status(http.StatusInternalServerError) // User is authorized but not in db? Something strange
		return
	}

	var updateUser schemas.UserData
	if err := c.BindJSON(&updateUser); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	updateUser.Login = user.Login
	updateUser.PasswordHash = user.PasswordHash

	res := db.DB.Model(&user).Updates(updateUser)
	if res.Error != nil {
		c.Status(http.StatusInternalServerError)
	} else {
		c.Status(http.StatusOK)
	}
}
