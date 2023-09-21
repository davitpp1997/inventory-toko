package handler

import (
	// "encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"backend-inventory-toko/model"
	"backend-inventory-toko/service"
)

type stockHandler struct {
	stockService service.StockService
}

func NewStockHandler(stockService service.StockService) *stockHandler {
	return &stockHandler{stockService}
}

func (sH *stockHandler) AddStock(ctx *gin.Context) {
	var dataStock model.Stock

	validasi := ctx.ShouldBindJSON(&dataStock)

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

	stock, err := sH.stockService.Create(dataStock)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Stock berhasil ditambahkan",
		"data":    stock,
	})
}

func (sH *stockHandler) NewStock(ctx *gin.Context) {
	var dataStock model.Stock

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Stock berhasil ditambahkan",
		"data":    dataStock,
	})
}
