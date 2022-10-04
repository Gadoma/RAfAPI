package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gadoma/rafapi/internal/category/domain"
	"github.com/gorilla/mux"
	"github.com/oklog/ulid/v2"
)

type CategoryRequestHandler struct{}

func NewCategoryRequestHandler() *CategoryRequestHandler {
	return &CategoryRequestHandler{}
}

func (h *CategoryRequestHandler) getCategoryIdParameter(r *http.Request) (*ulid.ULID, error) {
	params := mux.Vars(r)

	id, err := ulid.Parse(params["categoryId"])

	if err != nil {
		return nil, fmt.Errorf("failed while parsing categoryId parameter: %w", err)
	}

	return &id, nil
}

func (h *CategoryRequestHandler) getUpdateCategoryCommand(r *http.Request) (*domain.UpdateCategoryCommand, error) {
	var ccc domain.UpdateCategoryCommand

	if err := json.NewDecoder(r.Body).Decode(&ccc); err != nil {
		return nil, fmt.Errorf("invalid JSON payload: %w", err)
	}

	return &ccc, nil
}

func (h *CategoryRequestHandler) getCreateCategoryCommand(r *http.Request) (*domain.CreateCategoryCommand, error) {
	var ccc domain.CreateCategoryCommand

	if err := json.NewDecoder(r.Body).Decode(&ccc); err != nil {
		return nil, fmt.Errorf("invalid JSON payload: %w", err)
	}

	ccc.Id = ulid.Make()

	return &ccc, nil
}
