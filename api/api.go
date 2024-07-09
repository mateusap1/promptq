package api

type HealthResponse struct {
	Message string
}

type RequestPromptResponse struct {
	QueueId string
}

type PromptResponse struct {
	Prompt   string
	State    string
	Response string
}

type PeekPromptResponse = PromptResponse
type DeQueuePromptResponse = PromptResponse
type RespondPromptResponse = PromptResponse
