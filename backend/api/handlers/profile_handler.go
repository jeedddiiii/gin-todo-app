package user_handler

import (
	"gin-todo-app/pkg/orm"
	"gin-todo-app/pkg/orm/models"

	"github.com/gin-gonic/gin"
)

func Profile(c *gin.Context) {
	userId := c.MustGet("userId").(float64)
	var user []models.User
	orm.Db.First(&user, userId)
	c.JSON(200, gin.H{"status": "ok", "message": "User read successs", "users": user})

}
