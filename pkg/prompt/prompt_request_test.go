package prompt

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

func TestMakePromptRequest(t *testing.T) {
	client, ctx := setupDatabase(t)
	defer client.Close()

	us, err := client.User.
		Create().
		SetUsername("test123").
		SetAPIKey("secret123").
		Save(ctx)

	if err != nil {
		t.Fatalf("failed creating user: %v", err)
	}

	t.Run("Create prompt", func(t *testing.T) {
		pr, err := MakePromptRequest(ctx, client, "Prompt #1", us)
		if err != nil {
			t.Fatalf("failed creating prompt: %v", err)
		}

		prCount, err := client.PromptRequest.Query().Count(ctx)
		if err != nil {
			t.Fatalf("failed counting prompts: %v", err)
		}

		assert.Equal(t, prCount, 1)
		assert.Equal(t, pr.Prompt, "Prompt #1")
		assert.Equal(t, pr.IsQueued, false)
		assert.Equal(t, pr.IsAnswered, false)
	})
}

func TestHasQueuePromptRequest(t *testing.T) {
	client, ctx := setupDatabase(t)
	defer client.Close()

	response, err := HasQueuePromptRequest(ctx, client)
	assert.Nil(t, err)
	assert.False(t, response)

	_, err = client.PromptRequest.
		Create().
		SetPrompt("Prompt #1").
		Save(ctx)

	if err != nil {
		t.Fatalf("failed creating prompt request: %v", err)
	}

	response2, err2 := HasQueuePromptRequest(ctx, client)
	assert.Nil(t, err2)
	assert.True(t, response2)
}

func TestQueuePromptRequest(t *testing.T) {
	client, ctx := setupDatabase(t)
	defer client.Close()

	pr, err := client.PromptRequest.
		Create().
		SetPrompt("Prompt #1").
		Save(ctx)

	if err != nil {
		t.Fatalf("failed creating prompt request: %v", err)
	}

	t.Run("Queue prompt", func(t *testing.T) {
		pru, err := QueuePromptRequest(ctx, client)

		assert.Nil(t, err)
		assert.Equal(t, pr.ID, pru.ID)
		assert.Equal(t, pru.IsQueued, true)
		assert.Equal(t, pru.IsAnswered, false)
	})
}

func TestAnswerPromptRequest(t *testing.T) {
	client, ctx := setupDatabase(t)
	defer client.Close()

	pr, err := client.PromptRequest.
		Create().
		SetPrompt("Prompt #1").
		Save(ctx)

	if err != nil {
		t.Fatalf("failed creating prompt request: %v", err)
	}

	// Should not be able to answer a prompt request which is not queued
	_, err = AnswerPromptRequest(ctx, client, pr)
	assert.NotNil(t, err)

	pr2, err := pr.Update().SetIsQueued(true).Save(ctx)
	if err != nil {
		t.Fatalf("failed queueing prompt request: %v", err)
	}

	pr3, err := AnswerPromptRequest(ctx, client, pr2)
	assert.Nil(t, err)
	assert.Equal(t, pr2.ID, pr3.ID)
	assert.True(t, pr3.IsAnswered)
}
