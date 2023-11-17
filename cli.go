package main

import (
	"fmt"
	"os"
	"synapsis/config"
	"synapsis/repository"
	"synapsis/utils/data"
)

func CliHandler(args []string) {
	args = args[1:]

	db, err := config.NewDb(config.Envs.DB_URI)
	if err != nil {
		panic(err)
	}

	userRepo := repository.NewUserRepo(db)
	productRepo := repository.NewProductRepo(db)

	switch args[0] {
	case "seed-superuser":
		fmt.Println("running seed superuser...")
		data.SeedSuperuser(userRepo, args[1:]...)
	case "seed-data":
		fmt.Println("running seed data...")
		data.SeedData(productRepo)
	case "playground":
		fmt.Println("running playground...")
		Playground()
	default:
		fmt.Println("invalid command")
		os.Exit(1)
	}
	fmt.Println("done")
}
