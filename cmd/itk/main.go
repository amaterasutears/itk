package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/amaterasutears/itk/config"
	sqlc "github.com/amaterasutears/itk/internal/client/sql"
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

	pctx, pcancel := context.WithTimeout(ctx, 5*time.Second)
	defer pcancel()

	err = pgc.Ping(pctx)
	if err != nil {
		panic(err)
	}

	<-ctx.Done()

	err = pgc.Close()
	if err != nil {
		panic(err)
	}
}
