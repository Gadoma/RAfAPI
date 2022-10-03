package application

import (
	"context"
	"fmt"

	"github.com/gadoma/rafapi/internal/category/domain"
	common "github.com/gadoma/rafapi/internal/common/domain"
)

var _ domain.CategoryService = (*CategoryService)(nil)

type CategoryService struct {
	repo domain.CategoryRepository
}

func NewCategoryService(repo domain.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) GetCategories(ctx context.Context) ([]*domain.Category, int, error) {
	a, n, err := s.repo.GetCategories(ctx)

	if err != nil {
		return nil, 0, fmt.Errorf("there was an error getting all Categories: %w", err)
	}

	return a, n, nil
}

func (s *CategoryService) GetCategory(ctx context.Context, id int) (*domain.Category, error) {
	a, err := s.repo.GetCategory(ctx, id)

	if err != nil {
		return nil, fmt.Errorf("there was an error getting the Category by id: %w", err)
	}

	if a == nil {
		return nil, common.ErrorResourceNotFound
	}

	return a, nil
}

func (s *CategoryService) CreateCategory(ctx context.Context, cu *domain.CategoryUpdate) (int, error) {
	if err := cu.Validate(); err != nil {
		return 0, err
	}

	id, err := s.repo.CreateCategory(ctx, cu)

	if err != nil {
		return 0, fmt.Errorf("there was an error creating the Category: %w", err)
	}

	return id, nil
}

func (s *CategoryService) UpdateCategory(ctx context.Context, id int, cu *domain.CategoryUpdate) error {
	if err := cu.Validate(); err != nil {
		return err
	}

	if err := s.repo.UpdateCategory(ctx, id, cu); err != nil {
		return fmt.Errorf("there was an error updating the Category: %w", err)
	}

	return nil
}

func (s *CategoryService) DeleteCategory(ctx context.Context, id int) error {
	if err := s.repo.DeleteCategory(ctx, id); err != nil {
		return fmt.Errorf("there was an error deleting the Category: %w", err)
	}

	return nil
}
