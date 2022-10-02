package domain

import (
	"context"
)

type RandomAffirmation struct {
	Text string `json:"text"`
}

type RandomAffirmationService interface {
	GetRandomAffirmation(ctx context.Context, categortIds []int) (*RandomAffirmation, error)
}

type RandomAffirmationRepository interface {
	GetRandomAffirmations(ctx context.Context, categortIds []int) ([]*RandomAffirmation, error)
}
