package mock

import (
	"context"

	"github.com/gadoma/rafapi/internal/affirmation/domain"
	"github.com/oklog/ulid/v2"
)

var _ domain.AffirmationRepository = (*AffirmationRepository)(nil)

type AffirmationRepository struct {
	GetAffirmationFn    func(ctx context.Context, id ulid.ULID) (*domain.Affirmation, error)
	GetAffirmationsFn   func(ctx context.Context) ([]*domain.Affirmation, int, error)
	CreateAffirmationFn func(ctx context.Context, au *domain.CreateAffirmationCommand) error
	UpdateAffirmationFn func(ctx context.Context, id ulid.ULID, au *domain.UpdateAffirmationCommand) error
	DeleteAffirmationFn func(ctx context.Context, id ulid.ULID) error
}

func (r *AffirmationRepository) GetAffirmations(ctx context.Context) ([]*domain.Affirmation, int, error) {
	return r.GetAffirmationsFn(ctx)
}

func (r *AffirmationRepository) GetAffirmation(ctx context.Context, id ulid.ULID) (*domain.Affirmation, error) {
	return r.GetAffirmationFn(ctx, id)
}

func (r *AffirmationRepository) CreateAffirmation(ctx context.Context, cac *domain.CreateAffirmationCommand) error {
	return r.CreateAffirmationFn(ctx, cac)
}

func (r *AffirmationRepository) UpdateAffirmation(ctx context.Context, id ulid.ULID, uac *domain.UpdateAffirmationCommand) error {
	return r.UpdateAffirmationFn(ctx, id, uac)
}

func (r *AffirmationRepository) DeleteAffirmation(ctx context.Context, id ulid.ULID) error {
	return r.DeleteAffirmationFn(ctx, id)
}
