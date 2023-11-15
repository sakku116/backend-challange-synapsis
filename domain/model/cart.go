package model

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	ID            string         `json:"id" gorm:"primaryKey"`
	UserID        string         `json:"user_id"`
	ProductOrders []ProductOrder `json:"product_orders"`
}
