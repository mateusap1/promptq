package utils

import (
	"database/sql"

	"github.com/google/uuid"
)

type Thread struct {
	Identifier string `json:"id"`
	Name       string `json:"name"`
}

func GetThread(db *sql.DB, tid string) (id int64, name string, deleted bool, err error) {
	const query = "SELECT id, tname, deleted FROM threads WHERE tid=$1;"
	if err = db.QueryRow(query, tid).Scan(&id, &name, &deleted); err != nil {
		return -1, "", false, err
	}

	return id, name, deleted, nil
}

func GetThreads(db *sql.DB, userId int64) (threads []Thread, err error) {
	const query = "SELECT tid, tname FROM threads WHERE user_id=$1 AND deleted=false ORDER BY updated_at DESC, created_at DESC;"
	rows, err := db.Query(query, userId)
	if err != nil {
		return []Thread{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			tid  string
			name string
		)
		if err := rows.Scan(&tid, &name); err != nil {
			return []Thread{}, err
		}

		threads = append(threads, Thread{tid, name})
	}

	return threads, nil
}

func CreateThread(db *sql.DB, userId int64, name string) (id int64, tid string, err error) {
	tid = uuid.New().String()
	const query = "INSERT INTO threads (user_id, tid, tname) VALUES ($1, $2, $3) RETURNING id;"
	if err = db.QueryRow(query, userId, tid, name).Scan(&id); err != nil {
		return -1, "", err
	}

	return id, tid, nil
}

func RenameThread(db *sql.DB, threadId int64, name string) error {
	const query = "UPDATE threads SET tname=$1, updated_at=CURRENT_TIMESTAMP WHERE id=$2;"
	_, err := db.Exec(query, name, threadId)

	return err
}

func DeleteThread(db *sql.DB, threadId int64) error {
	const query = "UPDATE threads SET deleted=true, updated_at=CURRENT_TIMESTAMP WHERE id=$1;"
	_, err := db.Exec(query, threadId)

	return err
}
