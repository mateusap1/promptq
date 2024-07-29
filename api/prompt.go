package api

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/joho/godotenv"

	"github.com/mateusap1/promptq/ent"
	"github.com/mateusap1/promptq/ent/promptrequest"
	"github.com/mateusap1/promptq/pkg/prompt"
)

func loadDotEnv() error {
	err := godotenv.Load(".env")

	return err
}

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
	err := loadDotEnv()
	if err != nil {
		fmt.Print(err)

		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"Message": "Internal Server Error",
		})
		return
	}

	API_KEY, present := os.LookupEnv("API_KEY")
	if !present {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"Message": "Undefined env variable.",
		})
		return
	}

	var queueForm QueuePromptRequest

	if err := c.BindJSON(&queueForm); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"Message": "Authentication key not provided.",
		})
		return
	}

	if queueForm.Auth != API_KEY {
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

func AnswerPrompt(c *gin.Context, ctx context.Context, client *ent.Client) {
	err := loadDotEnv()
	if err != nil {
		fmt.Print(err)

		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"Message": "Internal Server Error",
		})
		return
	}

	API_KEY, present := os.LookupEnv("API_KEY")
	if !present {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"Message": "Undefined env variable.",
		})
		return
	}

	var responseForm RespondPromptRequest

	if err := c.BindJSON(&responseForm); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"Message": "Prompt response form not provided.",
		})
		return
	}

	if responseForm.Auth != API_KEY {
		c.IndentedJSON(http.StatusForbidden, gin.H{
			"Message": "Incorrect authentication key.",
		})
		return
	}

	promptId := c.Param("id")
	promptIdentifier, err := uuid.Parse(promptId)
	if err != nil {
		c.IndentedJSON(http.StatusForbidden, gin.H{
			"Message": "Prompt ID with wrong format.",
		})
		return
	}

	pr, err := client.PromptRequest.Query().Where(promptrequest.Identifier(promptIdentifier)).Only(ctx)
	if err != nil {
		fmt.Print(err)

		if ent.IsNotFound(err) {
			c.IndentedJSON(http.StatusNotFound, gin.H{
				"Message": "Queue ID not found.",
			})
			return
		} else {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{
				"Message": "Error while trying to query prompt request.",
			})
			return
		}
	}

	_, err = prompt.MakePromptResponse(ctx, client, pr, responseForm.Response)
	if err != nil {
		fmt.Print(err)

		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"Message": "Error while to make prompt response.",
		})
		return
	}

	_, err = prompt.AnswerPromptRequest(ctx, client, pr)
	if err != nil {
		fmt.Print(err)

		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"Message": "Error while to answer prompt request.",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, PromptResponse{Prompt: pr.Prompt, State: "answered", Response: responseForm.Response})
}

func GetPrompt(c *gin.Context, ctx context.Context, client *ent.Client) {
	promptId := c.Param("id")
	promptIdentifier, err := uuid.Parse(promptId)
	if err != nil {
		c.IndentedJSON(http.StatusForbidden, gin.H{
			"Message": "Prompt ID with wrong format.",
		})
		return
	}

	pr, err := client.PromptRequest.Query().Where(promptrequest.Identifier(promptIdentifier)).Only(ctx)
	if err != nil {
		fmt.Print(err)

		if ent.IsNotFound(err) {
			c.IndentedJSON(http.StatusNotFound, gin.H{
				"Message": "Queue ID not found.",
			})
			return
		} else {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{
				"Message": "Error while trying to query prompt request.",
			})
			return
		}
	}

	if pr.IsAnswered {
		prp, err := pr.QueryPromptResponse().Only(ctx)
		if err != nil {
			fmt.Print(err)

			c.IndentedJSON(http.StatusInternalServerError, gin.H{
				"Message": "Error while to query prompt response.",
			})
			return
		}

		c.IndentedJSON(http.StatusOK, PromptResponse{Prompt: pr.Prompt, State: "answered", Response: prp.Response})
	} else if pr.IsQueued {
		c.IndentedJSON(http.StatusOK, PromptResponse{Prompt: pr.Prompt, State: "queued"})
	} else {
		c.IndentedJSON(http.StatusOK, PromptResponse{Prompt: pr.Prompt, State: "awaiting"})
	}
}
