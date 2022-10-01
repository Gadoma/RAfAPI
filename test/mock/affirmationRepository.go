package mock

import (
	"context"

	"github.com/gadoma/rafapi/internal/domain"
)

var _ domain.AffirmationRepository = (*AffirmationRepository)(nil)

type AffirmationRepository struct {
	GetAffirmationFn    func(ctx context.Context, id int) (*domain.Affirmation, error)
	GetAffirmationsFn   func(ctx context.Context) ([]*domain.Affirmation, int, error)
	CreateAffirmationFn func(ctx context.Context, au *domain.AffirmationUpdate) (int, error)
	UpdateAffirmationFn func(ctx context.Context, id int, au *domain.AffirmationUpdate) error
	DeleteAffirmationFn func(ctx context.Context, id int) error
}

func (r *AffirmationRepository) GetAffirmations(ctx context.Context) ([]*domain.Affirmation, int, error) {
	return r.GetAffirmationsFn(ctx)
}

func (r *AffirmationRepository) GetAffirmation(ctx context.Context, id int) (*domain.Affirmation, error) {
	return r.GetAffirmationFn(ctx, id)
}

func (r *AffirmationRepository) CreateAffirmation(ctx context.Context, au *domain.AffirmationUpdate) (int, error) {
	return r.CreateAffirmationFn(ctx, au)
}

func (r *AffirmationRepository) UpdateAffirmation(ctx context.Context, id int, au *domain.AffirmationUpdate) error {
	return r.UpdateAffirmationFn(ctx, id, au)
}

func (r *AffirmationRepository) DeleteAffirmation(ctx context.Context, id int) error {
	return r.DeleteAffirmationFn(ctx, id)
}
