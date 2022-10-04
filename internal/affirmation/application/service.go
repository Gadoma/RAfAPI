package application

import (
	"context"
	"fmt"

	"github.com/gadoma/rafapi/internal/affirmation/domain"
	common "github.com/gadoma/rafapi/internal/common/domain"
	"github.com/oklog/ulid/v2"
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
		return nil, 0, fmt.Errorf("an error occurred while getting all Affirmations: %w", err)
	}

	return a, n, nil
}

func (s *AffirmationService) GetAffirmation(ctx context.Context, id ulid.ULID) (*domain.Affirmation, error) {
	a, err := s.repo.GetAffirmation(ctx, id)

	if err != nil {
		return nil, fmt.Errorf("an error occurred while getting the Affirmation by id: %w", err)
	}

	if a == nil {
		return nil, common.ErrorResourceNotFound
	}

	return a, nil
}

func (s *AffirmationService) CreateAffirmation(ctx context.Context, cac *domain.CreateAffirmationCommand) (*ulid.ULID, error) {
	if err := cac.Validate(); err != nil {
		return nil, err
	}

	err := s.repo.CreateAffirmation(ctx, cac)

	if err != nil {
		return nil, fmt.Errorf("an error occurred while creating the Affirmation: %w", err)
	}

	return &cac.Id, nil
}

func (s *AffirmationService) UpdateAffirmation(ctx context.Context, id ulid.ULID, uac *domain.UpdateAffirmationCommand) error {
	if err := uac.Validate(); err != nil {
		return err
	}

	if err := s.repo.UpdateAffirmation(ctx, id, uac); err != nil {
		return fmt.Errorf("an error occurred while updating the Affirmation: %w", err)
	}

	return nil
}

func (s *AffirmationService) DeleteAffirmation(ctx context.Context, id ulid.ULID) error {
	if err := s.repo.DeleteAffirmation(ctx, id); err != nil {
		return fmt.Errorf("an error occurred while deleting the Affirmation: %w", err)
	}

	return nil
}
