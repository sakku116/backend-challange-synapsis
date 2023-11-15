package main

import (
	"synapsis/config"
	"synapsis/repository"

	"github.com/gookit/goutil/dump"
)

func Playground() {
	db, err := config.NewDb(config.Envs.DB_URI)
	if err != nil {
		panic(err)
	}
	repo := repository.NewProductRepo(db)
	_ = repo

	test, err := repo.GetCategoryList()
	dump.P(test)
}
