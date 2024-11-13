package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/argon2"
)

func HashPassword(rawPassword string, salt []byte) []byte {
	return argon2.IDKey([]byte(rawPassword), salt, 1, 64*1024, 4, 32)
}

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
