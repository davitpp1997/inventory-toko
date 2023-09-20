package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(respon *gin.Context) {
	respon.JSON(http.StatusOK, gin.H{
		"message": "selamat datang di Aplikasi Inventory Toko",
	})
}
