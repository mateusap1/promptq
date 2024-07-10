package prompt

import (
	"context"
	"fmt"
	"log"

	"github.com/mateusap1/promptq/ent"
)

func MakePromptResponse(ctx context.Context, client *ent.Client, prompt_request *ent.PromptRequest, response string) (*ent.PromptResponse, error) {
	pr, err := client.PromptResponse.
		Create().
		SetResponse(response).
		SetPromptRequest(prompt_request).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed creating prompt response: %w", err)
	}

	log.Println("prompt response was created: ", pr)

	return pr, nil
}
