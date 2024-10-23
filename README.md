# Promptq

A REST API written in Go, designed to bridge private AI models with clients requests.

In scenarios where it is not desired to publically expose the server running the AI models, such as in a private local computer, Promptq creates a bridge allowing the AI server to remain private by querying and submitting responses to prompts through GET and POST requests.

## Running
```
$ go run main.go
```

## Testing
```
$ go test ./pkg/...
$ go test -v ./pkg/...
```
