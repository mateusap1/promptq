package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/mateusap1/promptq/api"
	"github.com/mateusap1/promptq/ent"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")

	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	defer client.Close()

	ctx := context.Background()

	// Run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	router := gin.Default()
	router.GET("/health", api.GetHealth)
	router.POST("/prompt", func(c *gin.Context) { api.CreatePrompt(c, ctx, client) })
	router.PUT("/prompt", func(c *gin.Context) { api.QueuePrompt(c, ctx, client) })

	// For running in production just use router.Run()
	router.Run("localhost:8080")
}
