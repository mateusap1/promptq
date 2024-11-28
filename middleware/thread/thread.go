package thread

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mateusap1/promptq/pkg/utils"
)

var (
	ErrDoesNotExist = "thread not found"
)

func ThreadMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Requires auth middleware
		userId := c.MustGet("userId").(int64)

		tid := c.Param("tid")

		id, name, deleted, err := utils.GetThread(db, userId, tid)
		if err == sql.ErrNoRows || deleted {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": ErrDoesNotExist, "error": "ErrDoesNotExist"})
			return
		} else if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.Set("threadId", id)
		c.Set("threadName", name)

		c.Next()
	}
}
