package user

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mateusap1/promptq/ent"
	_ "github.com/mattn/go-sqlite3"
)

func setupDatabase(t *testing.T) (*ent.Client, context.Context) {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		t.Fatalf("failed opening connection to sqlite: %v", err)
	}

	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		t.Fatalf("failed creating schema resources: %v", err)
	}

	return client, ctx
}

func TestMakeUser(t *testing.T) {
	client, ctx := setupDatabase(t)
	defer client.Close()

	t.Run("Create user", func(t *testing.T) {
		us, err := MakeUser(ctx, client, "test123")
		if err != nil {
			t.Fatalf("failed creating user: %v", err)
		}

		usCount, err := client.User.Query().Count(ctx)
		if err != nil {
			t.Fatalf("failed counting prompts: %v", err)
		}

		assert.Equal(t, usCount, 1)
		assert.Equal(t, us.Username, "test123")
	})
}

func TestGetUser(t *testing.T) {
	client, ctx := setupDatabase(t)
	defer client.Close()

	_, err := client.User.
		Create().
		SetUsername("test123").
		SetAPIKey("secret123").
		Save(ctx)

	if err != nil {
		t.Fatalf("failed creating user: %v", err)
	}

	t.Run("Get user", func(t *testing.T) {

		us, err := GetUser(ctx, client, "secret123")
		if err != nil {
			t.Fatalf("failed getting user: %v", err)
		}

		usCount, err := client.User.Query().Count(ctx)
		if err != nil {
			t.Fatalf("failed counting prompts: %v", err)
		}

		assert.Equal(t, usCount, 1)
		assert.Equal(t, us.Username, "test123")
	})
}
