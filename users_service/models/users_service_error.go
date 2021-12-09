package models

import (
	"encoding/json"
	"errors"
	"log"
)

type UsersServiceError struct {
	Code    int    `json:"code"`
	Err     error  `json:"err"`
	Message string `json:"message"`
}

func (e *UsersServiceError) Error() string {
	json, err := json.Marshal(e)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	return string(json)
}

func Wrap(err error, info string) *UsersServiceError {
	return &UsersServiceError{
		Message: info,
		Err:     err,
	}
}

func NewUsersServiceError() error {
	return &UsersServiceError{
		Code:    2,
		Err:     errors.New("Unknown"),
		Message: "Unknown error",
	}
}
