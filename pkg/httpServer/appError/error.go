package appError

import (
	"fmt"
	"net/http"
)

var (
	ErrOrderNotFound = NewAppError(nil, "order not found", http.StatusBadRequest)
)

type AppError struct {
	Err        error  `json:"err"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

func (e *AppError) Error() string {
	return fmt.Sprintf(e.Message+" "+e.Err.Error())
}

func NewAppError(err error, message string, statusCode int) *AppError {
	return &AppError{
		Err:              err,
		Message:          message,
		StatusCode:             statusCode,
	}
}

func systemError(err error) *AppError {
	return NewAppError(err, "internal system error", 500)
}