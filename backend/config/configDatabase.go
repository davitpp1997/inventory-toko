package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dsn = "root:@tcp(127.0.0.1:3306)/inventory_toko?charset=utf8mb4&parseTime=True&loc=Local"
var db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

func StartDatabaseConnection() {

	if err != nil {
		log.Fatal("Database Connection Error")
	}

	db.Begin()
	println("Database Connection Success")
}
