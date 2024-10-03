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

type CreatePromptRequest struct {
	Prompt string
	Auth   string
}

type QueuePromptRequest struct {
	Auth string
}

type RespondPromptRequest struct {
	Auth     string
	Response string
}

type CreateUserRequest struct {
	UserName string
}

type CreateUserResponse struct {
	UserName string
	ApiKey   string
}
