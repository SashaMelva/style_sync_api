package cmd

import (
	"context"
	"os/signal"
	"syscall"
	"time"

	"github.com/SashaMelva/style_sync_api/internal/app"
	"github.com/SashaMelva/style_sync_api/internal/config"
	"github.com/SashaMelva/style_sync_api/internal/logger"
	"github.com/SashaMelva/style_sync_api/internal/storage"
	"github.com/SashaMelva/style_sync_api/server/http"
)

func main() {
	configFile := "../config/"
	config := config.New(configFile)
	log := logger.New(config.Logger, "../log/")

	storage := storage.New(config.DataBase, log)
	// err := migrator.RunMigrationsPg(connectionDB, "migrations")

	// if err != nil {
	// 	log.Fatal(err)
	// }

	app := app.New(log, storage)

	httpServer := http.NewServer(log, app, config.HttpServer)

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		httpServer.Stop(ctx)
	}()

	log.Info("Services is running...")
	log.Debug("Debug mode enabled")

	httpServer.Start(ctx)
}
