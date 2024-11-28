package utils

import "database/sql"

func SendMessage(db *sql.DB, threadId int64, content string, ai bool) error {
	const query = "INSERT INTO prompts (thread_id, ai, content) VALUES ($1, $2, $3);"
	if _, err := db.Exec(query, threadId, ai, content); err != nil {
		return err
	}

	return nil
}
