package models

import "time"

type Saldo struct {
	ID          int       `gorm:"column:id;AUTO_INCREMENT" json:"id"`
	Namabank    string    `gorm:"column:nama_bank" json:"nama_bank" form:"nama_bank"`
	JumlahSaldo int       `gorm:"column:jumlah_saldo" json:"jumlah_saldo" form:"jumlah_saldo"`
	CreateAt    time.Time `gorm:"column:created_at" json:"created_at"`
	UpdateAt    time.Time `gorm:"column:updated_at" json:"update_dat"`
}
