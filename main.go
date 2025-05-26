package main

import (
	"authentication/config"
	"authentication/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	config.InitDB()

	e := echo.New()
	routes.InitRoutes(e)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))
}
