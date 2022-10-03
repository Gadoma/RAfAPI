package domain_test

import (
	"errors"
	"testing"

	"github.com/gadoma/rafapi/internal/affirmation/domain"
)

func TestAffirmationUpdateValidate(t *testing.T) {
	au := domain.AffirmationUpdate{
		Text:       "some text",
		CategoryId: 1,
	}

	if err := au.Validate(); err != nil {
		t.Error("Expected no validation errors")
	}
}

func TestAffirmationUpdateValidateTextError(t *testing.T) {
	au := domain.AffirmationUpdate{
		Text:       "",
		CategoryId: 1,
	}

	if err := au.Validate(); !errors.Is(err, domain.ErrorAffirmationUpdateInvalidText) {
		t.Error("Expected invalid text validation error")
	}
}

func TestAffirmationUpdateValidateCategoryIdError(t *testing.T) {
	var tests = []int{0, -1}

	for _, test := range tests {
		au := domain.AffirmationUpdate{
			Text:       "some text",
			CategoryId: test,
		}

		if err := au.Validate(); !errors.Is(err, domain.ErrorAffirmationUpdateInvalidCategoryId) {
			t.Error("Expected invalid categoryId validation error")
		}
	}
}
