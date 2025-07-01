package models

import "fmt"

type CustomError struct {
	Code    int
	Message string
}

// This makes CustomError implement the error interface
func (e CustomError) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}
