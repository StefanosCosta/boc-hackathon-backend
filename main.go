package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"OpenBankingAPI/config"
	"OpenBankingAPI/database"
	"OpenBankingAPI/routes"
)

func main() {
	err := config.LoadConfig("config.yml")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println("Client ID:", config.ClientID)
	fmt.Println("Client Secret:", config.ClientSecret)
	fmt.Println("Base URI:", config.BaseURI)

	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "http://localhost",
	}))

	routes.Setup(app)

	app.Listen(":8000")
}
