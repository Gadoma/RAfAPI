package main_test

import (
	"context"
	"testing"

	app "github.com/gadoma/rafapi/internal/common/infrastructure"
	commonHttp "github.com/gadoma/rafapi/internal/common/infrastructure/http"
	"github.com/gadoma/rafapi/test"
)

const (
	testServerAddr   = "0.0.0.0:5001"
	testServerDomain = "localhost"
	statusOk         = "OK"
	statusError      = "ERROR"
)

func MustRunMain(t *testing.T) *app.App {
	test.PrepareTestDB()
	b := commonHttp.NewBootstrap()
	main := app.NewApp(&app.AppConfig{
		DbDSN:        test.GetDSN(test.TestDbDSN),
		ServerAddr:   testServerAddr,
		ServerDomain: testServerDomain,
	}, b)

	if err := main.Run(context.Background()); err != nil {
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
