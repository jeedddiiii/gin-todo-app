package controller

import (
	"fmt"
	"gin-todo-app/orm"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var hmacSampleSecret []byte

type LoginBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var json LoginBody
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userExists orm.User
	orm.Db.Where("username = ?", json.Username).First(&userExists)
	if userExists.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "User does not exist"})
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(userExists.Password), []byte(json.Password))
	if err == nil {
		hmacSampleSecret = []byte("my_secret_key")
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"userId": userExists.ID,
		})
		tokenString, err := token.SignedString(hmacSampleSecret)
		fmt.Println(tokenString, err)
		c.JSON(http.StatusOK, gin.H{"status": "Login successful", "token": tokenString})

	} else {
		c.JSON(http.StatusOK, gin.H{"status": "Login failed"})

	}

}
