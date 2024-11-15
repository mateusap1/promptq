package utils

import (
	"database/sql"
	"log"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createMockUser(db *sql.DB, email string, passwordHash string, verified bool) (id int64) {
	const query = "INSERT INTO users (email, password_hash, email_verified) VALUES ($1, $2, $3);"
	result, err := db.Exec(query, email, passwordHash, verified)
	if err != nil {
		log.Fatal("Error inserting user: ", err)
	}

	id, err = result.LastInsertId()
	if err != nil {
		log.Fatal("Error getting last inserted id: ", err)
	}

	return id
}

func createMockSession(db *sql.DB, userId int64, userAgent string, ipAddress string, token string, active bool) (id int64) {
	const query = "INSERT INTO sessions (user_id, user_agent, ip_address, session_token, active) VALUES ($1, $2, $3, $4, $5);"
	result, err := db.Exec(query, userId, userAgent, ipAddress, token, active)
	if err != nil {
		log.Fatal("Error inserting session: ", err)
	}

	id, err = result.LastInsertId()
	if err != nil {
		log.Fatal("Error getting last inserted id: ", err)
	}

	return id
}

func TestValidEmailFormat(t *testing.T) {
	assert.Equal(t, true, ValidEmailFormat("a@b"))
	assert.Equal(t, false, ValidEmailFormat("@ba"))
	assert.Equal(t, false, ValidEmailFormat("ab@"))
	assert.Equal(t, false, ValidEmailFormat("a@"+strings.Repeat("a", 256)))
}

func TestValidPasswordFormat(t *testing.T) {
	assert.Equal(t, true, ValidPasswordFormat("P@ssw0rd"))
	assert.Equal(t, true, ValidPasswordFormat("Pass1-_*)!"))
	assert.Equal(t, false, ValidPasswordFormat("password"))
	assert.Equal(t, false, ValidPasswordFormat("pWd0!"))
	assert.Equal(t, false, ValidPasswordFormat("pAssw0rd"))
	assert.Equal(t, false, ValidPasswordFormat("12345a678@"))
	assert.Equal(t, false, ValidPasswordFormat("12345A678!"))
	assert.Equal(t, false, ValidPasswordFormat(strings.Repeat("a", 64)+"A@0"))
}

func TestEmailAlreadyExists(t *testing.T) {
	db := setup()

	if _, err := db.Exec(`
		INSERT INTO users (email, password_hash) VALUES ($1, $2);
	`, "alice@email.com", ""); err != nil {
		log.Fatal("Error inserting user:", err)
	}

	exists, err := EmailAlreadyExists(db, "bob@email.com")
	assert.Nil(t, err)
	assert.Equal(t, false, exists)

	exists, err = EmailAlreadyExists(db, "alice@email.com")
	assert.Nil(t, err)
	assert.Equal(t, true, exists)
}

func TestCreateUser(t *testing.T) {
	db := setup()

	expectedToken, err := CreateUser(db, "alice@email.com", "")
	assert.Nil(t, err)

	var token string
	err = db.QueryRow("SELECT validate_token FROM users WHERE email=$1;", "alice@email.com").Scan(&token)
	assert.Nil(t, err, "user was not created")

	assert.Equal(t, token, expectedToken, "token returned is different")
}

func TestGetUserLoginByEmail(t *testing.T) {
	db := setup()

	createMockUser(db, "alice@email.com", "pw", true)

	_, _, _, err := GetUserLoginByEmail(db, "bob@email.com")
	assert.ErrorIs(t, sql.ErrNoRows, err)

	_, passwordHash, emailVerified, err := GetUserLoginByEmail(db, "alice@email.com")
	assert.Nil(t, err)
	assert.Equal(t, passwordHash, "pw")
	assert.Equal(t, emailVerified, true)
}

func TestGetActiveSession(t *testing.T) {
	db := setup()

	var err error

	aliceId := createMockUser(db, "alice@email.com", "pw", true)
	activeSessionId := createMockSession(db, aliceId, "agent1", "ip", "token", true)

	id, token, err := GetActiveSession(db, aliceId)
	assert.Nil(t, err)
	assert.Equal(t, "token", token)
	assert.Equal(t, activeSessionId, id)

	// Test with not active

	// Creating a new session not active. Should not return it
	createMockSession(db, aliceId, "agent2", "ip2", "token2", false)

	id, token, err = GetActiveSession(db, aliceId)
	assert.Nil(t, err)
	assert.Equal(t, "token", token)
	assert.Equal(t, activeSessionId, id)

	// If user has no sessions, should return error
	bobId := createMockUser(db, "bob@email.com", "pw", true)

	_, _, err = GetActiveSession(db, bobId)
	assert.Equal(t, sql.ErrNoRows, err)
}

func TestCreateSession(t *testing.T) {
	db := setup()

	aliceId := createMockUser(db, "alice@email.com", "pw", true)

	sessionId, sessionToken, err := CreateSession(db, aliceId, "ua", "ip")
	assert.Nil(t, err)

	var userId int64
	var userAgent, ipAddress, dbToken string
	if err := db.QueryRow("SELECT user_id, user_agent, ip_address, session_token FROM sessions WHERE id=$1;", sessionId).Scan(&userId, &userAgent, &ipAddress, &dbToken); err != nil {
		if err == sql.ErrNoRows {
			assert.Fail(t, "CreateSession did not insert session in db with returned id")
		} else {
			log.Fatal(err)
		}
	}
	assert.Nil(t, err)
	assert.Equal(t, aliceId, userId)
	assert.Equal(t, "ua", userAgent)
	assert.Equal(t, "ip", ipAddress)
	assert.Equal(t, sessionToken, dbToken)
}

func TestDeactivateSession(t *testing.T) {
	db := setup()

	userId := createMockUser(db, "alice@email.com", "pwd", true)
	sessionId := createMockSession(db, userId, "ua", "ip", "token", true)

	DeactivateSession(db, sessionId)
	var active bool
	if err := db.QueryRow("SELECT active FROM sessions WHERE id=$1;", sessionId).Scan(&active); err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, active, false)
}

func setup() (db *sql.DB) {
	db = OpenSQLite(":memory:")

	if _, err := db.Exec(`
		CREATE TABLE users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email VARCHAR NOT NULL,
			password_hash VARCHAR NOT NULL,
			email_verified BOOLEAN DEFAULT FALSE NOT NULL,
			validate_token VARCHAR NULL,
			validate_token_expires TIMESTAMP NULL,
			reset_token VARCHAR NULL,
			reset_token_expires TIMESTAMP NULL,
			created_at TIMESTAMP DEFAULT NOW,
			updated_at TIMESTAMP
		);

		CREATE TABLE sessions (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER NOT NULL REFERENCES users ON DELETE CASCADE,
			user_agent VARCHAR NOT NULL,
			ip_address VARCHAR NOT NULL,
			session_token VARCHAR NOT NULL,
			active BOOLEAN DEFAULT TRUE NOT NULL,
			created_at TIMESTAMP DEFAULT NOW,
			expires_at TIMESTAMP
		);
    `); err != nil {
		log.Fatal("Error creating tables: ", err)
	}

	return db
}
