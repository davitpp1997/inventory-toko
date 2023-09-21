package repository

import (
	"backend-inventory-toko/model"

	"gorm.io/gorm"
)

type BarangRepository interface {
	FindAll() ([]model.Barang, error)
	Create(barang model.Barang) (model.Barang, error)
	Detail(ID int) (model.Barang, error)
	Update(barang model.Barang) (model.Barang, error)
	Delete(model.Barang) (model.Barang, error)
}

type barangRepository struct {
	databaseToko *gorm.DB
}

func NewBarangRepository(databaseToko *gorm.DB) *barangRepository {
	return &barangRepository{databaseToko}
}

func (bR *barangRepository) FindAll() ([]model.Barang, error) {
	var barangs []model.Barang

	err := bR.databaseToko.Find(&barangs).Error

	return barangs, err
}

func (bR *barangRepository) Detail(ID int) (model.Barang, error) {
	var barang model.Barang

	err := bR.databaseToko.Find(&barang, ID).Error

	return barang, err
}

func (bR *barangRepository) Create(barang model.Barang) (model.Barang, error) {
	err := bR.databaseToko.Create(&barang).Error

	return barang, err
}

func (bR *barangRepository) Update(barang model.Barang) (model.Barang, error) {
	err := bR.databaseToko.Save(&barang).Error

	return barang, err
}

func (bR *barangRepository) Delete(barang model.Barang) (model.Barang, error) {
	err := bR.databaseToko.Delete(&barang).Error

	return barang, err
}
