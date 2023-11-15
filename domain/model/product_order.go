package model

import "gorm.io/gorm"

type ProductOrder struct {
	gorm.Model
	ID       string  `json:"id" gorm:"primaryKey"`
	CartID   string  `json:"cart_id" gorm:"default:null"`
	Quantity int     `json:"quantity"`
	Product  Product `json:"product"`
}
