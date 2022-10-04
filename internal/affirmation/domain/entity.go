package domain

import (
	"time"

	"github.com/oklog/ulid/v2"
)

type Affirmation struct {
	Id         ulid.ULID `json:"id"`
	CategoryId ulid.ULID `json:"categoryId"`
	Text       string    `json:"text"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
