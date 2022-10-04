package application_test

import (
	"context"
	"errors"
	"testing"

	"github.com/gadoma/rafapi/internal/category/application"
	"github.com/gadoma/rafapi/internal/category/domain"
	common "github.com/gadoma/rafapi/internal/common/domain"
	"github.com/gadoma/rafapi/test/mock"
	"github.com/oklog/ulid/v2"
)

func getCategoryEntityStub() domain.Category {
	return domain.Category{
		Id:   ulid.Make(),
		Name: "I am a stub.",
	}
}

func getCreateCategoryCommandStub() domain.CreateCategoryCommand {
	return domain.CreateCategoryCommand{
		Id:   ulid.Make(),
		Name: "I am a create stub.",
	}
}

func getUpdateCategoryCommandStub() domain.UpdateCategoryCommand {
	return domain.UpdateCategoryCommand{
		Name: "I am an update stub.",
	}
}

func TestCategoryServiceGetCategories(t *testing.T) {
	repositoryMock := mock.CategoryRepository{}
	ctx := context.Background()
	categoryStub := getCategoryEntityStub()

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
	repositoryMock := mock.CategoryRepository{}
	ctx := context.Background()

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
	repositoryMock := mock.CategoryRepository{}
	ctx := context.Background()
	categoryStub := getCategoryEntityStub()

	repositoryMock.GetCategoryFn = func(ctx context.Context, id ulid.ULID) (*domain.Category, error) {
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
	repositoryMock := mock.CategoryRepository{}
	ctx := context.Background()
	categoryStub := getCategoryEntityStub()

	repositoryMock.GetCategoryFn = func(ctx context.Context, id ulid.ULID) (*domain.Category, error) {
		return nil, nil
	}

	sut := application.NewCategoryService(&repositoryMock)

	_, err := sut.GetCategory(ctx, categoryStub.Id)

	if !errors.Is(err, common.ErrorResourceNotFound) {
		t.Errorf("error=%q, want domain.ErrorResourceNotFound", err)
	}
}

func TestCategoryServiceGetCategoryError(t *testing.T) {
	repositoryMock := mock.CategoryRepository{}
	ctx := context.Background()
	categoryStub := getCategoryEntityStub()

	repositoryMock.GetCategoryFn = func(ctx context.Context, id ulid.ULID) (*domain.Category, error) {
		return nil, errors.New("something went wrong")
	}

	sut := application.NewCategoryService(&repositoryMock)

	_, err := sut.GetCategory(ctx, categoryStub.Id)

	if err == nil {
		t.Error("an error was expected")
	}
}

func TestCategoryServiceCreateCategory(t *testing.T) {
	repositoryMock := mock.CategoryRepository{}
	ctx := context.Background()
	createCategoryStub := getCreateCategoryCommandStub()

	repositoryMock.CreateCategoryFn = func(ctx context.Context, ccc *domain.CreateCategoryCommand) error {
		return nil
	}

	sut := application.NewCategoryService(&repositoryMock)

	result, err := sut.CreateCategory(ctx, &createCategoryStub)

	if err != nil {
		t.Errorf("error=%q, want nil", err)
	}

	if got, want := result, &createCategoryStub.Id; got != want {
		t.Errorf("result=%v, want %v", got, want)
	}
}

func TestCategoryServiceCreateCategoryValidationError(t *testing.T) {
	repositoryMock := mock.CategoryRepository{}
	ctx := context.Background()
	createCategoryStub := getCreateCategoryCommandStub()

	createCategoryStub.Name = ""

	sut := application.NewCategoryService(&repositoryMock)

	_, err := sut.CreateCategory(ctx, &createCategoryStub)

	if !errors.Is(err, domain.ErrorCreateCategoryCommandInvalidName) {
		t.Errorf("error=%q, want domain.ErrorCreateCategoryCommandInvalidName", err)
	}
}

func TestCategoryServiceCreateCategoryError(t *testing.T) {
	repositoryMock := mock.CategoryRepository{}
	ctx := context.Background()
	createCategoryStub := getCreateCategoryCommandStub()

	repositoryMock.CreateCategoryFn = func(ctx context.Context, ccc *domain.CreateCategoryCommand) error {
		return errors.New("something went wrong")
	}

	sut := application.NewCategoryService(&repositoryMock)

	_, err := sut.CreateCategory(ctx, &createCategoryStub)

	if err == nil {
		t.Error("an error was expected")
	}
}

func TestCategoryServiceUpdateCategory(t *testing.T) {
	repositoryMock := mock.CategoryRepository{}
	ctx := context.Background()
	updateCategoryStub := getUpdateCategoryCommandStub()

	repositoryMock.UpdateCategoryFn = func(ctx context.Context, id ulid.ULID, ccc *domain.UpdateCategoryCommand) error {
		return nil
	}

	sut := application.NewCategoryService(&repositoryMock)

	err := sut.UpdateCategory(ctx, ulid.Make(), &updateCategoryStub)

	if err != nil {
		t.Errorf("error=%q, want nil", err)
	}
}

func TestCategoryServiceUpdateCategoryValidationError(t *testing.T) {
	repositoryMock := mock.CategoryRepository{}
	ctx := context.Background()
	updateCategoryStub := getUpdateCategoryCommandStub()

	updateCategoryStub.Name = ""

	sut := application.NewCategoryService(&repositoryMock)

	err := sut.UpdateCategory(ctx, ulid.Make(), &updateCategoryStub)

	if !errors.Is(err, domain.ErrorUpdateCategoryCommandInvalidName) {
		t.Errorf("error=%q, want domain.ErrorCategoryUpdateInvalidName", err)
	}
}

func TestCategoryServiceUpdateCategoryError(t *testing.T) {
	repositoryMock := mock.CategoryRepository{}
	ctx := context.Background()
	updateCategoryStub := getUpdateCategoryCommandStub()

	repositoryMock.UpdateCategoryFn = func(ctx context.Context, id ulid.ULID, ccc *domain.UpdateCategoryCommand) error {
		return errors.New("something went wrong")
	}

	sut := application.NewCategoryService(&repositoryMock)

	err := sut.UpdateCategory(ctx, ulid.Make(), &updateCategoryStub)

	if err == nil {
		t.Error("an error was expected")
	}
}

func TestCategoryServiceDeleteCategory(t *testing.T) {
	repositoryMock := mock.CategoryRepository{}
	ctx := context.Background()

	repositoryMock.DeleteCategoryFn = func(ctx context.Context, id ulid.ULID) error {
		return nil
	}

	sut := application.NewCategoryService(&repositoryMock)

	err := sut.DeleteCategory(ctx, ulid.Make())

	if err != nil {
		t.Errorf("error=%q, want nil", err)
	}
}

func TestCategoryServiceDeleteCategoryError(t *testing.T) {
	repositoryMock := mock.CategoryRepository{}
	ctx := context.Background()

	repositoryMock.DeleteCategoryFn = func(ctx context.Context, id ulid.ULID) error {
		return errors.New("something went wrong")
	}

	sut := application.NewCategoryService(&repositoryMock)

	err := sut.DeleteCategory(ctx, ulid.Make())

	if err == nil {
		t.Error("an error was expected")
	}
}
