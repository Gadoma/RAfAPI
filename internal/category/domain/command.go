package domain

import (
	"github.com/oklog/ulid/v2"
)

// type RafapiULID ulid.ULID

// func (r *RafapiULID) UnmarshalJson(b []byte) error {
// 	id, err := ulid.Parse(string(b[:]))
// 	if err != nil {
// 		return err
// 	}
// 	*r = RafapiULID(id)
// 	return nil
// }

// func (r *RafapiULID) MarshalJson(r RafapiULID) ([]byte, error) {

// }

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
