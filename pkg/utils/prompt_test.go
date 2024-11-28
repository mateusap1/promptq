package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendMessage(t *testing.T) {
	db := setup()

	userId := CreateMockUser(db, "alice@email.com", "", true)
	threadId := CreateMockThread(db, userId, "tid", "tname", false)

	err := SendMessage(db, threadId, "Hello there!", false)
	assert.Nil(t, err)

	var (
		content string
		ai      bool
	)
	err = db.QueryRow("SELECT content, ai FROM prompts;").Scan(&content, &ai)
	assert.Nil(t, err)
	assert.Equal(t, "Hello there!", content)
	assert.Equal(t, false, ai)
}
