package models

import (
	"time"
)

type Finances struct {
	ID         int       `gorm:"column:id;AUTO_INCREMENT" json:"id"`
	Tipe       string    `gorm:"column:type" json:"type" form:"type"`
	Jumlah     int       `gorm:"column:jumlah" json:"jumlah" form:"jumlah"`
	AsalTujuan string    `gorm:"column:asaltujuan" json:"asaltujuan" form:"asaltujuan"`
	Tanggal    time.Time `gorm:"column:tanggal" json:"tanggal" form:"tanggal"`
	Keterangan string    `gorm:"column:keterangan" json:"keterangan" form:"keterangan"`
	CreateAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdateAt   time.Time `gorm:"column:updated_at" json:"update_dat"`
}
