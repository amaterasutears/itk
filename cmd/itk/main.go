package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/amaterasutears/itk/config"
	sqlc "github.com/amaterasutears/itk/internal/client/sql"
	wallet_handler "github.com/amaterasutears/itk/internal/http/handler/wallet"
	"github.com/amaterasutears/itk/internal/http/router"
	"github.com/amaterasutears/itk/internal/http/server"
	wallet_service "github.com/amaterasutears/itk/internal/service/wallet"
	"github.com/amaterasutears/itk/internal/storage/migrator"
	wallet_storage "github.com/amaterasutears/itk/internal/storage/wallet"
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

	mr, err := migrator.New(pgc.DB())
	if err != nil {
		panic(err)
	}

	mctx, mcancel := context.WithTimeout(ctx, time.Duration(c.Migrator.TimeoutSec)*time.Second)
	defer mcancel()

	err = mr.Up(mctx)
	if err != nil {
		panic(err)
	}

	wstorage := wallet_storage.New(pgc.DB())
	wservice := wallet_service.New(wstorage)
	whandler := wallet_handler.New(wservice)

	r := router.New(whandler)

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
