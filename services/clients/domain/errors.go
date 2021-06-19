package domain

import "fmt"

type ErrorCode string

var (
	InvalidPhone     ErrorCode = "C-001"
	DuplicatedPhone  ErrorCode = "C-002"
	PersistencyLayer ErrorCode = "C-003"
)

type ClientError struct {
	Msg  string      `json:"msg"`
	Code ErrorCode   `json:"code"`
	Data interface{} `json:"data"`
}

func (c *ClientError) Error() string {
	return fmt.Sprintf("[%v] %v ", c.Code, c.Msg)
}

var (
	InvalidPhoneErr      = &ClientError{Code: InvalidPhone, Msg: "phone number is not valid."}
	DuplicatedPhoneError = &ClientError{Code: DuplicatedPhone, Msg: "phone number is already been used."}
)

func NewPersistenceLayerError(data interface{}) *ClientError {
	return &ClientError{
		Code: PersistencyLayer,
		Msg:  "a unexpected error occurs on the persistency layer",
		Data: data}
}
