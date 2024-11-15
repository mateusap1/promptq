package auth

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mateusap1/promptq/pkg/utils"
)

var ErrInvalidSession = "session token does't exist"
var ErrInactiveSession = "session token is inactive"
var ErrExpiredSession = "session token has expired"

func AuthMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("session")
		if err != nil {
			if err == http.ErrNoCookie {
				c.JSON(http.StatusUnauthorized, gin.H{"message": http.ErrNoCookie.Error(), "error": "ErrNoCookie"})
				return
			} else {
				log.Fatal(err)
				return
			}
		}

		sessionId, userId, active, expiresAt, err := utils.GetSessionByToken(db, token)
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusUnauthorized, gin.H{"message": ErrInvalidSession, "error": "ErrInvalidSession"})
				return
			}
			log.Fatal(err)
			return
		}

		if !active {
			c.JSON(http.StatusUnauthorized, gin.H{"message": ErrInactiveSession, "error": "ErrInactiveSession"})
			return
		} else if time.Now().UTC().After(expiresAt) {
			c.JSON(http.StatusUnauthorized, gin.H{"message": ErrExpiredSession, "error": "ErrExpiredSession"})
			return
		}

		c.Set("sessionId", sessionId)
		c.Set("userId", userId)

		c.Next()
	}
}
