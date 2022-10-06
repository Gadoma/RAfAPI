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
	TestingDbDSN = "db_test.sqlite"
)

func MustOpenDB(t *testing.T) *database.DB {
	db := database.NewDB(GetDSN(TestingDbDSN))
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

func PrepareTestDB() {
	in, _ := os.Open(GetDSN(fixtureDbDSN))
	defer func(in *os.File) {
		_ = in.Close()
	}(in)

	out, _ := os.Create(GetDSN(TestingDbDSN))
	defer func(out *os.File) {
		_ = out.Close()
	}(out)

	_, _ = io.Copy(out, in)
	_ = out.Close()
}

func CleanupTestDB() {
	_ = os.Remove(GetDSN(TestingDbDSN))
}

func GetDSN(dsn string) string {
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(b)

	return basePath + "/../../../db/" + dsn
}
