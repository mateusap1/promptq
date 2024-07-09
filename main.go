package main

import (
	"context"
	"fmt"
	"log"

	"net/http"

	"github.com/gin-gonic/gin"

	"promptq/ent"

	_ "github.com/mattn/go-sqlite3"
)

func getHealth(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "healthy",
	})
}

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
	router.GET("/health", getHealth)

	router.Run("localhost:8080")
}
