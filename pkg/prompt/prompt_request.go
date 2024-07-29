package prompt

import (
	"context"
	"fmt"
	"log"

	"entgo.io/ent/dialect/sql"
	"github.com/mateusap1/promptq/ent"
	"github.com/mateusap1/promptq/ent/promptrequest"
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

func HasQueuePromptRequest(ctx context.Context, client *ent.Client) (bool, error) {
	return false, nil
}

func QueuePromptRequest(ctx context.Context, client *ent.Client) (*ent.PromptRequest, error) {
	pr, err := client.PromptRequest.
		Query().
		Where(promptrequest.And(promptrequest.IsQueued(false), promptrequest.IsAnswered(false))).Order(promptrequest.ByCreateDate(sql.OrderAsc())).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed getting latest prompt request: %w", err)
	}

	pru, err := pr.Update().SetIsQueued(true).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed updating prompt request: %w", err)
	}

	return pru, nil
}

func AnswerPromptRequest(ctx context.Context, client *ent.Client, promptRequest *ent.PromptRequest) (*ent.PromptRequest, error) {
	if !promptRequest.IsQueued {
		return nil, fmt.Errorf("prompt request needs to be queued in order to be answered")
	}

	pr, err := promptRequest.Update().SetIsAnswered(true).SetIsQueued(false).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed updating prompt request: %w", err)
	}

	return pr, nil
}
