package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/gadoma/rafapi/internal/application"
	"github.com/gadoma/rafapi/internal/infrastructure/database"
	"github.com/gadoma/rafapi/internal/infrastructure/http"
)

type AppConfig struct {
	DbDSN        string
	ServerAddr   string
	ServerDomain string
}

const (
	DefaultDSN          = "db/db.dist.sqlite"
	DefaultServerAddr   = "0.0.0.0:5000"
	DefaultServerDomain = "localhost"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() { <-c; cancel() }()

	m := NewApp(&AppConfig{
		DbDSN:        DefaultDSN,
		ServerAddr:   DefaultServerAddr,
		ServerDomain: DefaultServerDomain,
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
	DB *database.DB

	HTTPServer *http.Server

	Config *AppConfig
}

func NewApp(config *AppConfig) *App {
	return &App{
		DB:         database.NewDB(config.DbDSN),
		HTTPServer: http.NewServer(),
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

	affirmationService := application.NewAffirmationService(database.NewAffirmationRepository(app.DB))
	categoryService := application.NewCategoryService(database.NewCategoryRepository(app.DB))
	randomAffirmationService := application.NewRandomAffirmationService(database.NewRandomAffirmationRepository(app.DB))

	app.HTTPServer.Addr = app.Config.ServerAddr
	app.HTTPServer.Domain = app.Config.ServerDomain

	app.HTTPServer.AffirmationService = affirmationService
	app.HTTPServer.CategoryService = categoryService
	app.HTTPServer.RandomAffirmationService = randomAffirmationService

	if err := app.HTTPServer.Open(); err != nil {
		return fmt.Errorf("cannot open http server on %s because of %w", app.Config.ServerAddr, err)
	}

	return nil
}
