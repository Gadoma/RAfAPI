package domain

import (
	"github.com/oklog/ulid/v2"
)

type Category struct {
	Id   ulid.ULID `json:"id"`
	Name string    `json:"name"`
}
