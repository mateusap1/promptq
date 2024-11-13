package utils

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func OpenDB(dbUrl string) *sql.DB {
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
