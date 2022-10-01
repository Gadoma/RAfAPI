package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gadoma/rafapi/internal/domain"
	"github.com/gorilla/mux"
)

func (s *Server) getAffirmationIdParameter(r *http.Request) (int, error) {
	params := mux.Vars(r)
	id, err := s.parseIdParameter(params["affirmationId"])

	if err != nil {
		return 0, fmt.Errorf("failed while parsing affirmationId parameter: %w", err)
	}

	return int(id), nil
}

func (s *Server) getAffirmationUpdate(r *http.Request) (*domain.AffirmationUpdate, error) {
	var af domain.AffirmationUpdate

	if err := json.NewDecoder(r.Body).Decode(&af); err != nil {
		return nil, fmt.Errorf("invalid JSON payload: %w", err)
	}

	return &af, nil
}