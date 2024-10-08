package main

import (
	"job_portal/internal/repository"
	"job_portal/internal/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := repository.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	r := gin.Default()
	routes.InitRoutes(r, db)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
