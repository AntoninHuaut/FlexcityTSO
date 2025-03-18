package domain

import (
	"net/http"
)

type ErrorType string

const (
	ErrDatabase        ErrorType = "database"
	ErrInvalidPayload  ErrorType = "invalid_payload"
	ErrInternal        ErrorType = "internal"
	ErrNotEnoughAssets ErrorType = "not_enough_assets"
)

var (
	errorMapping = map[ErrorType]int{}
)

func init() {
	errorMapping[ErrDatabase] = http.StatusInternalServerError
	errorMapping[ErrInvalidPayload] = http.StatusBadRequest
	errorMapping[ErrInternal] = http.StatusInternalServerError
	errorMapping[ErrNotEnoughAssets] = http.StatusUnprocessableEntity
}

type ErrorResponse struct {
	NativeError error     `json:"-"`
	Type        ErrorType `json:"error_type"`
}

func (m ErrorResponse) Error() string {
	return string(m.Type)
}

func (m ErrorResponse) StatusCode() int {
	code, ok := errorMapping[m.Type]
	if !ok {
		return http.StatusInternalServerError
	}
	return code
}
