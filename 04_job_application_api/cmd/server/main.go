package main

import (
	"job_portal/internal/repository"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := repository.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	r := gin.Default()

	r.Run(":8080")
}
