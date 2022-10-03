package http

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	statusOk    = "OK"
	statusError = "ERROR"
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

type Responder struct{}

func NewResponder() *Responder {
	return &Responder{}
}

func (r *Responder) RespondSuccessOk(w http.ResponseWriter, data any, n int) {
	r.respondSuccess(w, data, n)
}

func (r *Responder) RespondSuccessNoContent(w http.ResponseWriter) {
	r.respondSuccess(w, []string{}, 0)
}

func (r *Responder) RespondErrorNotFound(w http.ResponseWriter) {
	r.RespondError(w, "The resource was not found", http.StatusNotFound)
}

func (r *Responder) RespondError(w http.ResponseWriter, message string, httpStatus int) {
	w.WriteHeader(httpStatus)
	w.Header().Set("Text-type", "application/json")
	if err := json.NewEncoder(w).Encode(ErrorResponse{
		Status:  statusError,
		Message: message,
	}); err != nil {
		panic(fmt.Sprintf("There was an error returning an error response: %q", err))
	}
}

func (r *Responder) respondSuccess(w http.ResponseWriter, data any, n int) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Text-type", "application/json")
	if err := json.NewEncoder(w).Encode(SuccessResponse{
		Status: statusOk,
		Data:   data,
		N:      n,
	}); err != nil {
		panic(fmt.Sprintf("There was an error returning a success response: %q", err))
	}
}
