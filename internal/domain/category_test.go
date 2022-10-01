package domain_test

import (
	"testing"

	"github.com/gadoma/rafapi/internal/domain"
)

func TestCategoryUpdateValidate(t *testing.T) {
	cu := domain.CategoryUpdate{
		Name: "some name",
	}

	if err := cu.Validate(); err != nil {
		t.Error("Expected no validation errors")
	}
}

func TestCategoryUpdateValidateNameError(t *testing.T) {
	cu := domain.CategoryUpdate{
		Name: "",
	}

	if err := cu.Validate(); err != domain.ErrorCategoryUpdateInvalidName {
		t.Error("Expected invalid category name error")
	}
}
