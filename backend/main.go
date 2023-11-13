package main

import (
	"fmt"
	Controller "gin-todo-app/controller"
	"gin-todo-app/orm"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
	}
	r := gin.Default()
	r.Use(cors.Default())

	orm.Dbconnect()
	r.POST("/register", Controller.Register)
	r.POST("/login", Controller.Login)

	r.Run(":8080")

}
