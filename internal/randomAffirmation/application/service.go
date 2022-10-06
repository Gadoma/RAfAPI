package application

import (
	"context"
	"fmt"
	"strings"

	"github.com/gadoma/rafapi/internal/randomAffirmation/domain"
	"github.com/oklog/ulid/v2"
)

var _ domain.RandomAffirmationService = (*RandomAffirmationService)(nil)

type RandomAffirmationService struct {
	repo domain.RandomAffirmationRepository
}

func NewRandomAffirmationService(repo domain.RandomAffirmationRepository) *RandomAffirmationService {
	return &RandomAffirmationService{repo: repo}
}

func (s *RandomAffirmationService) GetRandomAffirmation(ctx context.Context, categoryIds []ulid.ULID) (*domain.RandomAffirmation, error) {
	elements, err := s.repo.GetRandomAffirmations(ctx, categoryIds)

	if err != nil {
		return nil, fmt.Errorf("an error occurred while getting all RandomAffirmations: %w", err)
	}

	var stringElements []string
	for _, i := range elements {
		stringElements = append(stringElements, i.Text)
	}

	randomAffirmation := &domain.RandomAffirmation{
		Text: strings.Join(stringElements, " "),
	}

	return randomAffirmation, nil
}
