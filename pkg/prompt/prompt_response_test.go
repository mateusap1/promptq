package prompt

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mateusap1/promptq/ent"
	_ "github.com/mattn/go-sqlite3"
)

func TestMakePromptResponse(t *testing.T) {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")

	if err != nil {
		t.Fatalf("failed opening connection to sqlite: %v", err)
	}

	defer client.Close()

	ctx := context.Background()

	if err := client.Schema.Create(ctx); err != nil {
		t.Fatalf("failed creating schema resources: %v", err)
	}

	t.Run("Create prompt response", func(t *testing.T) {
		pr, err := MakePromptResponse(ctx, client, "Response #1")
		if err != nil {
			t.Fatalf("failed creating prompt response: %v", err)
		}

		prCount, err := client.PromptResponse.Query().Count(ctx)
		if err != nil {
			t.Fatalf("failed counting prompts: %v", err)
		}

		assert.Equal(t, prCount, 1)
		assert.Equal(t, pr.Response, "Response #1")
		assert.Equal(t, pr.IsAnswered, false)
	})
}
