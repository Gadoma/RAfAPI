package test

import "github.com/gadoma/rafapi/internal/randomAffirmation/domain"

type GetRandomAffirmationResponse struct {
	Status  string
	Data    domain.RandomAffirmation
	Count   int
	Message string
}
