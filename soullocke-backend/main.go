package main

import (
	"context"
	"fmt"
	"log/slog"
	db2 "soullocke-backend/db"
)

func main() {
	mainCtx, _ := context.WithCancel(context.Background())
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		"jailer", "sylv4n45", "localhost", 5445, "soullocker",
	)
	db, err := db2.NewDatabase(mainCtx, dsn)
	if err != nil {
		slog.Error("Failed to boot up database", err)
		return
	}
	slog.Info("successfully connected to postgres")

	err = db.Migrate(mainCtx)
	if err != nil {
		slog.Error("Failed to migrate database", err)
		return
	}
	slog.Info("successfully migrated database")
}
