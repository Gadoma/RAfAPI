package main_test

import (
	"context"
	"testing"

	main "github.com/gadoma/rafapi/cmd/rafapi"
	"github.com/gadoma/rafapi/test"
)

const (
	TestServerAddr   = "0.0.0.0:5001"
	TestServerDomain = "localhost"
)

func MustRunMain(t *testing.T) *main.App {
	test.PrepareTestDB()
	m := main.NewApp(&main.AppConfig{
		DbDSN:        test.GetDSN(test.TestDbDSN),
		ServerAddr:   TestServerAddr,
		ServerDomain: TestServerDomain,
	})

	if err := m.Run(context.Background()); err != nil {
		t.Fatal(err)
	}

	return m
}

func MustCloseMain(t *testing.T, m *main.App) {
	defer test.CleanupTestDB()
	if err := m.Close(); err != nil {
		t.Fatal(err)
	}
}
