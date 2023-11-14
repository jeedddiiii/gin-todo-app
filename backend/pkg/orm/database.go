package orm

import (
	"gin-todo-app/pkg/orm/models"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Dbconnect() {
	dsn := os.Getenv("MYSQL_DNS")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	if err := db.AutoMigrate(&models.User{}); err != nil {
		panic("failed to auto-migrate: " + err.Error())
	}

	Db = db

	db.AutoMigrate(&models.User{})
}
