package api

type Response struct {
	message string
}

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

type RespondPromptRequest struct {
	Response string
}

type GetPromptsRequest struct {
	ApiKey string
}

type PromptRequestResponse struct {
	Identifier string
	Prompt     string
	State      string
}

type GetPromptsRespose struct {
	responses []PromptRequestResponse
}
