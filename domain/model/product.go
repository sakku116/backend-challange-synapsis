package model

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID        string         `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`

	ProductOrderID string  `json:"product_order_id" gorm:"default:null"`
	Name           string  `json:"name"`
	Price          float64 `json:"price"`
	Category       string  `json:"category"`
}
