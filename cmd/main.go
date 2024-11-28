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

	// Change Origins when running in production
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.GET("/health", api.GetHealth)

	authRouter := router.Group("/auth")
	authRouter.POST("/register", func(c *gin.Context) { api.Register(c, db) })
	authRouter.POST("/login", func(c *gin.Context) { api.Login(c, db) })
	authRouter.POST("/signout", auth.AuthMiddleware(db), func(c *gin.Context) { api.SignOut(c, db) })

	authRouter.POST("/email/validate", func(c *gin.Context) { api.ValidateEmail(c, db) })
	authRouter.POST("/email/validate/resend", auth.AuthMiddleware(db), func(c *gin.Context) { api.ResendValidateEmail(c, db) })

	protectedRouter := router.Group("", auth.AuthMiddleware(db))

	threadRouter := protectedRouter.Group("/thread")
	threadRouter.POST("/create", func(c *gin.Context) { api.CreateThread(c, db) })

	// For running in production just use router.Run()
	router.Run()
}
