package test

import (
	"io"
	"os"
	"path/filepath"
	"runtime"
)

const (
	FixtureDbDSN = "db_test.dist.sqlite"
	TestDbDSN    = "db_test.sqlite"
)

func PrepareTestDB() error {
	in, err := os.Open(GetDSN(FixtureDbDSN))
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
	os.Remove(GetDSN(TestDbDSN) + "-wal")
	os.Remove(GetDSN(TestDbDSN) + "-shm")
	return os.Remove(GetDSN(TestDbDSN))
}

func GetDSN(dsn string) string {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	return basepath + "/../db/" + dsn
}
