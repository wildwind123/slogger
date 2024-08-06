package slogger

import (
	"log/slog"
	"os"
	"testing"
)

func TestLogger(t *testing.T) {
	// t.Skip()
	logger := NewLogger(&Options{
		Level:     slog.LevelDebug,
		AddSource: true,
		Writer:    os.Stdout,
		App:       "test_app",
		Build:     "v1.2",
	})
	logger.Info("test", slog.String("test", "test_attr"))
}
