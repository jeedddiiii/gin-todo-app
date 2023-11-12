package controller

import (
	"gin-todo-app/orm"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

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
		c.JSON(http.StatusOK, gin.H{"status": "Login successful"})

	} else {
		c.JSON(http.StatusOK, gin.H{"status": "Login failed"})

	}

}
