package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	ID       string    `json:"id" gorm:"primaryKey"`
	Name     string    `json:"name"`
	Products []Product `json:"products"`
}
