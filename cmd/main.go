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
	"github.com/mateusap1/promptq/pkg/user"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func Open(dbUrl string) *ent.Client {
	db, err := sql.Open("pgx", dbUrl)
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
		return
	}

	postgresUrl, present := os.LookupEnv("POSTGRES_URL")
	if !present {
		log.Fatalf("POSTGRES_URL not defined.")
		return
	}

	client := Open(postgresUrl)
	defer client.Close()

	ctx := context.Background()

	// Run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	us := user.CreateService(ctx, client)
	router := gin.Default()

	router.GET("/health", api.GetHealth)

	user := router.Group("")
	{
		user.POST("/login", loginEndpoint)
	}

	router.POST("/user", func(c *gin.Context) { api.CreateUser(c, us) })

	// For running in production just use router.Run()
	router.Run()
}
