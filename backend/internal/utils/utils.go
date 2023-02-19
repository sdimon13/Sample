package utils

import (
	"encoding/json"
	"git.sample.ru/sample/internal/errors"
	"net/http"
)

type ErrorResponse struct {
	Data    []interface{} `json:"data,omitempty"`
	Message string        `json:"message"`
	Code    int           `json:"code"`
}

type ItemResponse struct {
}

type ArrayResponse struct {
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(payload)
	if err != nil {
		return
	}
}

func EmptyRespond(w http.ResponseWriter) {
	RespondWithJSON(w, http.StatusNoContent, struct{}{})
}

func NotFoundRespond(w http.ResponseWriter, msg string) {
	RespondWithJSON(w, http.StatusNotFound, errors.NotFound(msg))
}

func ErrorRespond(w http.ResponseWriter, err error) {
	r := ErrorResponse{
		Message: err.Error(),
		Code:    0,
	}

	RespondWithJSON(w, http.StatusBadRequest, r)
}
