package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"

	"github.com/shiroemons/touhou_arrangement_chronicle/go/graph/resolver"
	"github.com/shiroemons/touhou_arrangement_chronicle/go/pkg/config"
	"github.com/shiroemons/touhou_arrangement_chronicle/go/pkg/loader"
	"github.com/shiroemons/touhou_arrangement_chronicle/go/pkg/repository"
	"github.com/shiroemons/touhou_arrangement_chronicle/go/pkg/server"
	"github.com/shiroemons/touhou_arrangement_chronicle/go/pkg/service"
	"github.com/shiroemons/touhou_arrangement_chronicle/go/pkg/store"
)

func inject() fx.Option {
	ctx := func() context.Context {
		return context.Background()
	}

	return fx.Options(
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
		fx.Provide(
			// Context
			fx.Annotate(ctx, fx.As(new(context.Context))),
			// Config
			config.Provider,
			// Logger
			zap.NewExample,
		),
		// Provide
		store.Module,
		repository.Module,
		service.Module,
		loader.InitModule,
		loader.Module,
		resolver.Module,
		server.Module,

		// Invoke
		server.InvokeModule,
	)
}

func main() {
	app := fx.New(inject())

	startCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := app.Start(startCtx); err != nil {
		log.Fatal(err)
	}

	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	stopCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := app.Stop(stopCtx); err != nil {
		log.Fatal(err)
	}
}
