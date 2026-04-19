package globalctx

import "log/slog"

// GlobalContext holds app-wide dependencies (logger, DB, etc.).
type GlobalContext struct {
	Logger *slog.Logger
}
