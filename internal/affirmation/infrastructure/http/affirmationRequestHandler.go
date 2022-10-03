package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gadoma/rafapi/internal/affirmation/domain"
	"github.com/gorilla/mux"
)

type AffirmationRequestHandler struct{}

func NewAffirmationRequestHandler() *AffirmationRequestHandler {
	return &AffirmationRequestHandler{}
}

func (h *AffirmationRequestHandler) getAffirmationIdParameter(r *http.Request) (int, error) {
	params := mux.Vars(r)

	id, err := strconv.ParseUint(params["affirmationId"], 10, 64)

	if err != nil {
		return 0, fmt.Errorf("failed while parsing affirmationId parameter: %w", err)
	}

	return int(id), nil
}

func (h *AffirmationRequestHandler) getAffirmationUpdate(r *http.Request) (*domain.AffirmationUpdate, error) {
	var af domain.AffirmationUpdate

	if err := json.NewDecoder(r.Body).Decode(&af); err != nil {
		return nil, fmt.Errorf("invalid JSON payload: %w", err)
	}

	return &af, nil
}
