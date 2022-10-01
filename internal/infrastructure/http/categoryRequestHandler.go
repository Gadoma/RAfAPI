package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gadoma/rafapi/internal/domain"
	"github.com/gorilla/mux"
)

func (s *Server) getCategoryIdParameter(r *http.Request) (int, error) {
	params := mux.Vars(r)
	id, err := s.parseIdParameter(params["categoryId"])

	if err != nil {
		return 0, fmt.Errorf("failed while parsing categoryId parameter: %w", err)
	}

	return int(id), nil
}

func (s *Server) getCategoryUpdate(r *http.Request) (*domain.CategoryUpdate, error) {
	var af domain.CategoryUpdate

	if err := json.NewDecoder(r.Body).Decode(&af); err != nil {
		return nil, fmt.Errorf("invalid JSON payload: %w", err)
	}

	return &af, nil
}
