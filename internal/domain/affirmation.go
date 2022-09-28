package domain

import (
	"context"
	"errors"
	"time"
)

type Affirmation struct {
	Id         int       `json:"id"`
	CategoryId int       `json:"categoryId"`
	Text       string    `json:"text"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type AffirmationUpdate struct {
	CategoryId int    `json:"categoryId"`
	Text       string `json:"text"`
}

func (au *AffirmationUpdate) Validate() error {
	if au.CategoryId < 1 {
		return errors.New("Affirmation CategoryId must be a positive integer")
	}

	if au.Text == "" {
		return errors.New("Affirmation Text cannot be empty")
	}

	return nil
}

type AffirmationService interface {
	GetAffirmations(ctx context.Context) ([]*Affirmation, int, error)
	GetAffirmation(ctx context.Context, id int) (*Affirmation, error)
	CreateAffirmation(ctx context.Context, au *AffirmationUpdate) (int, error)
	UpdateAffirmation(ctx context.Context, id int, au *AffirmationUpdate) error
	DeleteAffirmation(ctx context.Context, id int) error
}
