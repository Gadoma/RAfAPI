package test

import (
	aff "github.com/gadoma/rafapi/internal/affirmation/domain"
	cat "github.com/gadoma/rafapi/internal/category/domain"
	raf "github.com/gadoma/rafapi/internal/randomAffirmation/domain"
)

type GetRandomAffirmationResponse struct {
	Status  string
	Data    raf.RandomAffirmation
	Count   int
	Message string
}

type GetAffirmationsResponse struct {
	Status  string
	Data    []aff.Affirmation
	Count   int
	Message string
}

type GetAffirmationResponse struct {
	Status  string
	Data    aff.Affirmation
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

type GetCategoriesResponse struct {
	Status  string
	Data    []cat.Category
	Count   int
	Message string
}

type GetCategoryResponse struct {
	Status  string
	Data    cat.Category
	Count   int
	Message string
}

type CreateCategoryResponse struct {
	Status  string
	Data    string
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
