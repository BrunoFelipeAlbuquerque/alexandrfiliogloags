package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"imperial-fleet/config"
	"imperial-fleet/handlers"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDatabase()

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.SetTrustedProxies(nil)

	handlers.RegisterSpaceshipRoutes(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
