package model

type Request struct {
	Message map[string]interface{} `json:"msg,omitempty"`
}
