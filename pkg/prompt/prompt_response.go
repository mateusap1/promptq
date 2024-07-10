package prompt

import (
	"context"
	"fmt"
	"log"

	"github.com/mateusap1/promptq/ent"
)

func MakePromptResponse(ctx context.Context, client *ent.Client, response string) (*ent.PromptResponse, error) {
	u, err := client.PromptResponse.
		Create().
		SetResponse(response).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed creating prompt request: %w", err)
	}

	log.Println("prompt request was created: ", u)

	return u, nil
}
