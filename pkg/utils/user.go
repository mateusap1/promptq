package utils

import (
	"database/sql"
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"
)

func ValidEmailFormat(email string) bool {
	atIdx := strings.IndexRune(email, '@')
	if atIdx == -1 || atIdx == 0 || atIdx == len(email)-1 {
		return false
	}

	if len(email) > 256 {
		return false
	}

	return true
}

func ValidPasswordFormat(password string) bool {
	if len(password) > 64 || len(password) < 8 {
		return false
	}

	specials := `!@#$%^&*()_+\-=\[\]{};:'",.<>?~` + "`"

	// Allowed characters
	regex, err := regexp.Compile(fmt.Sprintf(`^[a-zA-Z\d%v]*$`, specials))
	if err != nil {
		log.Fatal(err)
	}
	if !regex.MatchString(password) {
		return false
	}

	// Required characters (at least once)
	conditions := []string{
		"abcdefghijklmnopqrstuvwxyz",
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"0123456789",
		specials,
	}
	for i := 0; i < len(conditions); i++ {
		if !strings.ContainsAny(password, conditions[i]) {
			return false
		}
	}

	return true
}

func EmailAlreadyExists(db *sql.DB, email string) (bool, error) {
	// Need to handle case where email exists but has not been confirmed
	// Not handling it right now

	var id int
	if err := db.QueryRow("SELECT id FROM users WHERE email=$1;", email).Scan(&id); err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		} else {
			return false, err
		}
	}

	return true, nil
}

func CreateUser(db *sql.DB, email, passwordHash string) (validateToken string, err error) {
	validateToken, err = GenerateValidateToken()
	if err != nil {
		return "", err
	}

	confirmDuration := 24 * time.Hour
	currentTime := time.Now().UTC()

	query := "INSERT INTO users (email, password_hash, validate_token, validate_token_expires) VALUES ($1, $2, $3, $4)"
	if _, err := db.Exec(query, email, passwordHash, validateToken, currentTime.Add(confirmDuration)); err != nil {
		return "", err
	}

	return validateToken, nil
}

func SendValidationEmail(confirmToken string) error {
	log.Printf("Sending validation e-mail with token %v...\n", confirmToken)

	return nil
}

func GetUserLoginByEmail(db *sql.DB, email string) (id int64, passwordHash string, emailVerified bool, err error) {
	if err = db.QueryRow("SELECT id, password_hash, email_verified FROM users WHERE email=$1;", email).Scan(&id, &passwordHash, &emailVerified); err != nil {
		// error sql.ErrNoRows means that the user does not exist
		return -1, "", false, err
	}

	return id, passwordHash, emailVerified, nil
}

func GetActiveSession(db *sql.DB, userId int64) (id int64, token string, err error) {
	const query = "SELECT id, session_token FROM sessions WHERE user_id=$1 AND active=TRUE ORDER BY expires_at DESC;"
	if err = db.QueryRow(query, userId).Scan(&id, &token); err != nil {
		return -1, "", err
	}

	return id, token, nil
}

func CreateSession(db *sql.DB, userId int64, userAgent string, ipAddress string) (token string, err error) {
	token, err = GenerateValidateToken()
	if err != nil {
		return "", err
	}

	sessionDuration := 24 * time.Hour
	currentTime := time.Now().UTC()

	const query = "INSERT INTO sessions (user_id, user_agent, ip_address, session_token, expires_at) VALUES ($1, $2, $3, $4);"
	if _, err = db.Exec(query, userId, userAgent, ipAddress, token, currentTime.Add(sessionDuration)); err != nil {
		return "", err
	}

	return token, nil
}

func DeactivateSessionById(db *sql.DB, id int64) error {
	const query = "UPDATE sessions SET active=FALSE WHERE id=$1;"
	_, err := db.Exec(query, id)
	return err
}

func DeactivateSessionByToken(db *sql.DB, token string) error {
	const query = "UPDATE sessions SET active=FALSE WHERE session_token=$1;"
	_, err := db.Exec(query, token)
	return err
}
