package main_test

import (
	"testing"

	"github.com/gadoma/rafapi/internal/affirmation/infrastructure/http"
	app "github.com/gadoma/rafapi/internal/common/infrastructure"
	"github.com/gadoma/rafapi/internal/common/test"
)

const (
	testServerAddr   = "0.0.0.0:5001"
	testServerDomain = "localhost"
	statusOk         = "OK"
	statusError      = "ERROR"
)

func MustRunMain(t *testing.T) *app.App {
	test.PrepareTestDB()
	b := http.NewBootstrap()
	main := app.NewApp(&app.AppConfig{
		DbDSN:        test.GetDSN(test.TestingDbDSN),
		ServerAddr:   testServerAddr,
		ServerDomain: testServerDomain,
	}, b)

	if err := main.Run(); err != nil {
		t.Fatal(err)
	}

	return main
}

func MustCloseMain(t *testing.T, main *app.App) {
	defer test.CleanupTestDB()
	if err := main.Halt(); err != nil {
		t.Fatal(err)
	}
}
