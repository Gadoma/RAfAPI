package domain

import "errors"

var ErrorCreateCategoryCommandInvalidId = errors.New("category Id must be a valid ULID")
var ErrorCreateCategoryCommandInvalidName = errors.New("category Name cannot be empty")

var ErrorUpdateCategoryCommandInvalidName = errors.New("category Name cannot be empty")
