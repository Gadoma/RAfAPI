package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gadoma/rafapi/internal/category/domain"
	"github.com/gorilla/mux"
)

type CategoryRequestHandler struct{}

func NewCategoryRequestHandler() *CategoryRequestHandler {
	return &CategoryRequestHandler{}
}

func (h *CategoryRequestHandler) getCategoryIdParameter(r *http.Request) (int, error) {
	params := mux.Vars(r)

	id, err := strconv.ParseUint(params["categoryId"], 10, 64)

	if err != nil {
		return 0, fmt.Errorf("failed while parsing categoryId parameter: %w", err)
	}

	return int(id), nil
}

func (h *CategoryRequestHandler) getCategoryUpdate(r *http.Request) (*domain.CategoryUpdate, error) {
	var af domain.CategoryUpdate

	if err := json.NewDecoder(r.Body).Decode(&af); err != nil {
		return nil, fmt.Errorf("invalid JSON payload: %w", err)
	}

	return &af, nil
}
