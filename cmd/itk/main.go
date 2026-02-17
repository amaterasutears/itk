package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/amaterasutears/itk/config"
	sqlc "github.com/amaterasutears/itk/internal/client/sql"
	"github.com/amaterasutears/itk/internal/http/router"
	"github.com/amaterasutears/itk/internal/http/server"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	c, err := config.Load()
	if err != nil {
		panic(err)
	}

	pgc, err := sqlc.New(c.DataSourceName())
	if err != nil {
		panic(err)
	}

	pctx, pcancel := context.WithTimeout(ctx, time.Duration(c.Postgres.PingTimeout)*time.Second)
	defer pcancel()

	err = pgc.Ping(pctx)
	if err != nil {
		panic(err)
	}

	r := router.New()

	srv := server.New(r, &c.Server)

	go func() {
		err = srv.Start()
		if err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	<-ctx.Done()

	sctx, scancel := context.WithTimeout(context.Background(), time.Duration(c.Server.ShutdownTimeoutSec)*time.Second)
	defer scancel()

	err = srv.Shutdown(sctx)
	if err != nil {
		panic(err)
	}

	err = pgc.Close()
	if err != nil {
		panic(err)
	}
}
