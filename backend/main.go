package main

import (
	"github.com/gin-gonic/gin"

	"backend-inventory-toko/handler"
)

func main() {
	route := gin.Default()
	apiV1 := route.Group("/v1")

	apiV1.GET("/", handler.Home)
	route.Run()
}
