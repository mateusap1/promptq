package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getHealth(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "healthy",
	})
}

func main() {
	router := gin.Default()
	router.GET("/health", getHealth)

	router.Run("localhost:8080")
}
