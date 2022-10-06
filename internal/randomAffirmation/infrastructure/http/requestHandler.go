package http

import (
	"fmt"
	"net/http"

	"github.com/oklog/ulid/v2"
)

type RandomAffirmationRequestHandler struct{}

func NewRandomAffirmationRequestHandler() *RandomAffirmationRequestHandler {
	return &RandomAffirmationRequestHandler{}
}

func (h *RandomAffirmationRequestHandler) getRandomAffirmationCategoryIdsParameter(r *http.Request) ([]ulid.ULID, error) {
	q := r.URL.Query()
	queryIds := q["categoryIds"]
	var categoryIds []ulid.ULID

	for _, i := range queryIds {
		categoryId, err := ulid.Parse(i)

		if err != nil {
			return []ulid.ULID{}, fmt.Errorf("failed while parsing categoryIds parameter: %w", err)
		}

		categoryIds = append(categoryIds, categoryId)
	}

	return categoryIds, nil
}
