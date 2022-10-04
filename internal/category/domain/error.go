package domain

import (
	"errors"
)

var ErrorCreateCategoryCommandInvalidId error = errors.New("Category Id must be a valid ULID")
var ErrorCreateCategoryCommandInvalidName error = errors.New("Category Name cannot be empty")

var ErrorUpdateCategoryCommandInvalidName error = errors.New("Category Name cannot be empty")
