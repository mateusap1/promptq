package utils

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

func ValidateEmail(email string) bool {
	atIdx := strings.IndexRune(email, '@')
	if atIdx == -1 || atIdx == 0 || atIdx == len(email)-1 {
		return false
	}

	if len(email) > 256 {
		return false
	}

	return true
}

func ValidatePassword(password string) bool {
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

// func EmailAlreadyExists(db *sql.DB, email string) (bool, bool) {
// 	// Returns (exists, expired)

// 	rows, err := db.Query("SELECT id, email_confirmed, confirm_token_expires FROM user WHERE email=$1;", email)
// 	if err != nil {
// 		log.Fatal("Error querying data:", err)
// 	}
// 	defer rows.Close()

// 	// var users [](int, string)
// 	for rows.Next() {
// 		var id int
// 		var name string
// 		if err := rows.Scan(&id, &name); err != nil {
// 			log.Fatal(err)
// 		}
// 		fmt.Printf("ID: %d, Name: %s\n", id, name)

// 		counter++
// 	}

// 	if err := rows.Err(); err != nil {
// 		log.Fatal(err)
// 	}
// }
