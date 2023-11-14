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
func Profile(c *gin.Context) {
	userId := c.MustGet("userId").(float64)
	var user []orm.User
	orm.Db.First(&user, userId)
	c.JSON(200, gin.H{"status": "ok", "message": "User read successs", "users": user})

}
