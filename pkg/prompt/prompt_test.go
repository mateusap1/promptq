package prompt

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

func setupService(t *testing.T) *PromptService {
	ctx, client := setupDatabase(t)

	return CreateService(ctx, client)
}

func TestMakePromptRequest(t *testing.T) {
	ps := setupService(t)
	defer ps.client.Close()

	us, err := ps.client.User.
		Create().
		SetUsername("test123").
		SetPassword("secret123").
		Save(ps.ctx)

	if err != nil {
		t.Fatalf("failed creating user: %v", err)
	}

	t.Run("Create prompt", func(t *testing.T) {
		pr, err := ps.MakePromptRequest("Prompt #1", us)
		if err != nil {
			t.Fatalf("failed creating prompt: %v", err)
		}

		prCount, err := ps.client.PromptRequest.Query().Count(ps.ctx)
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
	ps := setupService(t)
	defer ps.client.Close()

	response, err := ps.HasQueuePromptRequest()
	assert.Nil(t, err)
	assert.False(t, response)

	_, err = ps.client.PromptRequest.
		Create().
		SetPrompt("Prompt #1").
		Save(ps.ctx)

	if err != nil {
		t.Fatalf("failed creating prompt request: %v", err)
	}

	response2, err2 := ps.HasQueuePromptRequest()
	assert.Nil(t, err2)
	assert.True(t, response2)
}

func TestQueuePromptRequest(t *testing.T) {
	ps := setupService(t)
	defer ps.client.Close()

	pr, err := ps.client.PromptRequest.
		Create().
		SetPrompt("Prompt #1").
		Save(ps.ctx)

	if err != nil {
		t.Fatalf("failed creating prompt request: %v", err)
	}

	t.Run("Queue prompt", func(t *testing.T) {
		pru, err := ps.QueuePromptRequest()

		assert.Nil(t, err)
		assert.Equal(t, pr.ID, pru.ID)
		assert.Equal(t, pru.IsQueued, true)
		assert.Equal(t, pru.IsAnswered, false)
	})
}

func TestAnswerPromptRequest(t *testing.T) {
	ps := setupService(t)
	defer ps.client.Close()

	pr, err := ps.client.PromptRequest.
		Create().
		SetPrompt("Prompt #1").
		Save(ps.ctx)

	if err != nil {
		t.Fatalf("failed creating prompt request: %v", err)
	}

	// Should not be able to answer a prompt request which is not queued
	_, err = ps.AnswerPromptRequest(pr)
	assert.NotNil(t, err)

	pr2, err := pr.Update().SetIsQueued(true).Save(ps.ctx)
	if err != nil {
		t.Fatalf("failed queueing prompt request: %v", err)
	}

	pr3, err := ps.AnswerPromptRequest(pr2)
	assert.Nil(t, err)
	assert.Equal(t, pr2.ID, pr3.ID)
	assert.True(t, pr3.IsAnswered)
}

func TestGetPromptRequests(t *testing.T) {
	ps := setupService(t)
	defer ps.client.Close()

	us, err := ps.client.User.
		Create().
		SetUsername("test123").
		SetPassword("secret123").
		Save(ps.ctx)

	if err != nil {
		t.Fatalf("failed creating user: %v", err)
	}

	_, err = ps.client.PromptRequest.
		Create().
		SetPrompt("Prompt #1").
		SetUser(us).
		Save(ps.ctx)

	if err != nil {
		t.Fatalf("failed creating prompt request: %v", err)
	}

	_, err = ps.client.PromptRequest.
		Create().
		SetPrompt("Prompt #2").
		SetUser(us).
		Save(ps.ctx)

	if err != nil {
		t.Fatalf("failed creating prompt request: %v", err)
	}

	t.Run("Get prompts", func(t *testing.T) {
		prs, err := ps.GetPromptRequests(us)

		assert.Len(t, prs, 2)
		assert.Nil(t, err)
	})
}

func TestMakePromptResponse(t *testing.T) {
	ps := setupService(t)
	defer ps.client.Close()

	t.Run("Create prompt response", func(t *testing.T) {
		prq, err := ps.client.PromptRequest.
			Create().
			SetPrompt("Prompt #1").
			Save(ps.ctx)
		if err != nil {
			t.Fatalf("failed creating prompt request: %v", err)
		}

		prp, err := ps.MakePromptResponse(prq, "Response #1")
		if err != nil {
			t.Fatal(err)
		}

		prpCount, err := ps.client.PromptResponse.Query().Count(ps.ctx)
		if err != nil {
			t.Fatalf("failed counting prompts: %v", err)
		}

		prq2, err := prp.QueryPromptRequest().Only(ps.ctx)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, prpCount, 1)
		assert.Equal(t, prp.Response, "Response #1")
		assert.Equal(t, prq2.ID, prq.ID)
	})
}
