package service

import (
	"backend-inventory-toko/model"
	"backend-inventory-toko/repository"
)

type StockService interface {
	Create(barang model.Stock) (model.Stock, error)
}

type stockService struct {
	stockRepository repository.StockRepository
}

func NewStockService(stockRepository repository.StockRepository) *stockService {
	return &stockService{stockRepository}
}

func (sS *stockService) Create(stock model.Stock) (model.Stock, error) {
	brg := model.Stock{
		BarangID: stock.BarangID,
		Type:     stock.Type,
		Jumlah:   stock.Jumlah,
	}

	newBrg, err := sS.stockRepository.Create(brg)

	return newBrg, err
}
