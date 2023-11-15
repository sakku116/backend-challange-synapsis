package main

import (
	"fmt"
	"synapsis/config"
	"synapsis/repository"
)

func Playground() {
	db, err := config.NewDb(config.Envs.DB_URI)
	if err != nil {
		panic(err)
	}
	repo := repository.NewProductRepo(db)
	_ = repo

	test, err := repo.GetList("", "", 1, 2, "created_at", "desc")
	fmt.Println(test)
}
