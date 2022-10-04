package domain

import (
	"context"

	"github.com/oklog/ulid/v2"
)

type CategoryService interface {
	GetCategories(ctx context.Context) ([]*Category, int, error)
	GetCategory(ctx context.Context, id ulid.ULID) (*Category, error)
	CreateCategory(ctx context.Context, ccc *CreateCategoryCommand) (*ulid.ULID, error)
	UpdateCategory(ctx context.Context, id ulid.ULID, ccc *UpdateCategoryCommand) error
	DeleteCategory(ctx context.Context, id ulid.ULID) error
}

type CategoryRepository interface {
	GetCategories(ctx context.Context) ([]*Category, int, error)
	GetCategory(ctx context.Context, id ulid.ULID) (*Category, error)
	CreateCategory(ctx context.Context, ccc *CreateCategoryCommand) error
	UpdateCategory(ctx context.Context, id ulid.ULID, ccc *UpdateCategoryCommand) error
	DeleteCategory(ctx context.Context, id ulid.ULID) error
}
