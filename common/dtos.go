package common

type Result struct {
	Data  interface{} `json:"data"`
	Error error       `json:"error"`
}
