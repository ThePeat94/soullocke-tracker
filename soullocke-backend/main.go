package main

import (
	"context"
	"log/slog"
	"soullocke-backend/config"
	"soullocke-backend/db"
	"soullocke-backend/domain/lobby"
	"soullocke-backend/http"

	"golang.org/x/sync/errgroup"
)

func main() {
	mainCtx := context.Background()
	groupCtx, mainCtx := errgroup.WithContext(mainCtx)

	appConfig, err := config.LoadConfig("config.yml")
	if err != nil {
		slog.Error("Failed to load config", "error", err)
		return
	}

	slog.Info("successfully loaded config")
	database, err := db.NewDatabase(mainCtx, appConfig.Database.DSN())
	if err != nil {
		slog.Error("Failed to boot up database", "error", err)
		return
	}
	defer database.Close()
	slog.Info("successfully connected to postgres")

	err = database.Migrate()
	if err != nil {
		slog.Error("Failed to migrate database", "error", err)
		return
	}
	slog.Info("successfully migrated database")

	lr := lobby.NewRepository(*database)
	server := http.NewServer(1337, lr)
	err = server.Setup()
	if err != nil {
		slog.Error("Failed to start server", "error", err)
		return
	}

	err = server.ExportOpenAPISpec("../openapi.yaml")
	if err != nil {
		slog.Warn("Failed to export openapi.yaml", "error", err)
	} else {
		slog.Info("successfully exported openapi.yaml")
	}

	groupCtx.Go(func() error {
		return server.Serve()
	})

	select {}
}
