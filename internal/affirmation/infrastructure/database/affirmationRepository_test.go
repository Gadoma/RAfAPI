package database_test

import (
	"context"
	"testing"

	"github.com/gadoma/rafapi/internal/affirmation/domain"
	"github.com/gadoma/rafapi/internal/affirmation/infrastructure/database"
	"github.com/gadoma/rafapi/test"
)

func TestAffirmationRepositoryGetAffirmations(t *testing.T) {
	test.PrepareTestDB()
	db := test.MustOpenDB(t)
	defer test.CleanupTestDB()
	defer test.MustCloseDB(t, db)

	s := database.NewAffirmationRepository(db)

	ctx := context.Background()

	if a, n, err := s.GetAffirmations(ctx); err != nil {
		t.Error(err)
	} else if got, want := a[0].Id, 1; got != want {
		t.Errorf("a[0].Id=%v, want %v", got, want)
	} else if got, want := len(a), 8; got != want {
		t.Errorf("len=%v, want %v", got, want)
	} else if got, want := n, 8; got != want {
		t.Errorf("n=%v, want %v", got, want)
	}
}

func TestAffirmationRepositoryGetAffirmation(t *testing.T) {
	test.PrepareTestDB()
	db := test.MustOpenDB(t)
	defer test.CleanupTestDB()
	defer test.MustCloseDB(t, db)

	s := database.NewAffirmationRepository(db)

	ctx := context.Background()

	if a, err := s.GetAffirmation(ctx, 1); err != nil {
		t.Error(err)
	} else if got, want := a.Id, 1; got != want {
		t.Errorf("Id=%v, want %v", got, want)
	}
}

func TestAffirmationRepositoryCreateAffirmation(t *testing.T) {
	test.PrepareTestDB()
	db := test.MustOpenDB(t)
	defer test.CleanupTestDB()
	defer test.MustCloseDB(t, db)

	newAffirmationText := "I am fantastic."
	newAffirmationCategoryId := 2

	s := database.NewAffirmationRepository(db)

	ctx := context.Background()

	au := &domain.AffirmationUpdate{
		Text:       newAffirmationText,
		CategoryId: newAffirmationCategoryId,
	}

	id, err := s.CreateAffirmation(ctx, au)

	if err != nil {
		t.Error(err)
	}

	a, err := s.GetAffirmation(ctx, id)

	if err != nil {
		t.Error(err)
	}

	if got, want := a.Text, newAffirmationText; got != want {
		t.Errorf("Text=%v, want %v", got, want)
	} else if got, want := a.CategoryId, newAffirmationCategoryId; got != want {
		t.Errorf("Text=%v, want %v", got, want)
	}
}

func TestAffirmationRepositoryUpdateAffirmation(t *testing.T) {
	test.PrepareTestDB()
	db := test.MustOpenDB(t)
	defer test.CleanupTestDB()
	defer test.MustCloseDB(t, db)

	affirmationId := 1
	newAffirmationText := "I am fantastic."
	newAffirmationCategoryId := 2

	s := database.NewAffirmationRepository(db)

	ctx := context.Background()

	au := &domain.AffirmationUpdate{
		Text:       newAffirmationText,
		CategoryId: newAffirmationCategoryId,
	}

	err := s.UpdateAffirmation(ctx, affirmationId, au)

	if err != nil {
		t.Error(err)
	}

	a, err := s.GetAffirmation(ctx, affirmationId)

	if err != nil {
		t.Error(err)
	}

	if got, want := a.Text, newAffirmationText; got != want {
		t.Errorf("Text=%v, want %v", got, want)
	} else if got, want := a.CategoryId, newAffirmationCategoryId; got != want {
		t.Errorf("Text=%v, want %v", got, want)
	}
}

func TestAffirmationRepositoryDeleteAffirmation(t *testing.T) {
	test.PrepareTestDB()
	db := test.MustOpenDB(t)
	defer test.CleanupTestDB()
	defer test.MustCloseDB(t, db)

	affirmationId := 1

	s := database.NewAffirmationRepository(db)

	ctx := context.Background()

	err := s.DeleteAffirmation(ctx, affirmationId)

	if err != nil {
		t.Error(err)
	}

	id, err := s.GetAffirmation(ctx, affirmationId)

	if id != nil || err != nil {
		t.Errorf("Record was not deleted")
	}
}
