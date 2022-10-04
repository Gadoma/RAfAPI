package mock

import (
	"context"

	"github.com/gadoma/rafapi/internal/category/domain"
	"github.com/oklog/ulid/v2"
)

var _ domain.CategoryRepository = (*CategoryRepository)(nil)

type CategoryRepository struct {
	GetCategoryFn    func(ctx context.Context, id ulid.ULID) (*domain.Category, error)
	GetCategoriesFn  func(ctx context.Context) ([]*domain.Category, int, error)
	CreateCategoryFn func(ctx context.Context, au *domain.CreateCategoryCommand) error
	UpdateCategoryFn func(ctx context.Context, id ulid.ULID, au *domain.UpdateCategoryCommand) error
	DeleteCategoryFn func(ctx context.Context, id ulid.ULID) error
}

func (r *CategoryRepository) GetCategories(ctx context.Context) ([]*domain.Category, int, error) {
	return r.GetCategoriesFn(ctx)
}

func (r *CategoryRepository) GetCategory(ctx context.Context, id ulid.ULID) (*domain.Category, error) {
	return r.GetCategoryFn(ctx, id)
}

func (r *CategoryRepository) CreateCategory(ctx context.Context, cac *domain.CreateCategoryCommand) error {
	return r.CreateCategoryFn(ctx, cac)
}

func (r *CategoryRepository) UpdateCategory(ctx context.Context, id ulid.ULID, uac *domain.UpdateCategoryCommand) error {
	return r.UpdateCategoryFn(ctx, id, uac)
}

func (r *CategoryRepository) DeleteCategory(ctx context.Context, id ulid.ULID) error {
	return r.DeleteCategoryFn(ctx, id)
}
