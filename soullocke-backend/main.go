package main

import (
	"context"
	"fmt"
	"log/slog"
	"soullocke-backend/config"
	"soullocke-backend/db"
)

func main() {
	mainCtx, _ := context.WithCancel(context.Background())

	appConfig, err := config.LoadConfig("config.yml")
	if err != nil {
		slog.Error("Failed to load config", err)
	}

	slog.Info("successfully loaded config")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		appConfig.Database.User,
		appConfig.Database.Password,
		appConfig.Database.Host,
		appConfig.Database.Port,
		appConfig.Database.Database,
	)
	database, err := db.NewDatabase(mainCtx, dsn)
	if err != nil {
		slog.Error("Failed to boot up database", err)
		return
	}
	slog.Info("successfully connected to postgres")

	err = database.Migrate(mainCtx)
	if err != nil {
		slog.Error("Failed to migrate database", err)
		return
	}
	slog.Info("successfully migrated database")
}
