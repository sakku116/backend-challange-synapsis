package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ProductOrderID string  `json:"order_id"`
	Name           string  `json:"name"`
	Price          float64 `json:"price"`
	CategoryID     string  `json:"category_id"`
}
