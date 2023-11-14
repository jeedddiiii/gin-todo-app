package auth

import (
	"gin-todo-app/pkg/orm"

	"gin-todo-app/pkg/orm/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type RegisterBody struct {
	Username             string `json:"username" binding:"required"`
	Password             string `json:"password" binding:"required"`
	PasswordConfirmation string `json:"password_con" binding:"required"`
}

func Register(c *gin.Context) {
	var json RegisterBody
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userExists models.User
	orm.Db.Where("username = ?", json.Username).First(&userExists)
	if userExists.ID > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User Exists"})
		return
	}

	if json.Password != json.PasswordConfirmation {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password and password confirmation do not match"})
		return
	}
	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(json.Password), 10)
	user := models.User{Username: json.Username, Password: string(encryptedPassword)}
	orm.Db.Create(&user)
	if user.ID > 0 {
		c.JSON(200, gin.H{"status": "ok", "message": "User created successfully", "userId": user.ID})
	} else {
		c.JSON(400, gin.H{"status": "error", "message": "User created failed"})
	}
}
