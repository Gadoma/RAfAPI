package domain

import "errors"

var ErrorCreateAffirmationCommandInvalidId = errors.New("affirmation Id must be a valid ULID")
var ErrorCreateAffirmationCommandInvalidCategoryId = errors.New("affirmation CategoryId must be a valid ULID")
var ErrorCreateAffirmationCommandInvalidText = errors.New("affirmation Text cannot be empty")

var ErrorUpdateAffirmationCommandInvalidCategoryId = errors.New("affirmation CategoryId must be a valid ULID")
var ErrorUpdateAffirmationCommandInvalidText = errors.New("affirmation Text cannot be empty")
