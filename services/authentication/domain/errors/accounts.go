package errors

import "fmt"

type AuthenticationError struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

func (a AuthenticationError) Error() string {
	return fmt.Sprintf("error message : [%v], error code : [%v]", a.Message, a.Code)
}

var (
	EmailAlreadyUsed   = &AuthenticationError{Message: "email already been use", Code: "AU-001"}
	AccountNotFound    = &AuthenticationError{Message: "account not found", Code: "AU-002"}
	InvalidCredentials = &AuthenticationError{Message: "invalid credentials", Code: "AU-003"}
)
