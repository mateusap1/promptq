package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/mateusap1/promptq/api"
	"github.com/mateusap1/promptq/ent"

	_ "github.com/mattn/go-sqlite3"
)

func MakePromptRequest(ctx context.Context, client *ent.Client, prompt string) (*ent.PromptRequest, error) {
	u, err := client.PromptRequest.
		Create().
		SetPrompt(prompt).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed creating prompt request: %w", err)
	}

	log.Println("prompt request was created: ", u)

	return u, nil
}

func main() {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")

	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	defer client.Close()

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	router := gin.Default()
	router.GET("/health", api.GetHealth)

	router.Run("localhost:8080")
}
