package domain

import (
	"context"

	"github.com/oklog/ulid/v2"
)

type AffirmationService interface {
	GetAffirmations(ctx context.Context) ([]*Affirmation, int, error)
	GetAffirmation(ctx context.Context, id ulid.ULID) (*Affirmation, error)
	CreateAffirmation(ctx context.Context, cac *CreateAffirmationCommand) (*ulid.ULID, error)
	UpdateAffirmation(ctx context.Context, id ulid.ULID, uac *UpdateAffirmationCommand) error
	DeleteAffirmation(ctx context.Context, id ulid.ULID) error
}

type AffirmationRepository interface {
	GetAffirmations(ctx context.Context) ([]*Affirmation, int, error)
	GetAffirmation(ctx context.Context, id ulid.ULID) (*Affirmation, error)
	CreateAffirmation(ctx context.Context, cac *CreateAffirmationCommand) error
	UpdateAffirmation(ctx context.Context, id ulid.ULID, uac *UpdateAffirmationCommand) error
	DeleteAffirmation(ctx context.Context, id ulid.ULID) error
}
