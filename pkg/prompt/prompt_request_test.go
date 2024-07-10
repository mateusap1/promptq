package prompt

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mateusap1/promptq/ent"
	_ "github.com/mattn/go-sqlite3"
)

func TestMakePromptRequest(t *testing.T) {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")

	if err != nil {
		t.Fatalf("failed opening connection to sqlite: %v", err)
	}

	defer client.Close()

	ctx := context.Background()

	if err := client.Schema.Create(ctx); err != nil {
		t.Fatalf("failed creating schema resources: %v", err)
	}

	t.Run("Create prompt", func(t *testing.T) {
		pr, err := MakePromptRequest(ctx, client, "Prompt #1")
		if err != nil {
			t.Fatalf("failed creating prompt: %v", err)
		}

		prCount, err := client.PromptRequest.Query().Count(ctx)
		if err != nil {
			t.Fatalf("failed counting prompts: %v", err)
		}

		assert.Equal(t, prCount, 1)
		assert.Equal(t, pr.Prompt, "Prompt #1")
		assert.Equal(t, pr.Queued, false)
	})
}

func TestQueuePromptRequest(t *testing.T) {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")

	if err != nil {
		t.Fatalf("failed opening connection to sqlite: %v", err)
	}

	defer client.Close()

	ctx := context.Background()

	if err := client.Schema.Create(ctx); err != nil {
		t.Fatalf("failed creating schema resources: %v", err)
	}

	pr, err := client.PromptRequest.
		Create().
		SetPrompt("Prompt #1").
		Save(ctx)

	if err != nil {
		t.Fatalf("failed creating prompt request: %v", err)
	}

	t.Run("Queue prompt", func(t *testing.T) {
		pr, err := QueuePromptRequest(ctx, client, pr)
		if err != nil {
			t.Fatalf("failed queueing prompt: %v", err)
		}

		assert.Equal(t, pr.Queued, true)
	})
}
