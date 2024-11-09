package user

import (
	"github.com/mateusap1/promptq/ent"
	"golang.org/x/crypto/argon2"
)

type UserData struct {
	Username string
	Password []byte
	Salt     []byte
}

func HashPassword(rawPassword string, salt []byte) []byte {
	return argon2.IDKey([]byte(rawPassword), salt, 1, 64*1024, 4, 32)
}

type ModelInterface interface {
	CreateUser(UserData) error
	GetUser(int) (*ent.User, error)
}

type RegisterData struct {
	Email    string
	Password string
}

type ControllerInterface interface {
	Register(string, string) error
	GetUser(int) (*ent.User, error)
}
