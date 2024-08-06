package slogger

import (
	"log/slog"
	"net/http"
	"testing"
)

func TestWrite(t *testing.T) {
	t.Skip()
	v := &Vector{
		Client:   &http.Client{},
		Url:      "http://localhost",
		User:     "vector",
		Password: "vector_password",
	}

	logger := NewLogger(&Options{
		Level:     slog.LevelDebug,
		AddSource: true,
		Writer:    v,
		App:       "app",
		Build:     "v1.1",
	})
	logger.Info("test logger", slog.String("test_attr", "test_attr_value"))
}
