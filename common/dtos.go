package common

type Result struct {
	Data  interface{} `json:"data"`
	Error interface{}       `json:"error"`
}
