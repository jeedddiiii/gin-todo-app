package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterBody struct {
	Username             string `json:"username" binding:"required"`
	Password             string `json:"password" binding:"required"`
	PasswordConfirmation string `json:"passwordcon" binding:"required"`
}

func Register(c *gin.Context) {
	var json RegisterBody
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Your logic for registration here
	// ...

	// Respond with a success message or appropriate data
	c.JSON(200, gin.H{"Registration successful": json})
}
