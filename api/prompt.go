package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mateusap1/promptq/ent"
	"github.com/mateusap1/promptq/pkg/prompt"
)

const ApiKey = "secret"

func CreatePrompt(c *gin.Context, ctx context.Context, client *ent.Client) {
	var promptForm CreatePromptRequest

	if err := c.BindJSON(&promptForm); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"Message": "Prompt not provided or formatted incorrectly.",
		})
		return
	}

	pr, err := prompt.MakePromptRequest(ctx, client, promptForm.Prompt)
	if err != nil {
		fmt.Print(err)

		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"Message": "Error while trying to make prompt request",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, RequestPromptResponse{QueueId: pr.Identifier.String(), Prompt: pr.Prompt})
}

func QueuePrompt(c *gin.Context, ctx context.Context, client *ent.Client) {
	var queueForm QueuePromptRequest

	if err := c.BindJSON(&queueForm); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"Message": "Authentication key not provided.",
		})
		return
	}

	if queueForm.Auth != ApiKey {
		c.IndentedJSON(http.StatusForbidden, gin.H{
			"Message": "Incorrect authentication key.",
		})
		return
	}

	hasQueue, err := prompt.HasQueuePromptRequest(ctx, client)
	if err != nil {
		fmt.Print(err)

		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"Message": "Error while trying to query size of prompt request queue",
		})
		return
	}

	if !hasQueue {
		c.IndentedJSON(http.StatusOK, gin.H{
			"Message": "No prompts left.",
		})
	} else {
		pr, err := prompt.QueuePromptRequest(ctx, client)
		if err != nil {
			fmt.Print(err)

			c.IndentedJSON(http.StatusInternalServerError, gin.H{
				"Message": "Error while trying to query size of prompt request queue",
			})
			return
		}

		c.IndentedJSON(http.StatusOK, RequestPromptResponse{QueueId: pr.Identifier.String(), Prompt: pr.Prompt})
	}
}
