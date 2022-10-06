module github.com/gadoma/rafapi/internal/affirmation

go 1.19

exclude github.com/gadoma/rafapi v0.0.0-20221005180506-ac49759c4e07

require (
	github.com/gadoma/rafapi/internal/common v0.0.0-00010101000000-000000000000
	github.com/gorilla/mux v1.8.0
	github.com/oklog/ulid/v2 v2.1.0
)

require github.com/mattn/go-sqlite3 v1.14.15 // indirect

replace github.com/gadoma/rafapi/internal/common => ../common
