package api

import (
	"database/sql"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mateusap1/promptq/pkg/utils"
)

var (
	ErrEmailExists           = "email taken"
	ErrEmailUnverified       = "email unverified"
	ErrNoAccountEmail        = "no account with email"
	ErrWrongPassword         = "wrong password"
	ErrInvalidPasswordFormat = "invalid password format or weak password"
	ErrInvalidEmailFormat    = "invalid email format"
	ErrInvalidFormat         = "invalid request format"
)

type SignForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SignUp(c *gin.Context, db *sql.DB) {
	var form SignForm
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": ErrInvalidFormat, "error": "ErrInvalidFormat"})
		return
	}

	email := strings.ToLower(form.Email)
	password := form.Password

	if !utils.ValidEmailFormat(email) {
		c.JSON(http.StatusBadRequest, gin.H{"message": ErrInvalidEmailFormat, "error": "ErrInvalidEmailFormat"})
		return
	}

	if !utils.ValidPasswordFormat(password) {
		c.JSON(http.StatusBadRequest, gin.H{"message": ErrInvalidPasswordFormat, "error": "ErrInvalidPasswordFormat"})
		return
	}

	exists, err := utils.EmailAlreadyExists(db, email)
	if err != nil {
		log.Fatal(err)
		return
	}

	if exists {
		c.JSON(http.StatusBadRequest, gin.H{"message": ErrEmailExists, "error": "ErrEmailExists"})
		return
	}

	passwordHash := utils.EncodePassword(password)
	confirmToken, err := utils.CreateUser(db, email, passwordHash)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Send validation e-mail
	if err := utils.SendValidationEmail(confirmToken); err != nil {
		log.Fatal(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func SignIn(c *gin.Context, db *sql.DB) {
	var form SignForm
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": ErrInvalidFormat, "error": "ErrInvalidFormat"})
		return
	}

	email := strings.ToLower(form.Email)
	password := form.Password

	id, passwordHash, emailVerified, err := utils.GetUserLoginByEmail(db, email)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusBadRequest, gin.H{"message": ErrNoAccountEmail, "error": "ErrNoAccountEmail"})
			return
		} else {
			log.Fatal(err)
			return
		}
	}

	if !emailVerified {
		c.JSON(http.StatusBadRequest, gin.H{"message": ErrEmailUnverified, "error": "ErrEmailUnverified"})
		return
	}

	rightPassword, err := utils.ComparePasswordAndHash(password, passwordHash)
	if err != nil {
		log.Fatal(err)
		return
	}

	if !rightPassword {
		c.JSON(http.StatusBadRequest, gin.H{"message": ErrWrongPassword, "error": "ErrWrongPassword"})
		return
	}

}
