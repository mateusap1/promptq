package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mateusap1/promptq/ent"
	"github.com/mateusap1/promptq/pkg/user"
	"github.com/mateusap1/promptq/pkg/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization token required",
				"code":  "API_KEY_MISSING",
			})
			return
		}

		id, err := utils.VerifyToken(tokenString)
		if err != nil {
			fmt.Print(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid authentication token",
				"code":  "AUTH_TOKEN_INVALID",
			})
			return
		}

		c.Set("userId", id)
		c.Next()
	}
}

func GetUserFromSession(c *gin.Context, us *user.UserService) (*ent.User, error) {
	id, exists := c.Get("userId")
	if !exists {
		return nil, fmt.Errorf("user not found in context")
	}

	user, err := us.GetUser(id.(int))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token",
			"code":  "AUTH_TOKEN_CORRUPTED",
		})
		return nil, nil
	}

	return user, nil
}
