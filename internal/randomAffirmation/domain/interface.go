package domain

import (
	"context"

	"github.com/oklog/ulid/v2"
)

type RandomAffirmationService interface {
	GetRandomAffirmation(ctx context.Context, categoryIds []ulid.ULID) (*RandomAffirmation, error)
}

type RandomAffirmationRepository interface {
	GetRandomAffirmations(ctx context.Context, categoryIds []ulid.ULID) ([]*RandomAffirmation, error)
}
