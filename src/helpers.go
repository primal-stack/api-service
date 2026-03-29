package helpers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ValidateError struct {
	Errors map[string]string `json:"errors"`
}

func ErrorResponseWithCode(code int, message string) *ErrorResponse {
	return &ErrorResponse{
		Code:    code,
		Message: message,
	}
}

func ValidateErrorWithErrors(errors map[string]string) *ValidateError {
	return &ValidateError{
		Errors: errors,
	}
}

func IsError(err error) bool {
	return err!= nil
}

func GetLogger() *logrus.Logger {
	return logrus.New()
}

func GetHTTPStatusCode(status int) int {
	switch status {
	case http.StatusOK:
		return http.StatusOK
	case http.StatusCreated:
		return http.StatusCreated
	case http.StatusAccepted:
		return http.StatusAccepted
	case http.StatusNoContent:
		return http.StatusNoContent
	case http.StatusBadRequest:
		return http.StatusBadRequest
	case http.StatusUnauthorized:
		return http.StatusUnauthorized
	case http.StatusForbidden:
		return http.StatusForbidden
	case http.StatusInternalServerError:
		return http.StatusInternalServerError
	case http.StatusNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}

func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func GetUTCNow() time.Time {
	return time.Now().UTC()
}