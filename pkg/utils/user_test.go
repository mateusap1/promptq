package utils

import (
	"database/sql"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmailAlreadyExists(t *testing.T) {
	db := setup()

	CreateMockUser(db, "alice@email.com", "", false)

	exists, err := EmailAlreadyExists(db, "alice@email.com")
	assert.Nil(t, err)
	assert.Equal(t, true, exists)

	exists, err = EmailAlreadyExists(db, "bob@email.com")
	assert.Nil(t, err)
	assert.Equal(t, false, exists)
}

func TestCreateUser(t *testing.T) {
	db := setup()

	userId, expectedToken, err := CreateUser(db, "alice@email.com", "")
	assert.Nil(t, err)

	var token string
	var expectedId int64
	err = db.QueryRow("SELECT id, validate_token FROM users WHERE email=$1;", "alice@email.com").Scan(&expectedId, &token)
	assert.Nil(t, err, "user was not created")

	assert.Equal(t, expectedId, userId)
	assert.Equal(t, token, expectedToken)
}

func TestGetUserLoginByEmail(t *testing.T) {
	db := setup()

	CreateMockUser(db, "alice@email.com", "pw", true)

	_, _, err := GetUserLoginByEmail(db, "bob@email.com")
	assert.ErrorIs(t, sql.ErrNoRows, err)

	_, passwordHash, err := GetUserLoginByEmail(db, "alice@email.com")
	assert.Nil(t, err)
	assert.Equal(t, passwordHash, "pw")
}

func TestSessionByToken(t *testing.T) {
	db := setup()

	var err error

	aliceId := CreateMockUser(db, "alice@email.com", "pw", true)
	sessionId := CreateMockSession(db, aliceId, "agent1", "ip", "token", true)

	id, userId, active, _, err := GetSessionByToken(db, "token")
	assert.Nil(t, err)
	assert.Equal(t, userId, aliceId)
	assert.Equal(t, true, active)
	assert.Equal(t, sessionId, id)
}

func TestGetActiveSession(t *testing.T) {
	db := setup()

	var err error

	aliceId := CreateMockUser(db, "alice@email.com", "pw", true)
	activeSessionId := CreateMockSession(db, aliceId, "agent1", "ip", "token", true)

	id, token, err := GetActiveSession(db, aliceId)
	assert.Nil(t, err)
	assert.Equal(t, "token", token)
	assert.Equal(t, activeSessionId, id)

	// Test with not active

	// Creating a new session not active. Should not return it
	CreateMockSession(db, aliceId, "agent2", "ip2", "token2", false)

	id, token, err = GetActiveSession(db, aliceId)
	assert.Nil(t, err)
	assert.Equal(t, "token", token)
	assert.Equal(t, activeSessionId, id)

	// If user has no sessions, should return error
	bobId := CreateMockUser(db, "bob@email.com", "pw", true)

	_, _, err = GetActiveSession(db, bobId)
	assert.Equal(t, sql.ErrNoRows, err)
}

func TestCreateSession(t *testing.T) {
	db := setup()

	aliceId := CreateMockUser(db, "alice@email.com", "pw", true)

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

	userId := CreateMockUser(db, "alice@email.com", "pwd", true)
	sessionId := CreateMockSession(db, userId, "ua", "ip", "token", true)

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
