package main

import (
	"context"
	"log/slog"
	"soullocke-backend/config"
	"soullocke-backend/db"
	"soullocke-backend/domain/lobby"
	"soullocke-backend/http"
)

func main() {
	mainCtx := context.Background()

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
	server.Start()

	select {}
}
