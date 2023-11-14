package config

import (
	"fmt"
	"synapsis/domain/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDb(uri string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(Envs.DB_URI), &gorm.Config{})
	fmt.Println(uri)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&model.User{})
	return db, nil
}
