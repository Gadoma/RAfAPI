package database_test

import (
	"testing"

	"github.com/gadoma/rafapi/internal/common/test"
)

func TestMain(m *testing.M) {
	test.PrepareTestDB()
	m.Run()
	test.CleanupTestDB()
}

func TestDB(t *testing.T) {
	db := test.MustOpenDB(t)
	test.MustCloseDB(t, db)
}
