package prompt

import (
	"context"
	"fmt"
	"log"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/mateusap1/promptq/ent"
	"github.com/mateusap1/promptq/ent/promptrequest"
)

type PromptService struct {
	ctx    context.Context
	client *ent.Client
}

func CreateService(ctx context.Context, client *ent.Client) *PromptService {
	return &PromptService{ctx, client}
}

func (ps *PromptService) MakePromptRequest(prompt string, user *ent.User) (*ent.PromptRequest, error) {
	ctx, client := ps.ctx, ps.client

	pr, err := client.PromptRequest.
		Create().
		SetPrompt(prompt).
		SetUser(user).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed creating prompt request: %w", err)
	}

	log.Println("prompt request was created: ", pr)

	return pr, nil
}

func (ps *PromptService) HasQueuePromptRequest() (bool, error) {
	ctx, client := ps.ctx, ps.client

	num, err := client.PromptRequest.
		Query().
		Where(promptrequest.And(promptrequest.IsQueued(false), promptrequest.IsAnswered(false))).
		Count(ctx)
	if err != nil {
		return false, fmt.Errorf("failed querying count of unqueued prompt requests: %w", err)
	}

	if num > 0 {
		return true, nil
	} else {
		return false, nil
	}
}

func (ps *PromptService) QueuePromptRequest() (*ent.PromptRequest, error) {
	ctx, client := ps.ctx, ps.client

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

func (ps *PromptService) AnswerPromptRequest(promptRequest *ent.PromptRequest) (*ent.PromptRequest, error) {
	ctx := ps.ctx

	if !promptRequest.IsQueued {
		return nil, fmt.Errorf("prompt request needs to be queued in order to be answered")
	}

	pr, err := promptRequest.Update().SetIsAnswered(true).SetIsQueued(false).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed updating prompt request: %w", err)
	}

	return pr, nil
}

func (ps *PromptService) GetPromptRequest(id string) (*ent.PromptRequest, error) {
	ctx, client := ps.ctx, ps.client

	pid, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("prompt id %v with wrong format, parse UUID failed with error: %w", id, err)
	}

	pr, err := client.PromptRequest.Query().Where(promptrequest.Identifier(pid)).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed getting latest prompt request: %w", err)
	}

	return pr, nil
}

func (ps *PromptService) GetPromptRequests(user *ent.User) ([]*ent.PromptRequest, error) {
	ctx := ps.ctx

	prs, err := user.QueryPromptRequests().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed getting latest prompt request: %w", err)
	}

	return prs, nil
}

func (ps *PromptService) GetPromptResponse(p *ent.PromptRequest) (*ent.PromptResponse, error) {
	ctx := ps.ctx

	return p.QueryPromptResponse().Only(ctx)
}

func (ps *PromptService) MakePromptResponse(prompt_request *ent.PromptRequest, response string) (*ent.PromptResponse, error) {
	ctx, client := ps.ctx, ps.client

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
