package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mateusap1/promptq/pkg/utils"
)

func CreateThread(c *gin.Context, db *sql.DB) {
	// Requires auth middleware
	userId := c.MustGet("userId").(int64)

	var form struct {
		Name string `json:"name"`
	}
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": ErrInvalidFormat, "error": "ErrInvalidFormat"})
		return
	}

	_, tid, err := utils.CreateThread(db, userId, form.Name)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{"threadId": tid})
}

func GetThreads(c *gin.Context, db *sql.DB) {
	// Requires auth middleware
	userId := c.MustGet("userId").(int64)

	threads, err := utils.GetThreads(db, userId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"threads": threads})
}

func DeleteThread(c *gin.Context, db *sql.DB) {
	// Required thread middleware
	threadId := c.MustGet("threadId").(int64)

	if err := utils.DeleteThread(db, threadId); err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{})
}

func RenameThread(c *gin.Context, db *sql.DB) {
	// Required thread middleware
	threadId := c.MustGet("threadId").(int64)

	var form struct {
		Name string `json:"name"`
	}
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": ErrInvalidFormat, "error": "ErrInvalidFormat"})
		return
	}

	if err := utils.RenameThread(db, threadId, form.Name); err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{})
}
