package models

import "fmt"

type ErrorResponse struct {
	error
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("%d %s\n", e.Status, e.Message)
}
