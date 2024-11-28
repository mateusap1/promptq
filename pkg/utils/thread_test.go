package utils

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateThread(t *testing.T) {
	db := setup()

	userId := CreateMockUser(db, "alice@email.com", "", true)
	id, err := CreateThread(db, userId, "test")
	assert.Nil(t, err)

	var actualUserId int64
	var actualTName string
	if err := db.QueryRow("SELECT user_id, tname FROM threads WHERE id=$1;", id).Scan(&actualUserId, &actualTName); err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, actualUserId, userId)
	assert.Equal(t, actualTName, "test")
}
