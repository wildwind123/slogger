package slogger

import (
	"context"
	"log/slog"
)

type ctxKey string

const loggerKey = ctxKey("logger")
const trackIDKey = ctxKey("track_id")

func ToCtx(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}

func FromCtx(ctx context.Context) *slog.Logger {
	logger, ok := ctx.Value(loggerKey).(*slog.Logger)
	if !ok {
		slog.Error("context does not have logger.  will return slog default logger")
		return slog.Default()
	}

	return logger
}

func TrackIDToCtx(ctx context.Context, trackID string) context.Context {
	return context.WithValue(ctx, trackIDKey, trackID)
}

func TrackIDFromCtx(ctx context.Context) string {
	trackID, ok := ctx.Value(trackIDKey).(string)
	if !ok {
		return ""
	}
	return trackID
}
