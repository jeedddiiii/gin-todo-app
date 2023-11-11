package main

import (
	Controller "gin-todo-app/controller"
	"gin-todo-app/orm"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	orm.Dbconnect()
	r.POST("/register", Controller.Register)
	r.Use(cors.Default())

	r.Run(":8080")

}
