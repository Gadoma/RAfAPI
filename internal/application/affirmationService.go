package application

import (
	"context"
	"fmt"

	"github.com/gadoma/rafapi/internal/domain"
)

var _ domain.AffirmationService = (*AffirmationService)(nil)

type AffirmationService struct {
	repo domain.AffirmationRepository
}

func NewAffirmationService(repo domain.AffirmationRepository) *AffirmationService {
	return &AffirmationService{repo: repo}
}

func (s *AffirmationService) GetAffirmations(ctx context.Context) ([]*domain.Affirmation, int, error) {
	a, n, err := s.repo.GetAffirmations(ctx)

	if err != nil {
		return nil, 0, fmt.Errorf("there was an error getting all Affirmations: %w", err)
	}

	return a, n, nil
}

func (s *AffirmationService) GetAffirmation(ctx context.Context, id int) (*domain.Affirmation, error) {
	a, err := s.repo.GetAffirmation(ctx, id)

	if err != nil {
		return nil, fmt.Errorf("there was an error getting the Affirmation by id: %w", err)
	}

	if a == nil {
		return nil, domain.ErrorResourceNotFound
	}

	return a, nil
}

func (s *AffirmationService) CreateAffirmation(ctx context.Context, au *domain.AffirmationUpdate) (int, error) {

	if err := au.Validate(); err != nil {
		return 0, err
	}

	id, err := s.repo.CreateAffirmation(ctx, au)

	if err != nil {
		return 0, fmt.Errorf("there was an error creating the Affirmation: %w", err)
	}

	return id, nil
}

func (s *AffirmationService) UpdateAffirmation(ctx context.Context, id int, au *domain.AffirmationUpdate) error {

	if err := au.Validate(); err != nil {
		return err
	}

	if err := s.repo.UpdateAffirmation(ctx, id, au); err != nil {
		return fmt.Errorf("there was an error updating the Affirmation: %w", err)
	}

	return nil
}

func (s *AffirmationService) DeleteAffirmation(ctx context.Context, id int) error {
	if err := s.repo.DeleteAffirmation(ctx, id); err != nil {
		return fmt.Errorf("there was an error deleting the Affirmation: %w", err)
	}

	return nil
}
