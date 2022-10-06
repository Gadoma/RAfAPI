package database_test

import (
	"context"
	"testing"

	"github.com/gadoma/rafapi/internal/affirmation/domain"
	"github.com/gadoma/rafapi/internal/affirmation/infrastructure/database"
	"github.com/gadoma/rafapi/internal/common/test"
	"github.com/oklog/ulid/v2"
)

var affirmationIdString = "01GEJ0CNNA3VXV1HMJCKFNCYJV"
var categoryIdString = "01GEJ0CRM2JW0KY2Z4R5CH4349"

func TestAffirmationRepositoryGetAffirmations(t *testing.T) {
	test.PrepareTestDB()
	db := test.MustOpenDB(t)
	defer test.CleanupTestDB()
	defer test.MustCloseDB(t, db)

	s := database.NewAffirmationRepository(db)

	ctx := context.Background()

	if a, n, err := s.GetAffirmations(ctx); err != nil {
		t.Error(err)
	} else if got, want := a[0].Id.String(), affirmationIdString; got != want {
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

	id, _ := ulid.Parse(affirmationIdString)
	if a, err := s.GetAffirmation(ctx, id); err != nil {
		t.Error(err)
	} else if got, want := a.Id, id; got != want {
		t.Errorf("Id=%v, want %v", got, want)
	}
}

func TestAffirmationRepositoryCreateAffirmation(t *testing.T) {
	test.PrepareTestDB()
	db := test.MustOpenDB(t)
	defer test.CleanupTestDB()
	defer test.MustCloseDB(t, db)

	newAffirmationId := ulid.Make()
	newAffirmationText := "I am fantastic."
	newAffirmationCategoryId, _ := ulid.Parse(categoryIdString)

	s := database.NewAffirmationRepository(db)

	ctx := context.Background()

	cac := &domain.CreateAffirmationCommand{
		Id:         newAffirmationId,
		Text:       newAffirmationText,
		CategoryId: newAffirmationCategoryId,
	}

	err := s.CreateAffirmation(ctx, cac)

	if err != nil {
		t.Error(err)
	}

	a, err := s.GetAffirmation(ctx, newAffirmationId)

	if err != nil {
		t.Error(err)
	}

	if got, want := a.Id, newAffirmationId; got != want {
		t.Errorf("Id=%v, want %v", got, want)
	} else if got, want := a.Text, newAffirmationText; got != want {
		t.Errorf("Text=%v, want %v", got, want)
	} else if got, want := a.CategoryId, newAffirmationCategoryId; got != want {
		t.Errorf("CategoryId=%v, want %v", got, want)
	}
}

func TestAffirmationRepositoryUpdateAffirmation(t *testing.T) {
	test.PrepareTestDB()
	db := test.MustOpenDB(t)
	defer test.CleanupTestDB()
	defer test.MustCloseDB(t, db)

	affirmationId, _ := ulid.Parse(affirmationIdString)

	newAffirmationText := "I am fantastic."
	newAffirmationCategoryId, _ := ulid.Parse(categoryIdString)

	s := database.NewAffirmationRepository(db)

	ctx := context.Background()

	uac := &domain.UpdateAffirmationCommand{
		Text:       newAffirmationText,
		CategoryId: newAffirmationCategoryId,
	}

	err := s.UpdateAffirmation(ctx, affirmationId, uac)

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
		t.Errorf("CategoryId=%v, want %v", got, want)
	}
}

func TestAffirmationRepositoryDeleteAffirmation(t *testing.T) {
	test.PrepareTestDB()
	db := test.MustOpenDB(t)
	defer test.CleanupTestDB()
	defer test.MustCloseDB(t, db)

	affirmationId, _ := ulid.Parse(affirmationIdString)

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
