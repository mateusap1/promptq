package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mateusap1/promptq/ent"
	"github.com/mateusap1/promptq/pkg/user"
)

func CreateUser(c *gin.Context, ctx context.Context, client *ent.Client) {
	var userForm CreateUserRequest

	if err := c.BindJSON(&userForm); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"Message": "User not provided or formatted incorrectly.",
		})
		return
	}

	us, err := user.MakeUser(ctx, client, userForm.UserName)
	if err != nil {
		fmt.Print(err)

		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"Message": "Error while trying to register user",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, CreateUserResponse{UserName: us.Username, ApiKey: us.APIKey})
}
