package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mateusap1/promptq/pkg/user"
)

func CreateUser(c *gin.Context, us *user.UserService) {
	var form struct {
		Username string
		Password string
	}

	if err := c.BindJSON(&form); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"Message": "User not provided or formatted incorrectly.",
		})
		return
	}

	_, err := us.MakeUser(form.Username, form.Password)
	if err != nil {
		fmt.Print(err)

		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"Message": "Error while trying to register user",
		})
		return
	}

	c.Request.Header.Add("")
	c.IndentedJSON(http.StatusOK, Response{message: "created user successfully"})
}
