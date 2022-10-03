package database_test

import (
	"context"
	"testing"

	"github.com/gadoma/rafapi/internal/category/domain"
	"github.com/gadoma/rafapi/internal/category/infrastructure/database"
	"github.com/gadoma/rafapi/test"
)

func TestCategoryRepositoryGetCategories(t *testing.T) {
	test.PrepareTestDB()
	db := test.MustOpenDB(t)
	defer test.CleanupTestDB()
	defer test.MustCloseDB(t, db)

	s := database.NewCategoryRepository(db)

	ctx := context.Background()

	if a, n, err := s.GetCategories(ctx); err != nil {
		t.Error(err)
	} else if got, want := a[0].Id, 1; got != want {
		t.Errorf("a[0].Id=%v, want %v", got, want)
	} else if got, want := len(a), 5; got != want {
		t.Errorf("len=%v, want %v", got, want)
	} else if got, want := n, 5; got != want {
		t.Errorf("n=%v, want %v", got, want)
	}
}

func TestCategoryRepositoryGetCategory(t *testing.T) {
	test.PrepareTestDB()
	db := test.MustOpenDB(t)
	defer test.CleanupTestDB()
	defer test.MustCloseDB(t, db)

	s := database.NewCategoryRepository(db)

	ctx := context.Background()

	if a, err := s.GetCategory(ctx, 1); err != nil {
		t.Error(err)
	} else if got, want := a.Id, 1; got != want {
		t.Errorf("Id=%v, want %v", got, want)
	}
}

func TestCategoryRepositoryCreateCategory(t *testing.T) {
	test.PrepareTestDB()
	db := test.MustOpenDB(t)
	defer test.CleanupTestDB()
	defer test.MustCloseDB(t, db)

	newCategoryName := "Category"

	s := database.NewCategoryRepository(db)

	ctx := context.Background()

	au := &domain.CategoryUpdate{
		Name: newCategoryName,
	}

	id, err := s.CreateCategory(ctx, au)

	if err != nil {
		t.Error(err)
	}

	a, err := s.GetCategory(ctx, id)

	if err != nil {
		t.Error(err)
	}

	if got, want := a.Name, newCategoryName; got != want {
		t.Errorf("Name=%v, want %v", got, want)
	}
}

func TestCategoryRepositoryUpdateCategory(t *testing.T) {
	test.PrepareTestDB()
	db := test.MustOpenDB(t)
	defer test.CleanupTestDB()
	defer test.MustCloseDB(t, db)

	categoryId := 1
	newCategoryName := "Category"

	s := database.NewCategoryRepository(db)

	ctx := context.Background()

	au := &domain.CategoryUpdate{
		Name: newCategoryName,
	}

	err := s.UpdateCategory(ctx, categoryId, au)

	if err != nil {
		t.Error(err)
	}

	a, err := s.GetCategory(ctx, categoryId)

	if err != nil {
		t.Error(err)
	}

	if got, want := a.Name, newCategoryName; got != want {
		t.Errorf("Name=%v, want %v", got, want)
	}
}

func TestCategoryRepositoryDeleteCategory(t *testing.T) {
	test.PrepareTestDB()
	db := test.MustOpenDB(t)
	defer test.CleanupTestDB()
	defer test.MustCloseDB(t, db)

	s := database.NewCategoryRepository(db)

	ctx := context.Background()

	categoryId, _ := s.CreateCategory(ctx, &domain.CategoryUpdate{Name: "To be deleted"})
	err := s.DeleteCategory(ctx, categoryId)

	if err != nil {
		t.Error(err)
	}

	id, err := s.GetCategory(ctx, categoryId)

	if id != nil || err != nil {
		t.Errorf("Record was not deleted")
	}
}
