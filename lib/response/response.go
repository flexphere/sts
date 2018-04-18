package response

import (
	"fmt"
	"net/http"
)

type Response struct {
	Status bool        `json:"status"`
	Error  string      `json:"error,omitempty"`
	Body   interface{} `json:"body,omitempty"`
}

func New(b map[string]interface{}) *Response {
	return &Response{
		Status: true,
		Body:   b,
	}
}

func NewNotFound() (int, interface{}) {
	return http.StatusNotFound, &Response{
		Status: false,
		Error:  fmt.Sprintf("The requested id was not found."),
	}
}

func NewInternalServerError() (int, interface{}) {
	return http.StatusInternalServerError, &Response{
		Status: false,
		Error:  fmt.Sprintf("500 Internal Server Error. Contact the server administrator."),
	}
}
