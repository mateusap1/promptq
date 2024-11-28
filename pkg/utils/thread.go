package utils

import (
	"database/sql"

	"github.com/google/uuid"
)

func GetThread(db *sql.DB, userId int64, tid string) (id int64, name string, deleted bool, err error) {
	const query = "SELECT id, tname, deleted FROM threads WHERE user_id=$1 AND tid=$2;"
	if err = db.QueryRow(query, userId, tid).Scan(&id, &name, &deleted); err != nil {
		return -1, "", false, err
	}

	return id, name, deleted, nil
}

func CreateThread(db *sql.DB, userId int64, name string) (id int64, err error) {
	tid := uuid.New().String()
	const query = "INSERT INTO threads (user_id, tid, tname) VALUES ($1, $2, $3) RETURNING id;"
	if err = db.QueryRow(query, userId, tid, name).Scan(&id); err != nil {
		return -1, err
	}

	return id, nil
}
