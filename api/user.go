package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SignUpForm struct {
	Email    string
	Password string
}

func SignUp(c *gin.Context) {
	var user SignUpForm
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

}
