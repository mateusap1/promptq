package api

type HealthResponse struct {
	Message string
}

type RequestPromptResponse struct {
	QueueId string
	Prompt  string
}

type PromptResponse struct {
	Prompt   string
	State    string
	Response string
}

type PeekPromptResponse = PromptResponse
type DeQueuePromptResponse = PromptResponse
type RespondPromptResponse = PromptResponse

type CreatePromptRequest struct {
	Prompt string
}

type QueuePromptRequest struct {
	Auth string
}
