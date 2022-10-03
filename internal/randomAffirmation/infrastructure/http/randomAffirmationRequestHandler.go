package http

import (
	"fmt"
	"net/http"
	"strconv"
)

type RandomAffirmationRequestHandler struct{}

func NewRandomAffirmationRequestHandler() *RandomAffirmationRequestHandler {
	return &RandomAffirmationRequestHandler{}
}

func (h *RandomAffirmationRequestHandler) getRandomAffirmationCategoryIdsParameter(r *http.Request) ([]int, error) {
	q := r.URL.Query()
	queryIds := q["categoryIds"]
	categoryIds := []int{}

	for _, i := range queryIds {
		categoryId, err := strconv.ParseUint(i, 10, 64)

		if err != nil {
			return []int{}, fmt.Errorf("failed while parsing affirmationId parameter: %w", err)
		}

		categoryIds = append(categoryIds, int(categoryId))
	}

	return categoryIds, nil
}
