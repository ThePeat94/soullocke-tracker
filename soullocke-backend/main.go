package main

import (
	"context"
	"log/slog"
	"soullocke-backend/config"
	"soullocke-backend/db"
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

	err = database.Migrate(mainCtx)
	if err != nil {
		slog.Error("Failed to migrate database", "error", err)
		return
	}
	slog.Info("successfully migrated database")
}
