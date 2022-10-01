package application_test

import (
	"context"
	"errors"
	"testing"

	"github.com/gadoma/rafapi/internal/application"
	"github.com/gadoma/rafapi/internal/domain"
	"github.com/gadoma/rafapi/test/mock"
)

func prepareCategoryServiceTest() (
	repositoryMock mock.CategoryRepository,
	categoryStub domain.Category,
	categoryUpdateStub domain.CategoryUpdate,
	ctx context.Context) {
	repositoryMock = mock.CategoryRepository{}

	categoryStub = domain.Category{
		Id:   1,
		Name: "Stubs",
	}

	categoryUpdateStub = domain.CategoryUpdate{
		Name: "Other stubs",
	}

	ctx = context.Background()

	return
}

func TestCategoryServiceGetCategories(t *testing.T) {
	repositoryMock, categoryStub, _, ctx := prepareCategoryServiceTest()

	repositoryMock.GetCategoriesFn = func(ctx context.Context) ([]*domain.Category, int, error) {
		return []*domain.Category{&categoryStub}, 1, nil
	}

	sut := application.NewCategoryService(&repositoryMock)

	result, n, err := sut.GetCategories(ctx)

	if err != nil {
		t.Errorf("error=%q, want nil", err)
	}

	if got, want := result[0].Id, categoryStub.Id; got != want {
		t.Errorf("result[0].Id=%v, want %v", got, want)
	} else if got, want := result[0].Name, categoryStub.Name; got != want {
		t.Errorf("result[0].Name=%v, want %v", got, want)
	} else if got, want := n, 1; got != want {
		t.Errorf("count=%v, want %v", got, want)
	}
}

func TestCategoryServiceGetCategoriesError(t *testing.T) {
	repositoryMock, _, _, ctx := prepareCategoryServiceTest()

	repositoryMock.GetCategoriesFn = func(ctx context.Context) ([]*domain.Category, int, error) {
		return nil, 0, errors.New("something went wrong")
	}

	sut := application.NewCategoryService(&repositoryMock)

	_, _, err := sut.GetCategories(ctx)

	if err == nil {
		t.Error("an error was expected")
	}
}

func TestCategoryServiceGetCategory(t *testing.T) {
	repositoryMock, categoryStub, _, ctx := prepareCategoryServiceTest()

	repositoryMock.GetCategoryFn = func(ctx context.Context, id int) (*domain.Category, error) {
		return &categoryStub, nil
	}

	sut := application.NewCategoryService(&repositoryMock)

	result, err := sut.GetCategory(ctx, categoryStub.Id)

	if err != nil {
		t.Errorf("error=%q, want nil", err)
	}

	if got, want := result.Id, categoryStub.Id; got != want {
		t.Errorf("result.Id=%v, want %v", got, want)
	} else if got, want := result.Name, categoryStub.Name; got != want {
		t.Errorf("result.Name=%v, want %v", got, want)
	}
}

func TestCategoryServiceGetCategoryNotFound(t *testing.T) {
	repositoryMock, categoryStub, _, ctx := prepareCategoryServiceTest()

	repositoryMock.GetCategoryFn = func(ctx context.Context, id int) (*domain.Category, error) {
		return nil, nil
	}

	sut := application.NewCategoryService(&repositoryMock)

	_, err := sut.GetCategory(ctx, categoryStub.Id)

	if err != domain.ErrorResourceNotFound {
		t.Errorf("error=%q, want domain.ErrorResourceNotFound", err)
	}
}

func TestCategoryServiceGetCategoryError(t *testing.T) {
	repositoryMock, categoryStub, _, ctx := prepareCategoryServiceTest()

	repositoryMock.GetCategoryFn = func(ctx context.Context, id int) (*domain.Category, error) {
		return nil, errors.New("something went wrong")
	}

	sut := application.NewCategoryService(&repositoryMock)

	_, err := sut.GetCategory(ctx, categoryStub.Id)

	if err == nil {
		t.Error("an error was expected")
	}
}

func TestCategoryServiceCreateCategory(t *testing.T) {
	repositoryMock, _, categoryUpdateStub, ctx := prepareCategoryServiceTest()

	repositoryMock.CreateCategoryFn = func(ctx context.Context, cu *domain.CategoryUpdate) (int, error) {
		return 1, nil
	}

	sut := application.NewCategoryService(&repositoryMock)

	result, err := sut.CreateCategory(ctx, &categoryUpdateStub)

	if err != nil {
		t.Errorf("error=%q, want nil", err)
	}

	if got, want := result, 1; got != want {
		t.Errorf("result=%v, want %v", got, want)
	}
}

func TestCategoryServiceCreateCategoryValidationError(t *testing.T) {
	repositoryMock, _, categoryUpdateStub, ctx := prepareCategoryServiceTest()

	categoryUpdateStub.Name = ""

	sut := application.NewCategoryService(&repositoryMock)

	_, err := sut.CreateCategory(ctx, &categoryUpdateStub)

	if err != domain.ErrorCategoryUpdateInvalidName {
		t.Errorf("error=%q, want domain.ErrorCategoryUpdateInvalidName", err)
	}
}

func TestCategoryServiceCreateCategoryError(t *testing.T) {
	repositoryMock, _, categoryUpdateStub, ctx := prepareCategoryServiceTest()

	repositoryMock.CreateCategoryFn = func(ctx context.Context, cu *domain.CategoryUpdate) (int, error) {
		return 0, errors.New("something went wrong")
	}

	sut := application.NewCategoryService(&repositoryMock)

	_, err := sut.CreateCategory(ctx, &categoryUpdateStub)

	if err == nil {
		t.Error("an error was expected")
	}
}

func TestCategoryServiceUpdateCategory(t *testing.T) {
	repositoryMock, _, categoryUpdateStub, ctx := prepareCategoryServiceTest()

	repositoryMock.UpdateCategoryFn = func(ctx context.Context, id int, cu *domain.CategoryUpdate) error {
		return nil
	}

	sut := application.NewCategoryService(&repositoryMock)

	err := sut.UpdateCategory(ctx, 1, &categoryUpdateStub)

	if err != nil {
		t.Errorf("error=%q, want nil", err)
	}
}

func TestCategoryServiceUpdateCategoryValidationError(t *testing.T) {
	repositoryMock, _, categoryUpdateStub, ctx := prepareCategoryServiceTest()

	categoryUpdateStub.Name = ""

	sut := application.NewCategoryService(&repositoryMock)

	err := sut.UpdateCategory(ctx, 1, &categoryUpdateStub)

	if err != domain.ErrorCategoryUpdateInvalidName {
		t.Errorf("error=%q, want domain.ErrorCategoryUpdateInvalidName", err)
	}
}

func TestCategoryServiceUpdateCategoryError(t *testing.T) {
	repositoryMock, _, categoryUpdateStub, ctx := prepareCategoryServiceTest()

	repositoryMock.UpdateCategoryFn = func(ctx context.Context, id int, cu *domain.CategoryUpdate) error {
		return errors.New("something went wrong")
	}

	sut := application.NewCategoryService(&repositoryMock)

	err := sut.UpdateCategory(ctx, 1, &categoryUpdateStub)

	if err == nil {
		t.Error("an error was expected")
	}
}

func TestCategoryServiceDeleteCategory(t *testing.T) {
	repositoryMock, _, _, ctx := prepareCategoryServiceTest()

	repositoryMock.DeleteCategoryFn = func(ctx context.Context, id int) error {
		return nil
	}

	sut := application.NewCategoryService(&repositoryMock)

	err := sut.DeleteCategory(ctx, 1)

	if err != nil {
		t.Errorf("error=%q, want nil", err)
	}
}

func TestCategoryServiceDeleteCategoryError(t *testing.T) {
	repositoryMock, _, _, ctx := prepareCategoryServiceTest()

	repositoryMock.DeleteCategoryFn = func(ctx context.Context, id int) error {
		return errors.New("something went wrong")
	}

	sut := application.NewCategoryService(&repositoryMock)

	err := sut.DeleteCategory(ctx, 1)

	if err == nil {
		t.Error("an error was expected")
	}
}
