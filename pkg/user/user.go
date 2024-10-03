package user

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"

	"github.com/mateusap1/promptq/ent"
	"github.com/mateusap1/promptq/ent/user"
)

func randomStringCrypto(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(bytes)[:length], nil
}

func MakeUser(ctx context.Context, client *ent.Client, username string) (*ent.User, error) {
	apiKey, err := randomStringCrypto(64)
	if err != nil {
		return nil, fmt.Errorf("failed generating api key: %w", err)
	}

	user, err := client.User.
		Create().
		SetUsername(username).
		SetAPIKey(apiKey).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}

	log.Println("user was created: ", user)

	return user, nil
}

func GetUser(ctx context.Context, client *ent.Client, api_key string) (*ent.User, error) {
	us, err := client.User.Query().Where(user.APIKeyEQ(api_key)).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed getting latest user: %w", err)
	}

	return us, nil
}
