package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/mateusap1/promptq/api"
	"github.com/mateusap1/promptq/pkg/utils"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("failed accessing env variable: %v", err)
		return
	}

	pgURL, present := os.LookupEnv("POSTGRES_URL")
	if !present {
		log.Fatalf("POSTGRES_URL env not present")
		return
	}

	db := utils.OpenDB(pgURL)
	defer db.Close()

	router := gin.Default()

	router.GET("/health", api.GetHealth)

	// For running in production just use router.Run()
	router.Run()
}
