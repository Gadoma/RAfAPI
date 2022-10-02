package application

import (
	"context"
	"fmt"
	"strings"

	"github.com/gadoma/rafapi/internal/domain"
)

var _ domain.RandomAffirmationService = (*RandomAffirmationService)(nil)

type RandomAffirmationService struct {
	repo domain.RandomAffirmationRepository
}

func NewRandomAffirmationService(repo domain.RandomAffirmationRepository) *RandomAffirmationService {
	return &RandomAffirmationService{repo: repo}
}

func (s *RandomAffirmationService) GetRandomAffirmation(ctx context.Context, categoryIds []int) (*domain.RandomAffirmation, error) {
	elements, err := s.repo.GetRandomAffirmations(ctx, categoryIds)

	if err != nil {
		return nil, fmt.Errorf("there was an error getting all RandomAffirmations: %w", err)
	}

	stringElements := func(items []*domain.RandomAffirmation) []string {
		var strings []string
		for _, i := range items {
			strings = append(strings, i.Text)
		}
		return strings
	}(elements)

	randomAffirmation := &domain.RandomAffirmation{
		Text: strings.Join(stringElements, " "),
	}

	return randomAffirmation, nil
}
