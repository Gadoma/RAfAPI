package test

import "github.com/gadoma/rafapi/internal/domain"

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
	Data    int
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

type GetCategoriesResponse struct {
	Status  string
	Data    []domain.Category
	Count   int
	Message string
}

type GetCategoryResponse struct {
	Status  string
	Data    domain.Category
	Count   int
	Message string
}

type CreateCategoryResponse struct {
	Status  string
	Data    int
	Count   int
	Message string
}

type UpdateCategoryResponse struct {
	Status  string
	Data    []string
	Count   int
	Message string
}

type DeleteCategoryResponse struct {
	Status  string
	Data    []string
	Count   int
	Message string
}
