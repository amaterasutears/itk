package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/amaterasutears/itk/config"
	sqlc "github.com/amaterasutears/itk/internal/client/sql"
	transaction_handler "github.com/amaterasutears/itk/internal/http/handler/transaction"
	wallet_handler "github.com/amaterasutears/itk/internal/http/handler/wallet"
	"github.com/amaterasutears/itk/internal/http/router"
	"github.com/amaterasutears/itk/internal/http/server"
	transaction_service "github.com/amaterasutears/itk/internal/service/transaction"
	wallet_service "github.com/amaterasutears/itk/internal/service/wallet"
	"github.com/amaterasutears/itk/internal/storage/migrator"
	transaction_storage "github.com/amaterasutears/itk/internal/storage/transaction"
	"github.com/amaterasutears/itk/internal/storage/transactor"
	wallet_storage "github.com/amaterasutears/itk/internal/storage/wallet"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	c, err := config.Load()
	if err != nil {
		panic(err)
	}

	pgc, err := sqlc.New(
		c.DataSourceName(),
		sqlc.WithMaxOpenConns(c.Postgres.MaxOpenConns),
		sqlc.WithMaxIdleConns(c.Postgres.MaxIdleConns),
		sqlc.WithConnMaxLifetime(time.Duration(c.Postgres.ConnMaxLifetimeMin)*time.Minute),
		sqlc.WithConnMaxIdleTime(time.Duration(c.Postgres.ConnMaxIdleTimeMin)*time.Minute),
	)
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

	tactor := transactor.New(pgc.DB())

	wstorage := wallet_storage.New(pgc.DB())
	wservice := wallet_service.New(wstorage)
	whandler := wallet_handler.New(wservice)

	tstorage := transaction_storage.New(pgc.DB())
	tservice := transaction_service.New(wstorage, tstorage, tactor)
	thandler := transaction_handler.New(tservice)

	r := router.New(whandler, thandler)

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
