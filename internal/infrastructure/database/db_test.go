package database_test

import (
	"testing"

	"github.com/Gadoma/RandomAffirmationsApi/internal/infrastructure/database"
	"github.com/Gadoma/RandomAffirmationsApi/test"
)

func TestMain(m *testing.M) {
	test.PrepareTestDB()
	m.Run()
	test.CleanupTestDB()
}

func TestDB(t *testing.T) {
	db := MustOpenDB(t)
	MustCloseDB(t, db)
}

func MustOpenDB(t *testing.T) *database.DB {
	db := database.NewDB(test.GetDSN(test.TestDbDSN))
	if err := db.Open(); err != nil {
		t.Fatal(err)
	}
	return db
}

func MustCloseDB(t *testing.T, db *database.DB) {
	if err := db.Close(); err != nil {
		t.Fatal(err)
	}
}
