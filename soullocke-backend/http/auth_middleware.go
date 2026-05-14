package http

import (
	"log/slog"
	"net/http"
	"soullocke-backend/domain/token"

	"github.com/danielgtaylor/huma/v2"
)

type contextKey string

const lobbyAccessContextKey contextKey = "lobbyAccess"

func RequireLobbyTokenAuthentication(api huma.API, tr token.TokenRepository) func(ctx huma.Context, next func(ctx2 huma.Context)) {
	return func(ctx huma.Context, next func(huma.Context)) {
		lobbyID := ctx.Param("lobbyId")
		if lobbyID == "" {
			_ = huma.WriteErr(api, ctx, http.StatusBadRequest, "lobby id required")
			return
		}

		cookie, err := huma.ReadCookie(ctx, lobbyAccessCookieName)
		if err != nil {
			slog.Error("error reading cookie", "error", err)
			_ = huma.WriteErr(api, ctx, http.StatusUnauthorized, "unauthorized")
			return
		}

		tokenHash := hashAccessToken(cookie.Value)

		foundToken, err := tr.GetTokenByHash(ctx.Context(), tokenHash)
		if err != nil {
			_ = huma.WriteErr(api, ctx, http.StatusUnauthorized, "unauthorized")
			return
		}

		if foundToken.LobbyId != lobbyID {
			_ = huma.WriteErr(api, ctx, http.StatusUnauthorized, "unauthorized")
			return
		}

		ctx = huma.WithValue(ctx, lobbyAccessContextKey, foundToken)

		next(ctx)
	}
}
