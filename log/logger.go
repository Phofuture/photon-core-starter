package log

import (
	"context"
	"photon-core-starter/core"
)

type CoreLogger interface {
	Debug(context.Context, string, ...any)
	Info(context.Context, string, ...any)
	Warn(context.Context, string, ...any)
	Error(context.Context, string, ...any)
	DebugContext(string, ...any)
	InfoContext(string, ...any)
	WarnContext(string, ...any)
	ErrorContext(string, ...any)
}

type InitAction func(context.Context) (CoreLogger, error)

var (
	coreLogger     CoreLogger = GetDefaultLogger()
	thisInitAction InitAction = func(context context.Context) (CoreLogger, error) {
		return GetDefaultLogger(), nil
	}
)

var thisShutdownAction core.Action = func(context context.Context) error {
	Logger().Info(context, "logger is shutting down")
	return nil
}

func StartLogger() error {
	logger, err := thisInitAction(context.Background())
	if err != nil {
		return err
	}
	if logger != nil {
		coreLogger = logger
	}
	return nil
}

func ShutdownLogger() error {
	return thisShutdownAction(context.Background())
}

func RegisterInitAction(initAction InitAction) {
	thisInitAction = initAction
}

func RegisterShutdownAction(shutdownAction core.Action) {
	thisShutdownAction = shutdownAction
}

func Logger() CoreLogger {
	return coreLogger
}

func Default() CoreLogger {
	return GetDefaultLogger()
}
