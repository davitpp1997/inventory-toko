package model

import (
	"time"
)

type Barang struct {
	ID          uint32    `json:"id"`
	Kode        string    `gorm:"size:10;not null;unique" json:"kode" binding:"required" usage:"unique"`
	Nama        string    `gorm:"size:20;" json:"nama"`
	Deskripsi   string    `gorm:"type:text;" json:"deskripsi"`
	Jumlah      int32     `json:"jumlah"`
	StatusAktif bool      `gorm:"default:true;" json:"status_aktif"`
	CreatedAt   time.Time `json:"created_ad"`
	UpdatedAt   time.Time `json:"updated_at"`
	Stock       []Stock   `gorm:"foreignKey:BarangID;references:ID" json:"stock"`
}
