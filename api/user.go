package api

import (
	"database/sql"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mateusap1/promptq/pkg/utils"
)

var (
	ErrEmailExists           = errors.New("email taken")
	ErrInvalidPasswordFormat = errors.New("invalid password format or weak password")
	ErrInvalidEmailFormat    = errors.New("invalid email format")
	ErrInvalidFormat         = errors.New("invalid request format")
)

type SignUpForm struct {
	Email    string
	Password string
}

func SignUp(c *gin.Context, db *sql.DB) {
	var form SignUpForm
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrInvalidFormat})
		return
	}

	if !utils.ValidEmailFormat(form.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrInvalidEmailFormat})
		return
	}

	if !utils.ValidPasswordFormat(form.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrInvalidPasswordFormat})
		return
	}

	exists, err := utils.EmailAlreadyExists(db, form.Email)
	if err != nil {
		log.Fatal(err)
		return
	}

	if exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrEmailExists})
		return
	}

	passwordHash := utils.EncodePassword(form.Password)
	confirmToken, err := utils.CreateUser(db, form.Email, passwordHash)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Send confirmation e-mail
	if err := utils.SendConfirmationEmail(confirmToken); err != nil {
		log.Fatal(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
