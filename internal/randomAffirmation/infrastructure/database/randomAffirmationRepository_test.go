package database_test

import (
	"context"
	"testing"

	"github.com/gadoma/rafapi/internal/randomAffirmation/infrastructure/database"
	"github.com/gadoma/rafapi/test"
)

func TestRandomAffirmationRepositoryGetRandomAffirmationsContent(t *testing.T) {
	test.PrepareTestDB()
	db := test.MustOpenDB(t)
	defer test.CleanupTestDB()
	defer test.MustCloseDB(t, db)

	s := database.NewRandomAffirmationRepository(db)

	ctx := context.Background()

	categoryIds := []int{1}

	if a, err := s.GetRandomAffirmations(ctx, categoryIds); err != nil {
		t.Error(err)
	} else if a[0].Text != "I am blessed." && a[0].Text != "My life is wonderful." {
		t.Errorf("Text=%v, want `I am blessed.` or `My life is wonderful.`", a[0].Text)
	}
}

func TestRandomAffirmationRepositoryGetRandomAffirmationsLenght(t *testing.T) {
	test.PrepareTestDB()
	db := test.MustOpenDB(t)
	defer test.CleanupTestDB()
	defer test.MustCloseDB(t, db)

	s := database.NewRandomAffirmationRepository(db)

	ctx := context.Background()

	categoryIds := []int{1, 2, 3}

	if a, err := s.GetRandomAffirmations(ctx, categoryIds); err != nil {
		t.Error(err)
	} else if got, want := len(a), len(categoryIds); got != want {
		t.Errorf("len=%v, want %v", got, want)
	}
}
