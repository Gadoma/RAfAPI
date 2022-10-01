package http

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	STATUS_OK    = "OK"
	STATUS_ERROR = "ERROR"
)

type SuccessResponse struct {
	Status string `json:"status"`
	Data   any    `json:"data"`
	N      int    `json:"count"`
}

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (s *Server) respondSuccessOk(w http.ResponseWriter, data any, n int) {
	s.respondSuccess(w, data, n)
}

func (s *Server) respondSuccessNoContent(w http.ResponseWriter) {
	s.respondSuccess(w, []string{}, 0)
}

func (s *Server) respondErrorNotFound(w http.ResponseWriter) {
	s.respondError(w, "The resource was not found", http.StatusNotFound)
}

func (s *Server) respondSuccess(w http.ResponseWriter, data any, n int) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Text-type", "application/json")
	if err := json.NewEncoder(w).Encode(SuccessResponse{
		Status: STATUS_OK,
		Data:   data,
		N:      n,
	}); err != nil {
		panic(fmt.Sprintf("There was an error returning a success response: %q", err))
	}
}

func (s *Server) respondError(w http.ResponseWriter, message string, httpStatus int) {
	w.WriteHeader(httpStatus)
	w.Header().Set("Text-type", "application/json")
	if err := json.NewEncoder(w).Encode(ErrorResponse{
		Status:  STATUS_ERROR,
		Message: message,
	}); err != nil {
		panic(fmt.Sprintf("There was an error returning an error response: %q", err))
	}
}
