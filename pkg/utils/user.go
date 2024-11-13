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

	rows, err := db.Query("SELECT id FROM users WHERE email=$1;", email)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	if rows.Next() {
		return true, nil
	} else {
		if err := rows.Err(); err != nil {
			return false, err
		}

		return false, nil
	}
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
