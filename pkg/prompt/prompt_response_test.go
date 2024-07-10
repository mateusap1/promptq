package prompt

import (
	"context"
	"testing"

	"github.com/mateusap1/promptq/ent"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
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
		prq, err := client.PromptRequest.
			Create().
			SetPrompt("Prompt #1").
			Save(ctx)
		if err != nil {
			t.Fatalf("failed creating prompt request: %v", err)
		}

		prp, err := MakePromptResponse(ctx, client, prq, "Response #1")
		if err != nil {
			t.Fatal(err)
		}

		prpCount, err := client.PromptResponse.Query().Count(ctx)
		if err != nil {
			t.Fatalf("failed counting prompts: %v", err)
		}

		prq2, err := prp.QueryPromptRequest().Only(ctx)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, prpCount, 1)
		assert.Equal(t, prp.Response, "Response #1")
		assert.Equal(t, prq2.ID, prq.ID)
	})
}
