package user

import (
	"gin-todo-app/orm"

	"github.com/gin-gonic/gin"
)

func ReadAll(c *gin.Context) {

	var users []orm.User
	orm.Db.Find(&users)
	c.JSON(200, gin.H{"status": "ok", "message": "User read successs", "users": users})

}
