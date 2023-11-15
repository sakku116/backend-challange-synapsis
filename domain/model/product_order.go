package model

import "gorm.io/gorm"

type ProductOrder struct {
	gorm.Model
	CartID   uint      `json:"cart_id" gorm:"default:null"`
	Quantity int       `json:"quantity"`
	Products []Product `json:"products"`
}
