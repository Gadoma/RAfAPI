package mock

import (
	"context"

	"github.com/gadoma/rafapi/internal/domain"
)

var _ domain.RandomAffirmationRepository = (*RandomAffirmationRepository)(nil)

type RandomAffirmationRepository struct {
	GetRandomAffirmationsFn func(ctx context.Context, categoryIds []int) ([]*domain.RandomAffirmation, error)
}

func (r *RandomAffirmationRepository) GetRandomAffirmations(ctx context.Context, categoryIds []int) ([]*domain.RandomAffirmation, error) {
	return r.GetRandomAffirmationsFn(ctx, categoryIds)
}
