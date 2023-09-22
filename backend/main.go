package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"backend-inventory-toko/handler"
	"backend-inventory-toko/repository"
	service "backend-inventory-toko/service"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/inventory_toko?charset=utf8mb4&parseTime=True&loc=Local"
	var db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Database Connection Error")
	}

	route := gin.Default()
	route.Use(cors.Default())
	apiV1 := route.Group("/v1")

	//Barang
	barangRepo := repository.NewBarangRepository(db)
	barangService := service.NewBarangService(barangRepo)
	barangHandler := handler.NewBarangHandler(barangService)

	apiV1.GET("/list_barang", barangHandler.ListBarangs)
	apiV1.GET("/detail_barang/:id", barangHandler.DetailBarang)
	apiV1.POST("/add_barang", barangHandler.AddBarang)
	apiV1.POST("/update_barang/:id", barangHandler.UpdateBarang)
	apiV1.GET("/delete_barang/:id", barangHandler.DeleteBarang)

	// Stock
	stockRepo := repository.NewStockRepository(db)
	stockService := service.NewStockService(stockRepo)
	stockHandler := handler.NewStockHandler(stockService)

	apiV1.POST("/add_stock", stockHandler.NewStock)

	route.Run(":8001")
}

// func koneksiDatabase() {
// 	dsn := "root:@tcp(127.0.0.1:3306)/inventory_toko?charset=utf8mb4&parseTime=True&loc=Local"
// 	var db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		log.Fatal("Database Connection Error")
// 	}

// 	db.Begin()
// 	println("Database Connection Success")

// db.Begin()

// barang := repository.NewBarangRepository(db)

// brg := model.Barang{
// 	Kode:        "BRG0123",
// 	Nama:        "Gula",
// 	Deskripsi:   "Gula putih",
// 	Jumlah:      5,
// 	StatusAktif: true,
// }
// barang.Create(brg)

// barangs, errors := barang.FindAll()
// if errors != nil {
// 	log.Fatal("Database Connection Error")
// }

// for _, data := range barangs {
// 	fmt.Println("nama barang = ", data.Nama)
// }
// }
