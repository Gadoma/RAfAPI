package domain

import (
	"errors"
)

var ErrorCreateAffirmationCommandInvalidId error = errors.New("Affirmation Id must be a valid ULID")
var ErrorCreateAffirmationCommandInvalidCategoryId error = errors.New("Affirmation CategoryId must be a valid ULID")
var ErrorCreateAffirmationCommandInvalidText error = errors.New("Affirmation Text cannot be empty")

var ErrorUpdateAffirmationCommandInvalidCategoryId error = errors.New("Affirmation CategoryId must be a valid ULID")
var ErrorUpdateAffirmationCommandInvalidText error = errors.New("Affirmation Text cannot be empty")
