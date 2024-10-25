package main

import (
	"context"
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/mateusap1/promptq/api"
	"github.com/mateusap1/promptq/ent"
	"github.com/mateusap1/promptq/pkg/prompt"
	"github.com/mateusap1/promptq/pkg/user"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func Open(databaseUrl string) *ent.Client {
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		log.Fatalf("failed oppening connecting to postgresql %v", err)
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv))
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("failed accessing env variable: %v", err)
	}

	postgresUrl, present := os.LookupEnv("POSTGRES_URL")
	if !present {
		log.Fatalf("POSTGRES_URL not defined.")
	}

	client := Open(postgresUrl)
	defer client.Close()

	ctx := context.Background()

	// Run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	us := user.CreateService(ctx, client)
	ps := prompt.CreateService(ctx, client)

	router := gin.Default()
	router.GET("/health", api.GetHealth)
	router.POST("/prompt", func(c *gin.Context) { api.CreatePrompt(c, us, ps) })
	router.PUT("/prompt", func(c *gin.Context) { api.QueuePrompt(c, ps) })
	router.POST("/prompt/:id", func(c *gin.Context) { api.AnswerPrompt(c, ps) })
	router.GET("/prompt/:id", func(c *gin.Context) { api.GetPrompt(c, ps) })
	router.GET("/prompt/list", func(c *gin.Context) { api.GetPrompts(c, us, ps) })
	router.POST("/user", func(c *gin.Context) { api.CreateUser(c, us) })

	// For running in production just use router.Run()
	router.Run()
}
