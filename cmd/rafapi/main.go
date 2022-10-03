package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	applicationAff "github.com/gadoma/rafapi/internal/affirmation/application"
	applicationCat "github.com/gadoma/rafapi/internal/category/application"
	appicationRaf "github.com/gadoma/rafapi/internal/randomAffirmation/application"

	httpAff "github.com/gadoma/rafapi/internal/affirmation/infrastructure/http"
	httpCat "github.com/gadoma/rafapi/internal/category/infrastructure/http"
	httpRaf "github.com/gadoma/rafapi/internal/randomAffirmation/infrastructure/http"

	dbAff "github.com/gadoma/rafapi/internal/affirmation/infrastructure/database"
	dbCat "github.com/gadoma/rafapi/internal/category/infrastructure/database"
	dbRaf "github.com/gadoma/rafapi/internal/randomAffirmation/infrastructure/database"

	commonDb "github.com/gadoma/rafapi/internal/common/infrastructure/database"
	commonHttp "github.com/gadoma/rafapi/internal/common/infrastructure/http"
)

type AppConfig struct {
	DbDSN        string
	ServerAddr   string
	ServerDomain string
}

const (
	defaultDSN          = "db/db.dist.sqlite"
	defaultServerAddr   = "0.0.0.0:5000"
	defaultServerDomain = "localhost"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() { <-c; cancel() }()

	m := NewApp(&AppConfig{
		DbDSN:        defaultDSN,
		ServerAddr:   defaultServerAddr,
		ServerDomain: defaultServerDomain,
	})

	if err := m.Run(ctx); err != nil {
		m.Close()
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	<-ctx.Done()

	if err := m.Close(); err != nil {
		m.Close()
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

type App struct {
	DB *commonDb.DB

	HTTPServer *commonHttp.Server

	Config *AppConfig
}

func NewApp(config *AppConfig) *App {
	return &App{
		DB:         commonDb.NewDB(config.DbDSN),
		HTTPServer: commonHttp.NewServer(),
		Config:     config,
	}
}

func (m *App) Close() error {
	if m.HTTPServer != nil {
		if err := m.HTTPServer.Close(); err != nil {
			return err
		}
	}

	if m.DB != nil {
		if err := m.DB.Close(); err != nil {
			return err
		}
	}

	return nil
}

func (app *App) Run(ctx context.Context) (err error) {

	app.DB.DSN = app.Config.DbDSN

	if err := app.DB.Open(); err != nil {
		return fmt.Errorf("cannot open db %s because of %w", app.Config.DbDSN, err)
	}

	app.HTTPServer.Addr = app.Config.ServerAddr
	app.HTTPServer.Domain = app.Config.ServerDomain

	responder := commonHttp.NewResponder()

	affService := applicationAff.NewAffirmationService(dbAff.NewAffirmationRepository(app.DB))
	catService := applicationCat.NewCategoryService(dbCat.NewCategoryRepository(app.DB))
	rafService := appicationRaf.NewRandomAffirmationService(dbRaf.NewRandomAffirmationRepository(app.DB))

	affReqHandler := httpAff.NewAffirmationRequestHandler()
	catReqHandler := httpCat.NewCategoryRequestHandler()
	rafReqHandler := httpRaf.NewRandomAffirmationRequestHandler()

	affirmationController := httpAff.NewAffirmationController(affService, responder, affReqHandler)
	categoryController := httpCat.NewCategoryController(catService, responder, catReqHandler)
	randomAffirmationController := httpRaf.NewRandomAffirmationController(rafService, responder, rafReqHandler)

	app.HTTPServer.AffirmationController = affirmationController
	app.HTTPServer.CategoryController = categoryController
	app.HTTPServer.RandomAffirmationController = randomAffirmationController

	app.HTTPServer.RegisterRoutes()

	if err := app.HTTPServer.Open(); err != nil {
		return fmt.Errorf("cannot open http server on %s because of %w", app.Config.ServerAddr, err)
	}

	return nil
}
