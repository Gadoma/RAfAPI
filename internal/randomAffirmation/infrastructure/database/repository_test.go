package database_test

import (
	"context"
	"testing"

	"github.com/gadoma/rafapi/internal/common/test"
	"github.com/gadoma/rafapi/internal/randomAffirmation/infrastructure/database"
	"github.com/oklog/ulid/v2"
)

func TestRandomAffirmationRepositoryGetRandomAffirmationsContent(t *testing.T) {
	test.PrepareTestDB()
	db := test.MustOpenDB(t)
	defer test.CleanupTestDB()
	defer test.MustCloseDB(t, db)

	s := database.NewRandomAffirmationRepository(db)

	ctx := context.Background()

	id, _ := ulid.Parse("01GEJ0CR9DWN7SA1QBSJE4DVKF")

	categoryIds := []ulid.ULID{id}

	if a, err := s.GetRandomAffirmations(ctx, categoryIds); err != nil {
		t.Error(err)
	} else if a[0].Text != "I am blessed." && a[0].Text != "My life is wonderful." {
		t.Errorf("Text=%v, want `I am blessed.` or `My life is wonderful.`", a[0].Text)
	}
}

func TestRandomAffirmationRepositoryGetRandomAffirmationsLength(t *testing.T) {
	test.PrepareTestDB()
	db := test.MustOpenDB(t)
	defer test.CleanupTestDB()
	defer test.MustCloseDB(t, db)

	s := database.NewRandomAffirmationRepository(db)

	ctx := context.Background()

	id1, _ := ulid.Parse("01GEJ0CR9DWN7SA1QBSJE4DVKF")
	id2, _ := ulid.Parse("01GEJ0CRM2JW0KY2Z4R5CH4349")
	id3, _ := ulid.Parse("01GEJ0CRYJ1AAGQZDS9BR13AKS")

	categoryIds := []ulid.ULID{id1, id2, id3}

	if a, err := s.GetRandomAffirmations(ctx, categoryIds); err != nil {
		t.Error(err)
	} else if got, want := len(a), len(categoryIds); got != want {
		t.Errorf("len=%v, want %v", got, want)
	}
}
