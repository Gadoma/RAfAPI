package domain

import (
	"github.com/oklog/ulid/v2"
)

type CreateCategoryCommand struct {
	Id   ulid.ULID
	Name string `json:"name"`
}

type UpdateCategoryCommand struct {
	Name string `json:"name"`
}

func (ccc *CreateCategoryCommand) Validate() error {
	if ccc.Id.String() == "00000000000000000000000000" {
		return ErrorCreateCategoryCommandInvalidId
	}

	if ccc.Name == "" {
		return ErrorCreateCategoryCommandInvalidName
	}

	return nil
}

func (ccc *UpdateCategoryCommand) Validate() error {
	if ccc.Name == "" {
		return ErrorUpdateCategoryCommandInvalidName
	}

	return nil
}
