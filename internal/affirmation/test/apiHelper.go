package test

import "github.com/gadoma/rafapi/internal/affirmation/domain"

type GetAffirmationsResponse struct {
	Status  string
	Data    []domain.Affirmation
	Count   int
	Message string
}

type GetAffirmationResponse struct {
	Status  string
	Data    domain.Affirmation
	Count   int
	Message string
}

type CreateAffirmationResponse struct {
	Status  string
	Data    string
	Count   int
	Message string
}

type UpdateAffirmationResponse struct {
	Status  string
	Data    []string
	Count   int
	Message string
}

type DeleteAffirmationResponse struct {
	Status  string
	Data    []string
	Count   int
	Message string
}
