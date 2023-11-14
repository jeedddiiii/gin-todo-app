package user_handler

import (
	"gin-todo-app/pkg/orm"
	"gin-todo-app/pkg/orm/models"

	"github.com/gin-gonic/gin"
)

func ReadAll(c *gin.Context) {

	var users []models.User
	orm.Db.Find(&users)
	c.JSON(200, gin.H{"status": "ok", "message": "User read successs", "users": users})

}
