package mock

import (
	"context"

	"github.com/gadoma/rafapi/internal/randomAffirmation/domain"
	"github.com/oklog/ulid/v2"
)

var _ domain.RandomAffirmationRepository = (*RandomAffirmationRepository)(nil)

type RandomAffirmationRepository struct {
	GetRandomAffirmationsFn func(ctx context.Context, categoryIds []ulid.ULID) ([]*domain.RandomAffirmation, error)
}

func (r *RandomAffirmationRepository) GetRandomAffirmations(ctx context.Context, categoryIds []ulid.ULID) ([]*domain.RandomAffirmation, error) {
	return r.GetRandomAffirmationsFn(ctx, categoryIds)
}
