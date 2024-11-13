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
	ErrInvalidPasswordFormat = "invalid password format or weak password"
	ErrInvalidEmailFormat    = "invalid email format"
	ErrInvalidFormat         = "invalid request format"
)

type SignUpForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SignUp(c *gin.Context, db *sql.DB) {
	var form SignUpForm
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrInvalidFormat})
		return
	}

	email := strings.ToLower(form.Email)
	password := form.Password

	if !utils.ValidEmailFormat(email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrInvalidEmailFormat})
		return
	}

	if !utils.ValidPasswordFormat(password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrInvalidPasswordFormat})
		return
	}

	exists, err := utils.EmailAlreadyExists(db, email)
	if err != nil {
		log.Fatal(err)
		return
	}

	if exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrEmailExists})
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
