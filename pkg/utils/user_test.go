package utils

import (
	"database/sql"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var db *sql.DB

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
	validateToken, err := CreateUser(db, "charlie@email.com", "")
	assert.Nil(t, err)

	var validateToken2 string
	err = db.QueryRow("SELECT validate_token FROM users WHERE email=$1;", "charlie@email.com").Scan(&validateToken2)
	assert.Nil(t, err, "user was not created")

	assert.Equal(t, validateToken2, validateToken, "token returned is different")
}

func setup() {
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
			updated_at TIMESTAMP DEFAULT NOW
		);
    `); err != nil {
		log.Fatal("Error creating table:", err)
	}
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	// shutdown()
	os.Exit(code)
}
