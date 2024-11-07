package user

import (
	"context"
	"fmt"

	"github.com/mateusap1/promptq/ent"
	"github.com/mateusap1/promptq/ent/user"
)

// Maybe change this to interface in the future
type UserService struct {
	ctx    context.Context
	client *ent.Client
}

func CreateService(ctx context.Context, client *ent.Client) *UserService {
	return &UserService{ctx, client}
}

func (s *UserService) MakeUser(username string, password string) (*ent.User, error) {
	ctx, client := s.ctx, s.client

	user, err := client.User.
		Create().
		SetUsername(username).
		SetPassword(password).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}

	return user, nil
}

func (s *UserService) GetUser(id int) (*ent.User, error) {
	ctx, client := s.ctx, s.client

	us, err := client.User.Query().Where(user.ID(id)).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed getting latest user: %w", err)
	}

	return us, nil
}
