package models

type Income struct {
	Jumlah int `gorm:"column:total" json:"total"`
}

type Outcome struct {
	Jumlah int `gorm:"column:total" json:"total"`
}
