package auth

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"gin-todo-app/pkg/orm"
	"gin-todo-app/pkg/orm/models"
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

	var userExists models.User
	orm.Db.Where("username = ?", json.Username).First(&userExists)
	if userExists.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "User does not exist"})
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(userExists.Password), []byte(json.Password))
	if err == nil {
		hmacSampleSecret = []byte(os.Getenv("JWT_SECRET_KEY"))
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"userId": userExists.ID,
			"exp":    time.Now().Add(time.Minute * 30).Unix(),
		})
		tokenString, err := token.SignedString(hmacSampleSecret)
		fmt.Println(tokenString, err)
		c.JSON(http.StatusOK, gin.H{"status": "Login successful", "token": tokenString})

	} else {
		c.JSON(http.StatusOK, gin.H{"status": "Login failed"})

	}

}
