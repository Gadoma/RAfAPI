package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gadoma/rafapi/internal/affirmation/domain"
	"github.com/gorilla/mux"
	"github.com/oklog/ulid/v2"
)

type AffirmationRequestHandler struct{}

func NewAffirmationRequestHandler() *AffirmationRequestHandler {
	return &AffirmationRequestHandler{}
}

func (h *AffirmationRequestHandler) getAffirmationIdParameter(r *http.Request) (*ulid.ULID, error) {
	params := mux.Vars(r)

	id, err := ulid.Parse(params["affirmationId"])

	if err != nil {
		return nil, fmt.Errorf("failed while parsing affirmationId parameter: %w", err)
	}

	return &id, nil
}

func (h *AffirmationRequestHandler) getUpdateAffirmationCommand(r *http.Request) (*domain.UpdateAffirmationCommand, error) {
	var uac domain.UpdateAffirmationCommand

	if err := json.NewDecoder(r.Body).Decode(&uac); err != nil {
		return nil, fmt.Errorf("invalid JSON payload: %w", err)
	}

	return &uac, nil
}

func (h *AffirmationRequestHandler) getCreateAffirmationCommand(r *http.Request) (*domain.CreateAffirmationCommand, error) {
	var cac domain.CreateAffirmationCommand

	if err := json.NewDecoder(r.Body).Decode(&cac); err != nil {
		return nil, fmt.Errorf("invalid JSON payload: %w", err)
	}

	cac.Id = ulid.Make()

	return &cac, nil
}
