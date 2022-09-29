package domain

import (
	"context"
	"errors"
)

type Category struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type CategoryUpdate struct {
	Name string `json:"name"`
}

var ErrorCategoryUpdateInvalidName error = errors.New("Category Name cannot be empty")

func (cu *CategoryUpdate) Validate() error {
	if cu.Name == "" {
		return ErrorCategoryUpdateInvalidName
	}

	return nil
}

type CategoriesService interface {
	GetCategories(ctx context.Context) ([]*Category, int, error)
	GetCategory(ctx context.Context, id int) (*Category, error)
	CreateCategory(ctx context.Context, cu *CategoryUpdate) (int, error)
	UpdateCategory(ctx context.Context, id int, cu *CategoryUpdate) error
	DeleteCategory(ctx context.Context, id int) error
}
