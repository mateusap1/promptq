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

type UserServiceInterface interface {
	CreateUser(string, string) error
	GetUser(int) (*ent.User, error)
}
