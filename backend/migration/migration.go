package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"backend-inventory-toko/model"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/inventory_toko?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Database Connection Error")
	}

	db.AutoMigrate(
		&model.Barang{},
		&model.Stock{},
	)

	print("Migration Success")
}
