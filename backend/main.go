package main

import (
	"fmt"
	UserHandler "gin-todo-app/api/handlers"
	Auth "gin-todo-app/internal/auth"
	"gin-todo-app/pkg/middleware"

	"gin-todo-app/pkg/orm"

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
	r.POST("/register", Auth.Register)
	r.POST("/login", Auth.Login)
	authorized := r.Group("/users", middleware.JWTAuthen())
	authorized.GET("/readall", UserHandler.ReadAll)
	authorized.GET("/profile", UserHandler.Profile)

	r.Run(":8080")

}
