package prompt

import (
	"context"
	"fmt"
	"log"

	"github.com/mateusap1/promptq/ent"
)

func MakePromptRequest(ctx context.Context, client *ent.Client, prompt string) (*ent.PromptRequest, error) {
	pr, err := client.PromptRequest.
		Create().
		SetPrompt(prompt).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed creating prompt request: %w", err)
	}

	log.Println("prompt request was created: ", pr)

	return pr, nil
}

func QueuePromptRequest(ctx context.Context, client *ent.Client, promptRequest *ent.PromptRequest) (*ent.PromptRequest, error) {
	pr, err := promptRequest.Update().SetIsQueued(true).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed updating prompt request: %w", err)
	}

	return pr, nil
}
