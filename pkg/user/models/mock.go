package model

import (
	"fmt"

	usr "github.com/mateusap1/promptq/pkg/user"
)

// Maybe change this to interface in the future
type MockModel struct {
	users []usr.UserData
}

func (m *MockModel) CreateUser(user usr.UserData) error {
	m.users = append(m.users, user)
	return nil
}

func (m *MockModel) GetUser(id int) (*usr.UserData, error) {
	if id >= len(m.users) {
		return nil, fmt.Errorf("user does not exist with id %v", id)
	}

	return &m.users[id], nil
}
