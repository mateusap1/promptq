package main

import (
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/mateusap1/promptq/api"
	auth "github.com/mateusap1/promptq/middleware"
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

	db := utils.OpenPostgresDB(pgURL)
	defer db.Close()

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.GET("/health", api.GetHealth)
	router.POST("/signup", func(c *gin.Context) { api.SignUp(c, db) })
	router.POST("/signin", func(c *gin.Context) { api.SignIn(c, db) })

	protected := router.Group("", auth.AuthMiddleware(db))
	protected.GET("/protected", func(c *gin.Context) {})

	// For running in production just use router.Run()
	router.Run()
}
