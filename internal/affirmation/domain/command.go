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

type CreateAffirmationCommand struct {
	Id         ulid.ULID
	CategoryId ulid.ULID `json:"categoryId"`
	Text       string    `json:"text"`
}

type UpdateAffirmationCommand struct {
	CategoryId ulid.ULID `json:"categoryId"`
	Text       string    `json:"text"`
}

func (cac *CreateAffirmationCommand) Validate() error {
	if cac.Id.String() == "00000000000000000000000000" {
		return ErrorCreateAffirmationCommandInvalidId
	}

	if cac.CategoryId.String() == "00000000000000000000000000" {
		return ErrorCreateAffirmationCommandInvalidCategoryId
	}

	if cac.Text == "" {
		return ErrorCreateAffirmationCommandInvalidText
	}

	return nil
}

func (uac *UpdateAffirmationCommand) Validate() error {
	if uac.CategoryId.String() == "00000000000000000000000000" {
		return ErrorUpdateAffirmationCommandInvalidCategoryId
	}

	if uac.Text == "" {
		return ErrorUpdateAffirmationCommandInvalidText
	}

	return nil
}
