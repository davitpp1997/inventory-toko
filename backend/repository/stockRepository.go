package repository

import (
	"backend-inventory-toko/model"

	"gorm.io/gorm"
)

type StockRepository interface {
	Create(stock model.Stock) (model.Stock, error)
}

type stockRepository struct {
	databaseToko *gorm.DB
}

func NewStockRepository(databaseToko *gorm.DB) *stockRepository {
	return &stockRepository{databaseToko}
}

func (sR *stockRepository) Create(stock model.Stock) (model.Stock, error) {
	err := sR.databaseToko.Create(&stock).Error

	return stock, err
}
