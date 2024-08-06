package slogger

import (
	"context"
	"log/slog"
	"testing"
)

func TestCtxLogger(t *testing.T) {
	l := slog.Default()
	ctx := context.Background()
	ctxWithLogger := ToCtx(ctx, l)
	logger := FromCtx(ctxWithLogger)

	logger.Info("test logger")

}

func TestTrackIDCtx(t *testing.T) {
	ctx := context.Background()

	sampleTrackID := "sample_track_id"

	ctx = TrackIDToCtx(ctx, sampleTrackID)

	trackID := TrackIDFromCtx(ctx)

	if trackID != sampleTrackID {
		t.Error("wrong request id")
	}

	l := slog.Default()

	l = l.With(AttrTrackID(ctx))

	ctx = ToCtx(ctx, l)

	ctxLogger := FromCtx(ctx)

	ctxLogger.Info("test logger")
}
