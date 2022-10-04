package application

import (
	"context"
	"fmt"

	"github.com/gadoma/rafapi/internal/category/domain"
	common "github.com/gadoma/rafapi/internal/common/domain"
	"github.com/oklog/ulid/v2"
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
		return nil, 0, fmt.Errorf("an error occurred while getting all Categories: %w", err)
	}

	return a, n, nil
}

func (s *CategoryService) GetCategory(ctx context.Context, id ulid.ULID) (*domain.Category, error) {
	a, err := s.repo.GetCategory(ctx, id)

	if err != nil {
		return nil, fmt.Errorf("an error occurred while getting the Category by id: %w", err)
	}

	if a == nil {
		return nil, common.ErrorResourceNotFound
	}

	return a, nil
}

func (s *CategoryService) CreateCategory(ctx context.Context, ccc *domain.CreateCategoryCommand) (*ulid.ULID, error) {
	if err := ccc.Validate(); err != nil {
		return nil, err
	}

	err := s.repo.CreateCategory(ctx, ccc)

	if err != nil {
		return nil, fmt.Errorf("an error occurred while creating the Category: %w", err)
	}

	return &ccc.Id, nil
}

func (s *CategoryService) UpdateCategory(ctx context.Context, id ulid.ULID, ccc *domain.UpdateCategoryCommand) error {
	if err := ccc.Validate(); err != nil {
		return err
	}

	if err := s.repo.UpdateCategory(ctx, id, ccc); err != nil {
		return fmt.Errorf("an error occurred while updating the Category: %w", err)
	}

	return nil
}

func (s *CategoryService) DeleteCategory(ctx context.Context, id ulid.ULID) error {
	if err := s.repo.DeleteCategory(ctx, id); err != nil {
		return fmt.Errorf("an error occurred while deleting the Category: %w", err)
	}

	return nil
}
