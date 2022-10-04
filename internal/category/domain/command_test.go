package domain_test

import (
	"errors"
	"testing"

	"github.com/gadoma/rafapi/internal/category/domain"
	"github.com/oklog/ulid/v2"
)

func TestUpdateCategoryCommandValidate(t *testing.T) {
	ccc := domain.UpdateCategoryCommand{
		Name: "some name",
	}

	if err := ccc.Validate(); err != nil {
		t.Error("Expected no validation errors")
	}
}

func TestUpdateCategoryCommandValidateNameError(t *testing.T) {
	ccc := domain.UpdateCategoryCommand{
		Name: "",
	}

	if err := ccc.Validate(); !errors.Is(err, domain.ErrorUpdateCategoryCommandInvalidName) {
		t.Error("Expected invalid Name validation error")
	}
}

func TestCreateCategoryCommandValidate(t *testing.T) {
	ccc := domain.CreateCategoryCommand{
		Id:   ulid.Make(),
		Name: "some name",
	}

	if err := ccc.Validate(); err != nil {
		t.Error("Expected no validation errors")
	}
}

func TestCreateCategoryCommandValidateIdError(t *testing.T) {
	ccc := domain.CreateCategoryCommand{
		Name: "some name",
	}

	if err := ccc.Validate(); !errors.Is(err, domain.ErrorCreateCategoryCommandInvalidId) {
		t.Error("Expected invalid Id validation error")
	}
}

func TestCreateCategoryCommandValidateNameError(t *testing.T) {
	ccc := domain.CreateCategoryCommand{
		Id:   ulid.Make(),
		Name: "",
	}

	if err := ccc.Validate(); !errors.Is(err, domain.ErrorCreateCategoryCommandInvalidName) {
		t.Error("Expected invalid Name validation error")
	}
}
