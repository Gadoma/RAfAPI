package database_test

import (
	"context"
	"testing"

	"github.com/gadoma/rafapi/internal/category/domain"
	"github.com/gadoma/rafapi/internal/category/infrastructure/database"
	"github.com/gadoma/rafapi/test"
	"github.com/oklog/ulid/v2"
)

var categoryIdString string = "01GEJ0CR9DWN7SA1QBSJE4DVKF"

func TestCategoryRepositoryGetCategories(t *testing.T) {
	test.PrepareTestDB()
	db := test.MustOpenDB(t)
	defer test.CleanupTestDB()
	defer test.MustCloseDB(t, db)

	s := database.NewCategoryRepository(db)

	ctx := context.Background()

	if a, n, err := s.GetCategories(ctx); err != nil {
		t.Error(err)
	} else if got, want := a[0].Id.String(), categoryIdString; got != want {
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

	id, _ := ulid.Parse(categoryIdString)
	if a, err := s.GetCategory(ctx, id); err != nil {
		t.Error(err)
	} else if got, want := a.Id, id; got != want {
		t.Errorf("Id=%v, want %v", got, want)
	}
}

func TestCategoryRepositoryCreateCategory(t *testing.T) {
	test.PrepareTestDB()
	db := test.MustOpenDB(t)
	defer test.CleanupTestDB()
	defer test.MustCloseDB(t, db)

	newCategoryId := ulid.Make()
	newCategoryName := "I am a category."

	s := database.NewCategoryRepository(db)

	ctx := context.Background()

	ccc := &domain.CreateCategoryCommand{
		Id:   newCategoryId,
		Name: newCategoryName,
	}

	err := s.CreateCategory(ctx, ccc)

	if err != nil {
		t.Error(err)
	}

	a, err := s.GetCategory(ctx, newCategoryId)

	if err != nil {
		t.Error(err)
	}

	if got, want := a.Id, newCategoryId; got != want {
		t.Errorf("Id=%v, want %v", got, want)
	} else if got, want := a.Name, newCategoryName; got != want {
		t.Errorf("Name=%v, want %v", got, want)
	}
}

func TestCategoryRepositoryUpdateCategory(t *testing.T) {
	test.PrepareTestDB()
	db := test.MustOpenDB(t)
	defer test.CleanupTestDB()
	defer test.MustCloseDB(t, db)

	categoryId, _ := ulid.Parse(categoryIdString)

	newCategoryName := "Love"

	s := database.NewCategoryRepository(db)

	ctx := context.Background()

	ucc := &domain.UpdateCategoryCommand{
		Name: newCategoryName,
	}

	err := s.UpdateCategory(ctx, categoryId, ucc)

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

	categoryId, _ := ulid.Parse("01GEJ0CSKKVPY3PR8VXJMPDQAY")

	s := database.NewCategoryRepository(db)

	ctx := context.Background()

	err := s.DeleteCategory(ctx, categoryId)

	if err != nil {
		t.Error(err)
	}

	id, err := s.GetCategory(ctx, categoryId)

	if id != nil || err != nil {
		t.Errorf("Record was not deleted")
	}
}
