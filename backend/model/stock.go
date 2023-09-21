package model

import (
	"time"
)

type Stock struct {
	ID        uint32 `json:"id"`
	BarangID  uint32 `json:"id_barang" binding:"required"`
	Type      string `gorm:"size:10;" json:"type" binding:"required"`
	Jumlah    int32  `json:"jumlah" binding:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
