package utils

import (
	"database/sql"
	"log"
)

func CreateMockUser(db *sql.DB, email string, passwordHash string, verified bool) (id int64) {
	const query = "INSERT INTO users (email, password_hash, email_verified) VALUES ($1, $2, $3) RETURNING id;"
	if err := db.QueryRow(query, email, passwordHash, verified).Scan(&id); err != nil {
		log.Fatal("Error inserting user: ", err)
	}

	return id
}

func CreateMockSession(db *sql.DB, userId int64, userAgent string, ipAddress string, token string, active bool) (id int64) {
	const query = "INSERT INTO sessions (user_id, user_agent, ip_address, session_token, active) VALUES ($1, $2, $3, $4, $5) RETURNING id;"
	if err := db.QueryRow(query, userId, userAgent, ipAddress, token, active).Scan(&id); err != nil {
		log.Fatal("Error inserting session: ", err)
	}

	return id
}
