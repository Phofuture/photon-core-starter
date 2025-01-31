package log

import (
	"context"
	"log/slog"
)

type DefaultLogger struct {
	Logger *slog.Logger
}

// TODO: 處理slog設定
var defaultLogger = &DefaultLogger{Logger: slog.Default()}

func GetDefaultLogger() *DefaultLogger {
	return defaultLogger
}

func (l *DefaultLogger) DebugContext(msg string, args ...any) {
	l.Logger.Debug(msg, args...)
}

func (l *DefaultLogger) InfoContext(msg string, args ...any) {
	l.Logger.Info(msg, args...)
}

func (l *DefaultLogger) WarnContext(msg string, args ...any) {
	l.Logger.Warn(msg, args...)
}

func (l *DefaultLogger) ErrorContext(msg string, args ...any) {
	l.Logger.Error(msg, args...)
}

func (l *DefaultLogger) Debug(ctx context.Context, msg string, args ...any) {
	l.Logger.Debug(msg, args...)
}

func (l *DefaultLogger) Info(ctx context.Context, msg string, args ...any) {
	l.Logger.Info(msg, args...)
}

func (l *DefaultLogger) Warn(ctx context.Context, msg string, args ...any) {
	l.Logger.Warn(msg, args...)
}

func (l *DefaultLogger) Error(ctx context.Context, msg string, args ...any) {
	l.Logger.Error(msg, args...)
}
