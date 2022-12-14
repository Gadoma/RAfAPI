package infrastructure

import (
	"fmt"

	commonDb "github.com/gadoma/rafapi/internal/common/infrastructure/database"
	commonHttp "github.com/gadoma/rafapi/internal/common/infrastructure/http"
)

type AppConfig struct {
	DbDSN        string
	ServerAddr   string
	ServerDomain string
}

type App struct {
	DB         *commonDb.DB
	HTTPServer *commonHttp.Server
	Config     *AppConfig
}

func NewApp(config *AppConfig, bootstrap commonHttp.Bootstrapper) *App {
	db := commonDb.NewDB(config.DbDSN)
	controllers := bootstrap.BootstrapControllers(db)
	server := commonHttp.NewServer(controllers)

	return &App{
		DB:         db,
		HTTPServer: server,
		Config:     config,
	}
}

func (app *App) Run() error {
	app.DB.DSN = app.Config.DbDSN

	if err := app.DB.Open(); err != nil {
		return fmt.Errorf("cannot open db %s because of %w", app.Config.DbDSN, err)
	}

	app.HTTPServer.Addr = app.Config.ServerAddr
	app.HTTPServer.Domain = app.Config.ServerDomain
	app.HTTPServer.RegisterRoutes()

	if err := app.HTTPServer.Open(); err != nil {
		return fmt.Errorf("cannot open http server on %s because of %w", app.Config.ServerAddr, err)
	}

	return nil
}

func (app *App) Halt() error {
	if app.HTTPServer != nil {
		if err := app.HTTPServer.Close(); err != nil {
			return err
		}
	}

	if app.DB != nil {
		if err := app.DB.Close(); err != nil {
			return err
		}
	}

	return nil
}
