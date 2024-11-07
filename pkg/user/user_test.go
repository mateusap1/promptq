package user

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mateusap1/promptq/ent"
	_ "github.com/mattn/go-sqlite3"
)

func setupDatabase(t *testing.T) (context.Context, *ent.Client) {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		t.Fatalf("failed opening connection to sqlite: %v", err)
	}

	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		t.Fatalf("failed creating schema resources: %v", err)
	}

	return ctx, client
}

func setupService(t *testing.T) (s *UserService) {
	ctx, client := setupDatabase(t)

	return CreateService(ctx, client)
}

func TestMakeUser(t *testing.T) {
	us := setupService(t)
	defer us.client.Close()

	t.Run("Create user", func(t *testing.T) {
		user, err := us.MakeUser("test123", "secret123")
		if err != nil {
			t.Fatalf("failed creating user: %v", err)
		}

		userCount, err := us.client.User.Query().Count(us.ctx)
		if err != nil {
			t.Fatalf("failed counting prompts: %v", err)
		}

		assert.Equal(t, userCount, 1)
		assert.Equal(t, user.Username, "test123")
		assert.Equal(t, user.Password, "secret123")
	})
}

func TestGetUser(t *testing.T) {
	us := setupService(t)
	defer us.client.Close()

	user, err := us.client.User.
		Create().
		SetUsername("test123").
		SetPassword("secret123").
		Save(us.ctx)

	if err != nil {
		t.Fatalf("failed creating user: %v", err)
	}

	t.Run("Get user", func(t *testing.T) {

		user, err := us.GetUser(user.ID)
		if err != nil {
			t.Fatalf("failed getting user: %v", err)
		}

		userCount, err := us.client.User.Query().Count(us.ctx)
		if err != nil {
			t.Fatalf("failed counting prompts: %v", err)
		}

		assert.Equal(t, userCount, 1)
		assert.Equal(t, user.Username, "test123")
	})
}
