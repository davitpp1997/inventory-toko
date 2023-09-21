package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"backend-inventory-toko/model"
	"backend-inventory-toko/service"
)

type barangHandler struct {
	brgService service.BarangService
}

func NewBarangHandler(brgService service.BarangService) *barangHandler {
	return &barangHandler{brgService}
}

func (bH *barangHandler) ListBarangs(ctx *gin.Context) {
	barangs, err := bH.brgService.FindAll()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "List Data Barang",
		"data":    barangs,
	})

}

func (bH *barangHandler) DetailBarang(ctx *gin.Context) {
	idBarang := ctx.Param("id")
	id, _ := strconv.Atoi(idBarang)
	barang, err := bH.brgService.Detail(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Detail Barang",
		"data":    barang,
	})

}

func (bH *barangHandler) AddBarang(ctx *gin.Context) {
	var dataBarang model.Barang

	validasi := ctx.ShouldBindJSON(&dataBarang)

	if validasi != nil {

		for _, e := range validasi.(validator.ValidationErrors) {
			pesanError := fmt.Sprintf("Error on field %s, Condition: %s", e.Field(), e.ActualTag())

			ctx.JSON(http.StatusBadRequest, gin.H{
				"status":  "failed",
				"message": pesanError,
			})
			return
		}

	}

	barang, err := bH.brgService.Create(dataBarang)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Barang berhasil ditambahkan",
		"data":    barang,
	})
}

func (bH *barangHandler) UpdateBarang(ctx *gin.Context) {
	var dataBarang model.Barang

	validasi := ctx.ShouldBindJSON(&dataBarang)

	if validasi != nil {

		for _, e := range validasi.(validator.ValidationErrors) {
			pesanError := fmt.Sprintf("Error on field %s, Condition: %s", e.Field(), e.ActualTag())

			ctx.JSON(http.StatusBadRequest, gin.H{
				"status":  "failed",
				"message": pesanError,
			})
			return
		}

	}

	idBarang := ctx.Param("id")
	id, _ := strconv.Atoi(idBarang)
	barang, err := bH.brgService.Update(id, dataBarang)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Barang berhasil dirubah",
		"data":    barang,
	})
}

func (bH *barangHandler) DeleteBarang(ctx *gin.Context) {

	idBarang := ctx.Param("id")
	id, _ := strconv.Atoi(idBarang)
	barang, err := bH.brgService.Delete(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Barang berhasil dihapus",
		"data":    barang,
	})
}
