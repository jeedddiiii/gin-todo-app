package main

import (
	Controller "gin-todo-app/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/register", Controller.Register)

	r.Run(":8080")

}
