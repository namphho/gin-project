package common

import (
	"errors"
	"net/http"
)

type successRes struct {
	Data   interface{} `json:"data"`
	Paging interface{} `json:"paging,omitempty"`
	Filter interface{} `json:"filter,omitempty"`
}

func NewSuccessResponse(data, paging, filter interface{}) *successRes {
	return &successRes{data, paging, filter}
}

func SimpleSuccessResponse(data interface{}) *successRes {
	return &successRes{data, nil, nil}
}

type AppError struct {
	StatusCode int    `json:"status_code"`
	RootErr    error  `json:"-"`
	Message    string `json:"message"`
	Log        string `json:"log"`
	Key        string `json:"key"`
}

func NewErrorResponse(root error, message, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr:    root,
		Message:    message,
		Log:        log,
		Key:        key,
	}
}

func NewUnauthorized(root error, message, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusUnauthorized,
		RootErr:    root,
		Message:    message,
		Key:        key,
	}
}

func NewCustomError(root error, message, key string) *AppError {
	if root != nil {
		return NewErrorResponse(root, message, root.Error(), key)
	}
	return NewErrorResponse(errors.New(message), message, message, key)
}

func (e *AppError) RootError() error {
	if err, ok := e.RootErr.(*AppError); ok {
		return err.RootErr
	}
	return e.RootErr
}

func (e *AppError) Error() string {
	return e.RootErr.Error()
}
