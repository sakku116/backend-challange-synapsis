package main

import (
	"fmt"
	"os"
	"synapsis/cli"
	"synapsis/config"
	"synapsis/repository"
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
		cli.SeedSuperuser(userRepo, args[1:]...)
	case "seed-data":
		fmt.Println("running seed data...")
		fmt.Println("seeding products")
		cli.SeedProduct(productRepo)
	default:
		fmt.Println("invalid command")
		os.Exit(1)
	}
	fmt.Println("done")
}
