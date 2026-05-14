package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"soullocke-backend/config"
	"soullocke-backend/db"
	"soullocke-backend/domain/game_edition"
	"soullocke-backend/domain/language"
	"soullocke-backend/domain/lobby"
	"soullocke-backend/domain/token"
	"soullocke-backend/http"
	"strings"
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

	if len(strings.TrimSpace(appConfig.Localization.DefaultLocale)) == 0 {
		slog.Error("Default locale not specified")
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

	langR := language.NewRepository(database)
	defaultLanguage, err := langR.GetLanguageByName(ctx, appConfig.Localization.DefaultLocale)
	if err != nil {
		slog.Error("Failed to get default language", "error", err)
		return
	}

	var supportedLocales []language.Language
	for _, locale := range appConfig.Localization.SupportedLocales {
		if locale == appConfig.Localization.DefaultLocale {
			continue
		}
		lang, err := langR.GetLanguageByName(ctx, locale)
		if err != nil {
			slog.Warn("Failed to get supported language", "error", err, "locale", locale)
			continue
		}
		supportedLocales = append(supportedLocales, *lang)
	}

	allLang := make([]language.Language, 0, len(supportedLocales)+1)
	allLang = append(allLang, supportedLocales...)
	allLang = append(allLang, *defaultLanguage)

	lr := lobby.NewRepository(database)
	ger := game_edition.NewRepository(database)
	tr := token.NewRepository(database)
	server := http.NewServer(appConfig.Server.Port, appConfig.Server.AllowedOrigins, lr, ger, tr, allLang, langR)
	server.Setup()

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
