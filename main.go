package main

import (
	"log"
	"os"

	"atm-api.com/api"
	"atm-api.com/api/routes"
	"github.com/joho/godotenv"
)

func main() {
	initDotEnv()
	initAPI()
}

func initDotEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func initAPI() {
	withdrawalRoutes := InitializeWithdrawalRoutes()
	routes := []routes.BaseRoutes{withdrawalRoutes}
	server := api.Server{Routes: routes}
	r := server.SetupRoutes()
	r.Run(":" + os.Getenv("APP_PORT"))
}
