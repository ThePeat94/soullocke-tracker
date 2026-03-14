package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"soullocke-backend/config"
	"soullocke-backend/db"
	"soullocke-backend/domain/lobby"
	"soullocke-backend/http"
	"syscall"

	"golang.org/x/sync/errgroup"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	grp, gCtx := errgroup.WithContext(ctx)

	appConfig, err := config.LoadConfig("config.yml")
	if err != nil {
		slog.Error("Failed to load config", "error", err)
		return
	}

	slog.Info("successfully loaded config")
	database, err := db.NewDatabase(gCtx, appConfig.Database.DSN())
	if err != nil {
		slog.Error("Failed to boot up database", "error", err)
		return
	}
	defer database.Close()
	slog.Info("successfully connected to postgres")

	if len(os.Args) <= 1 || (len(os.Args) > 1 && os.Args[1] != "export-openapi") {
		err = database.Migrate()
		if err != nil {
			slog.Error("Failed to migrate database", "error", err)
			return
		}
		slog.Info("successfully migrated database")
	}

	lr := lobby.NewRepository(*database)
	server := http.NewServer(appConfig.Server.Port, lr)
	err = server.Setup()
	if err != nil {
		slog.Error("Failed to start server", "error", err)
		return
	}

	if len(os.Args) > 1 && os.Args[1] == "export-openapi" {
		path := "../openapi/openapi.yaml"
		if len(os.Args) > 2 {
			path = os.Args[2]
		}
		if err := server.ExportOpenAPISpec(path); err != nil {
			slog.Error("Failed to export openapi spec", "error", err)
			os.Exit(1)
		}
		slog.Info("exported openapi spec", "path", path)
		return
	}

	grp.Go(func() error {
		return server.Serve(gCtx)
	})

	err = grp.Wait()
	if err != nil {
		slog.Error("main err", "error", err)
	}
}
