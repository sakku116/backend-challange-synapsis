package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ProductOrderID uint    `json:"order_id" gorm:"default:null"`
	Name           string  `json:"name"`
	Price          float64 `json:"price"`
	CategoryID     uint    `json:"category_id" gorm:"default:null"`
}
