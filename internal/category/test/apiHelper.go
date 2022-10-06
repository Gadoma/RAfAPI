package test

import "github.com/gadoma/rafapi/internal/category/domain"

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
