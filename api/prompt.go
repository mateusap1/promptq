package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/mateusap1/promptq/ent"
	"github.com/mateusap1/promptq/middleware/auth"
	"github.com/mateusap1/promptq/pkg/prompt"
	"github.com/mateusap1/promptq/pkg/user"
)

func raiseInvalidPrompt(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"error": "Invalid prompt",
		"code":  "WRONG_FORMAT",
	})
}

func raiseServerError(c *gin.Context, err string) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
		"error": "internal server error",
		"code":  "INTERNAL_SERVER_ERROR",
	})
	return
}

func CreatePrompt(c *gin.Context, us *user.UserService, ps *prompt.PromptService) {
	var form struct {
		Prompt string
	}

	if err := c.BindJSON(&form); err != nil {
		raiseInvalidPrompt(c)
		return
	}

	user, err := auth.GetUser(c)
	if err != nil {
		raiseServerError(c, "failed getting user from session")
		return
	}

	pr, err := ps.MakePromptRequest(form.Prompt, user)
	if err != nil {
		raiseServerError(c, "failed making prompt request")
		return
	}

	c.IndentedJSON(http.StatusOK, RequestPromptResponse{QueueId: pr.Identifier.String(), Prompt: pr.Prompt})
}

func QueuePrompt(c *gin.Context, us *user.UserService, ps *prompt.PromptService) {
	apiKey, present := os.LookupEnv("API_KEY")
	if !present {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"Message": "Undefined env variable.",
		})
		return
	}

	hasQueue, err := ps.HasQueuePromptRequest()
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
		pr, err := ps.QueuePromptRequest()
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

func AnswerPrompt(c *gin.Context, ps *prompt.PromptService) {
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

	pr, err := ps.GetPromptRequest(c.Param("id"))
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

	_, err = ps.MakePromptResponse(pr, responseForm.Response)
	if err != nil {
		fmt.Print(err)

		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"Message": "Error while to make prompt response.",
		})
		return
	}

	_, err = ps.AnswerPromptRequest(pr)
	if err != nil {
		fmt.Print(err)

		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"Message": "Error while to answer prompt request.",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, PromptResponse{Prompt: pr.Prompt, State: "answered", Response: responseForm.Response})
}

func GetPrompt(c *gin.Context, ps *prompt.PromptService) {
	pr, err := ps.GetPromptRequest(c.Param("id"))
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
		prp, err := ps.GetPromptResponse(pr)
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

func GetPrompts(c *gin.Context, us *user.UserService, ps *prompt.PromptService) {
	var promptsForm GetPromptsRequest

	if err := c.BindJSON(&promptsForm); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"Message": "Authentication key not provided.",
		})
		return
	}

	user, err := us.GetUser(promptsForm.ApiKey)
	if err != nil {
		fmt.Print(err)

		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"Message": "User authentication error.",
		})
		return
	}

	requests, err := ps.GetPromptRequests(user)
	if err != nil {
		fmt.Print(err)

		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"Message": "Error while trying to get user requests.",
		})
		return
	}

	responses := make([]PromptRequestResponse, len(requests))
	for i := range requests {
		pr := requests[i]

		var state string = "awaiting"
		if pr.IsAnswered {
			state = "answered"
		} else if pr.IsQueued {
			state = "queued"
		}

		responses[i] = PromptRequestResponse{Identifier: pr.Identifier.String(), Prompt: pr.Prompt, State: state}
	}

	c.IndentedJSON(http.StatusOK, GetPromptsRespose{responses: responses})
}
