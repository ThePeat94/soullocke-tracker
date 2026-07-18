package token

import (
	"context"
	"fmt"
	"log/slog"
	"time"
)

func LoopCleanup(ctx context.Context, tr Repository, interval time.Duration) {
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				if count, err := tr.DeleteExpiredTokens(ctx); err != nil {
					slog.Error("token cleanup failed", "error", err)
				} else {
					slog.Info(fmt.Sprintf("deleted %d expired tokens", count))
				}
			case <-ctx.Done():
				return
			}
		}
	}()
}
