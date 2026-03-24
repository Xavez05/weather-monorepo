package apiclient

import "fmt"

type APIError struct {
	Source  string // "REST" o "SOAP"
	Message string
	Code    int
}

func (e *APIError) Error() string {
	return fmt.Sprintf("[%s] código %d: %s", e.Source, e.Code, e.Message)
}

func NewAPIError(source, message string, code int) *APIError {
	return &APIError{Source: source, Message: message, Code: code}
}
