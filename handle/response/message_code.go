package response

import (
	"net/http"
)

var (
	ErrIDRequired     = &APIError{Message: "ID must not be null or blank", HTTPStatus: http.StatusBadRequest}
	ErrInvalidReq     = &APIError{Message: "Invalid request", HTTPStatus: http.StatusBadRequest}
	ErrEmailDuplicate = &APIError{Message: "Email already exist", HTTPStatus: http.StatusConflict}
	ErrUserNotFound   = &APIError{Message: "User not found", HTTPStatus: http.StatusNotFound}
	ErrUnauthorized   = &APIError{Message: "Unauthorized", HTTPStatus: http.StatusUnauthorized}
	ErrForbidden      = &APIError{Message: "Forbidden", HTTPStatus: http.StatusForbidden}
	ErrInternalServer = &APIError{
		Message: "Internal server error", HTTPStatus: http.StatusInternalServerError,
	}
)
