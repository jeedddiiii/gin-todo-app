package orm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Dbconnect() {
	dsn := "root@tcp(127.0.0.1:3306)/gotodo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	if err := db.AutoMigrate(&User{}); err != nil {
		panic("failed to auto-migrate: " + err.Error())
	}

	Db = db

	db.AutoMigrate(&User{})
}
