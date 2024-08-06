package slogger

import (
	"context"
	"io"
	"log/slog"
	"os"
	"time"

	slogformatter "github.com/samber/slog-formatter"
)

const defaultTimeFormat = "2006-01-02 15:04:05.999999999"

type Options struct {
	Level     slog.Level
	AddSource bool
	Writer    io.Writer
	App       string
	Build     string
}

func NewLogger(options *Options) *slog.Logger {

	var writer io.Writer = options.Writer
	if writer == nil {
		writer = os.Stdout
	}

	logger := slog.New(
		slogformatter.NewFormatterHandler(
			slogformatter.TimeFormatter(defaultTimeFormat, time.UTC),
			slogformatter.ErrorFormatter("err"),
		)(

			slog.NewJSONHandler(options.Writer, &slog.HandlerOptions{
				ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
					if a.Key == "time" {
						a = slog.String("time", a.Value.Time().UTC().Format(defaultTimeFormat))
					}
					return a
				},
				AddSource: options.AddSource,
				Level:     options.Level,
			}),
		),
	)
	logger = logger.With(
		slog.String("app", options.App),
		slog.String("build", options.Build),
	).WithGroup("attrs")

	return logger
}

func AttrTrackID(ctx context.Context) slog.Attr {
	trackID := TrackIDFromCtx(ctx)
	if trackID == "" {
		// return empty attribute
		return slog.Attr{}
	}

	return slog.String("track_id", trackID)
}
