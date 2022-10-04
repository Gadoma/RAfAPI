package test

import (
	"io"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/gadoma/rafapi/internal/common/infrastructure/database"
)

const (
	fixtureDbDSN = "db_test.dist.sqlite"
	TestDbDSN    = "db_test.sqlite"
)

func MustOpenDB(t *testing.T) *database.DB {
	db := database.NewDB(GetDSN(TestDbDSN))
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

func PrepareTestDB() error {
	in, err := os.Open(GetDSN(fixtureDbDSN))
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(GetDSN(TestDbDSN))
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}

func CleanupTestDB() error {
	return os.Remove(GetDSN(TestDbDSN))
}

func GetDSN(dsn string) string {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	return basepath + "/../db/" + dsn
}
