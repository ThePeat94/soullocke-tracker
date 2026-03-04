package main

import (
	"context"
	"fmt"
	"log/slog"
	db2 "soullocke-backend/db"
)

func main() {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		"jailer", "sylv4n45", "localhost", 5445, "soullocker",
	)
	_, err := db2.NewDatabase(context.Background(), dsn)
	if err != nil {
		panic(err)
	}
	slog.Info("successfully connected to postgres")
}
