package main

import (
	"fmt"
	"log"
	"os"
	"synapsis/config"

	"github.com/gin-gonic/gin"

	_ "synapsis/docs"
)

// @title Backend Challange Synapsis API
// @securitydefinitions.apiKey BearerAuth
// @in header
// @name Authorization
// @default Bearer {token}
func main() {
	args := os.Args
	if len(args) > 1 {
		if args[1] != "playground" {
			CliHandler(args)
		} else {
			Playground()
		}
	} else {
		log.Printf("Envs: %v", config.Envs)
		log.Println("starting rest api app...")

		router := gin.Default()
		SetupServer(router)
		router.Run(config.Envs.ADDR)

		fmt.Println("starting rest api app...")
	}
}
