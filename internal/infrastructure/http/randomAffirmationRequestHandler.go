package http

import (
	"fmt"
	"net/http"
)

func (s *Server) getRandomAffirmationCategoryIdsParameter(r *http.Request) ([]int, error) {
	q := r.URL.Query()
	queryIds := q["categoryIds"]
	categoryIds := []int{}

	for _, i := range queryIds {
		categoryId, err := s.parseIdParameter(i)

		if err != nil {
			return []int{}, fmt.Errorf("failed while parsing affirmationId parameter: %w", err)
		}

		categoryIds = append(categoryIds, int(categoryId))

	}

	return categoryIds, nil
}
