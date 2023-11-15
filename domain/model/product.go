package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID             string   `json:"id" gorm:"primaryKey"`
	ProductOrderID string   `json:"order_id" gorm:"default:null"`
	Name           string   `json:"name"`
	Price          float64  `json:"price"`
	Category       Category `json:"category"`
	CategoryID     string   `json:"category_id"`
}
