package logger

import (
	"log/slog"
	"os"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/emgeorrk/sinbitus/internal/config"
	"github.com/emgeorrk/sinbitus/internal/constants"
	"github.com/muesli/termenv"
	"github.com/samber/slog-multi"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

var Module = fx.Options(
	fx.Provide(NewLogger),
	fx.WithLogger(NewFxLogger),
)

type Logger struct {
	*slog.Logger
}

func (l Logger) Write(p []byte) (n int, err error) {
	l.Info(string(p))
	return len(p), nil
}

func NewLogger(config *config.Config) *Logger {
	level := log.InfoLevel
	switch strings.ToLower(config.Log.Level) {
	case "debug":
		level = log.DebugLevel
	case "info":
		level = log.InfoLevel
	case "warn":
		level = log.WarnLevel
	case "error":
		level = log.ErrorLevel
	}

	formatter := log.TextFormatter
	switch strings.ToLower(config.Log.Formatter) {
	case "json":
		formatter = log.JSONFormatter
	case "text":
		formatter = log.TextFormatter
	case "fmt":
		formatter = log.LogfmtFormatter
	}

	sourceFormat := log.ShortCallerFormatter
	if strings.ToLower(config.Log.SourceFormat) == "long" {
		sourceFormat = log.LongCallerFormatter
	}

	file, err := os.OpenFile(constants.LogOutputFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	terminalOpts := log.Options{
		ReportTimestamp: config.Log.Timestamp != "",
		TimeFormat:      config.Log.TSFormat,
		Level:           level,
		Prefix:          config.Log.Prefix,
		ReportCaller:    strings.ToLower(config.Log.Source) == "yes",
		CallerFormatter: sourceFormat,
		Formatter:       formatter,
	}

	fileOpts := log.Options{
		ReportTimestamp: config.Log.Timestamp != "",
		TimeFormat:      config.Log.TSFormat,
		Level:           level,
		Prefix:          config.Log.Prefix,
		ReportCaller:    strings.ToLower(config.Log.Source) == "yes",
		CallerFormatter: sourceFormat,
		Formatter:       log.LogfmtFormatter,
	}

	terminalHandler := log.NewWithOptions(os.Stdout, terminalOpts)
	fileHandler := log.NewWithOptions(file, fileOpts)

	if strings.ToLower(config.Log.Color) == "yes" {
		terminalHandler.SetColorProfile(termenv.TrueColor)
	}

	if config.Log.Label != "" {
		terminalHandler = terminalHandler.With("label", config.Log.Label)
		fileHandler = fileHandler.With("label", config.Log.Label)
	}

	logger := slog.New(slogmulti.Fanout(terminalHandler, fileHandler))

	logger.Info("Logger initialized", "level", config.Log.Level)

	return &Logger{
		Logger: logger,
	}
}

func NewFxLogger(l *Logger, config *config.Config) fxevent.Logger {
	if strings.ToLower(config.Log.EnableFxLogs) == "yes" {
		return &fxevent.SlogLogger{
			Logger: l.Logger.With("source", "fx"),
		}
	}

	return &fxevent.NopLogger
}

func NewTestLogger() Logger {
	handler := log.NewWithOptions(os.Stderr, log.Options{
		Level:        log.DebugLevel,
		Formatter:    log.TextFormatter,
		ReportCaller: true,
	})

	logger := slog.New(handler)

	return Logger{
		Logger: logger,
	}
}

func (l Logger) With(key, value string) Logger {
	l.Logger = l.Logger.With(key, value)

	return l
}
