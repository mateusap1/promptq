package model

import (
	"context"
	"fmt"

	"github.com/mateusap1/promptq/ent"
	"github.com/mateusap1/promptq/ent/user"
	usr "github.com/mateusap1/promptq/pkg/user"
)

type DBModel struct {
	ctx    context.Context
	client *ent.Client
}

func CreateDbModel(ctx context.Context, client *ent.Client) *DBModel {
	return &DBModel{ctx, client}
}

func (s *DBModel) CreateUser(user usr.UserData) error {
	ctx, client := s.ctx, s.client

	_, err := client.User.
		Create().
		SetUsername(user.Username).
		SetPassword(user.Password).
		SetSalt(user.Salt).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed creating user: %w", err)
	}

	return nil
}

func (s *DBModel) GetUser(id int) (*ent.User, error) {
	ctx, client := s.ctx, s.client

	us, err := client.User.Query().Where(user.ID(id)).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed getting latest user: %w", err)
	}

	return us, nil
}
