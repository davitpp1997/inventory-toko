package service

import (
	"backend-inventory-toko/model"
	"backend-inventory-toko/repository"
)

type BarangService interface {
	FindAll() ([]model.Barang, error)
	Detail(ID int) (model.Barang, error)
	Create(barang model.Barang) (model.Barang, error)
	Update(ID int, barang model.Barang) (model.Barang, error)
	Delete(ID int) (model.Barang, error)
}

type barangService struct {
	barangRepository repository.BarangRepository
}

func NewBarangService(barangRepository repository.BarangRepository) *barangService {
	return &barangService{barangRepository}
}

func (bS *barangService) FindAll() ([]model.Barang, error) {
	return bS.barangRepository.FindAll()
}

func (bS *barangService) Detail(ID int) (model.Barang, error) {
	return bS.barangRepository.Detail(ID)
}

func (bS *barangService) Create(barang model.Barang) (model.Barang, error) {
	brg := model.Barang{
		Kode:        barang.Kode,
		Nama:        barang.Nama,
		Deskripsi:   barang.Deskripsi,
		Jumlah:      barang.Jumlah,
		StatusAktif: barang.StatusAktif,
	}

	newBrg, err := bS.barangRepository.Create(brg)

	return newBrg, err
}

func (bS *barangService) Update(ID int, barang model.Barang) (model.Barang, error) {
	brg, err := bS.barangRepository.Detail(ID)

	brg.Kode = barang.Kode
	brg.Nama = barang.Nama
	brg.Deskripsi = barang.Deskripsi
	brg.Jumlah = barang.Jumlah
	brg.StatusAktif = barang.StatusAktif

	newBrg, err := bS.barangRepository.Update(brg)

	return newBrg, err
}

func (bS *barangService) Delete(ID int) (model.Barang, error) {
	brg, err := bS.barangRepository.Detail(ID)

	bS.barangRepository.Delete(brg)

	return brg, err
}
