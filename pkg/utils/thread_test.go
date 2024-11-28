package utils

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetThread(t *testing.T) {
	db := setup()

	userId := CreateMockUser(db, "alice@email.com", "", true)
	id := CreateMockThread(db, userId, "test_id", "test_name", false)

	threadId, tname, deleted, err := GetThread(db, userId, "test_id")
	assert.Nil(t, err)

	assert.Equal(t, threadId, id)
	assert.Equal(t, tname, "test_name")
	assert.Equal(t, deleted, false)
}

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
