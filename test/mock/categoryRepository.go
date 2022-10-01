package mock

import (
	"context"

	"github.com/gadoma/rafapi/internal/domain"
)

var _ domain.CategoryRepository = (*CategoryRepository)(nil)

type CategoryRepository struct {
	GetCategoryFn    func(ctx context.Context, id int) (*domain.Category, error)
	GetCategoriesFn  func(ctx context.Context) ([]*domain.Category, int, error)
	CreateCategoryFn func(ctx context.Context, cu *domain.CategoryUpdate) (int, error)
	UpdateCategoryFn func(ctx context.Context, id int, cu *domain.CategoryUpdate) error
	DeleteCategoryFn func(ctx context.Context, id int) error
}

func (r *CategoryRepository) GetCategories(ctx context.Context) ([]*domain.Category, int, error) {
	return r.GetCategoriesFn(ctx)
}

func (r *CategoryRepository) GetCategory(ctx context.Context, id int) (*domain.Category, error) {
	return r.GetCategoryFn(ctx, id)
}

func (r *CategoryRepository) CreateCategory(ctx context.Context, cu *domain.CategoryUpdate) (int, error) {
	return r.CreateCategoryFn(ctx, cu)
}

func (r *CategoryRepository) UpdateCategory(ctx context.Context, id int, cu *domain.CategoryUpdate) error {
	return r.UpdateCategoryFn(ctx, id, cu)
}

func (r *CategoryRepository) DeleteCategory(ctx context.Context, id int) error {
	return r.DeleteCategoryFn(ctx, id)
}
