package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	app "github.com/gadoma/rafapi/internal/common/infrastructure"
	commonHttp "github.com/gadoma/rafapi/internal/common/infrastructure/http"
)

const (
	defaultDSN          = "db/db.dist.sqlite"
	defaultServerAddr   = "0.0.0.0:5000"
	defaultServerDomain = "localhost"
)

func getRuntimeConfig() (dbDsn, serverAddr, serverDomain string) {
	dbDsn = os.Getenv("RAFAPI_DB_DSN")
	if dbDsn == "" {
		dbDsn = defaultDSN
	}

	serverAddr = os.Getenv("RAFAPI_SERVER_ADDR")
	if serverAddr == "" {
		serverAddr = defaultServerAddr
	}

	serverDomain = os.Getenv("RAFAPI_SERVER_DOMAIN")
	if serverDomain == "" {
		serverDomain = defaultServerDomain
	}
	return
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() { <-c; cancel() }()

	dbDsn, serverAddr, serverDomain := getRuntimeConfig()

	bootstrap := commonHttp.NewBootstrap()

	main := app.NewApp(&app.AppConfig{
		DbDSN:        dbDsn,
		ServerAddr:   serverAddr,
		ServerDomain: serverDomain,
	}, bootstrap)

	if runErr := main.Run(ctx); runErr != nil {
		fmt.Fprintln(os.Stderr, runErr)

		if haltErr := main.Halt(); haltErr != nil {
			fmt.Fprintln(os.Stderr, haltErr)
		}

		os.Exit(1)
	}

	<-ctx.Done()

	if haltErr := main.Halt(); haltErr != nil {
		fmt.Fprintln(os.Stderr, haltErr)

		if haltRetryErr := main.Halt(); haltRetryErr != nil {
			fmt.Fprintln(os.Stderr, haltRetryErr)
		}

		os.Exit(1)
	}
}
