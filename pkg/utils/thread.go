package utils

import (
	"database/sql"

	"github.com/google/uuid"
)

func CreateThread(db *sql.DB, userId int64, name string) (id int64, err error) {
	tid := uuid.New().String()
	const query = "INSERT INTO threads (user_id, tid, tname) VALUES ($1, $2, $3) RETURNING id;"
	if err = db.QueryRow(query, userId, tid, name).Scan(&id); err != nil {
		return -1, err
	}

	return id, nil
}
