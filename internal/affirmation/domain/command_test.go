package domain_test

import (
	"errors"
	"testing"

	"github.com/gadoma/rafapi/internal/affirmation/domain"
	"github.com/oklog/ulid/v2"
)

func TestUpdateAffirmationCommandValidate(t *testing.T) {
	uac := domain.UpdateAffirmationCommand{
		Text:       "some text",
		CategoryId: ulid.Make(),
	}

	if err := uac.Validate(); err != nil {
		t.Error("Expected no validation errors")
	}
}

func TestUpdateAffirmationCommandValidateTextError(t *testing.T) {
	uac := domain.UpdateAffirmationCommand{
		Text:       "",
		CategoryId: ulid.Make(),
	}

	if err := uac.Validate(); !errors.Is(err, domain.ErrorUpdateAffirmationCommandInvalidText) {
		t.Error("Expected invalid Text validation error")
	}
}

func TestUpdateAffirmationCommandValidateCategoryIdError(t *testing.T) {
	uac := domain.UpdateAffirmationCommand{
		Text: "some text",
	}

	if err := uac.Validate(); !errors.Is(err, domain.ErrorUpdateAffirmationCommandInvalidCategoryId) {
		t.Error("Expected invalid CategoryId validation error")
	}
}

func TestCreateAffirmationCommandValidate(t *testing.T) {
	cac := domain.CreateAffirmationCommand{
		Id:         ulid.Make(),
		Text:       "some text",
		CategoryId: ulid.Make(),
	}

	if err := cac.Validate(); err != nil {
		t.Error("Expected no validation errors")
	}
}

func TestCreateAffirmationCommandValidateIdError(t *testing.T) {
	cac := domain.CreateAffirmationCommand{
		Text:       "some text",
		CategoryId: ulid.Make(),
	}

	if err := cac.Validate(); !errors.Is(err, domain.ErrorCreateAffirmationCommandInvalidId) {
		t.Error("Expected invalid Id validation error")
	}
}

func TestCreateAffirmationCommandValidateTextError(t *testing.T) {
	cac := domain.CreateAffirmationCommand{
		Id:         ulid.Make(),
		Text:       "",
		CategoryId: ulid.Make(),
	}

	if err := cac.Validate(); !errors.Is(err, domain.ErrorCreateAffirmationCommandInvalidText) {
		t.Error("Expected invalid Text validation error")
	}
}

func TestCreateAffirmationCommandValidateCategoryIdError(t *testing.T) {
	cac := domain.CreateAffirmationCommand{
		Id:   ulid.Make(),
		Text: "some text",
	}

	if err := cac.Validate(); !errors.Is(err, domain.ErrorCreateAffirmationCommandInvalidCategoryId) {
		t.Error("Expected invalid CategoryId validation error")
	}
}
