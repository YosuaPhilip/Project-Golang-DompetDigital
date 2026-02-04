package handler

import (
	"net/http"

	errX "wallet-service/internal/errors"
)

func MapErrorToHTTP(err error) (int, string) {
	switch err {
	case nil:
		return http.StatusOK, ""
	case errX.ErrUserNotFound:
		return http.StatusNotFound, err.Error()
	case errX.ErrInsufficientBalance:
		return http.StatusConflict, err.Error()
	default:
		return http.StatusInternalServerError, "internal server error"
	}
}
