package response

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func respondJSON(w http.ResponseWriter, HTTPStatus int, resource *APIResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(HTTPStatus)
	if err := json.NewEncoder(w).Encode(resource); err != nil {
		zap.L().Error("respondJSON", zap.Error(err))
		return
	}
}

func OK(w http.ResponseWriter, resource any) {
	respondJSON(w, http.StatusOK, &APIResponse{resource})
}

func NoContent(w http.ResponseWriter) {
	respondJSON(w, http.StatusNoContent, nil)
}

func Created(w http.ResponseWriter, resource any) {
	respondJSON(w, http.StatusCreated, &APIResponse{resource})
}

func ErrorJSON(w http.ResponseWriter, err error) {
	var apiErr *APIError
	ok := errors.As(err, &apiErr)
	if !ok {
		zap.L().Error("ErrorJSON", zap.Error(err))
		apiErr = ErrInternalServer
	}
	respondJSON(w, apiErr.HTTPStatus, &APIResponse{apiErr.Message})
}

func InvalidReq(w http.ResponseWriter) {
	ErrorJSON(w, ErrInvalidReq)
}

func ErrValidation(w http.ResponseWriter, err error) {
	var validationErrors validator.ValidationErrors
	ok := errors.As(err, &validationErrors)
	if ok {
		ErrorJSON(w, &APIError{validationErrors.Error(), ErrInvalidReq.HTTPStatus})
		return
	}
	ErrorJSON(w, ErrInvalidReq)
}
