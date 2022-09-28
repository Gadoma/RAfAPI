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

func (c *CategoryUpdate) Validate() error {
	if c.Name == "" {
		return errors.New("Category Name cannot be empty")
	}

	return nil
}

type CategoriesService interface {
	GetCategories(ctx context.Context) ([]*Category, int, error)
	GetCategory(ctx context.Context, id int) (*Category, error)
	CreateCategory(ctx context.Context, au *CategoryUpdate) (int, error)
	UpdateCategory(ctx context.Context, id int, au *CategoryUpdate) error
	DeleteCategory(ctx context.Context, id int) error
}
