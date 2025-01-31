package core

import "context"

type Action func(context context.Context) error

var (
	coreDependencies         = []Action{}
	shutdownCoreDependencies = []Action{}
)

var (
	initAddModuleActions     = []Action{}
	shutdownAddModuleActions = []Action{}
)

var initProjectActions = []Action{}

func RegisterProjectInit(initAction Action) {
	initProjectActions = append(initProjectActions, initAction)
}

func GetProjectInit() []Action {
	return initProjectActions
}

func RegisterCoreDependency(initAction Action) {
	coreDependencies = append(coreDependencies, initAction)
}

func RegisterShutdownCoreDependency(shutdownAction Action) {
	shutdownCoreDependencies = append(shutdownCoreDependencies, shutdownAction)
}

func RegisterAddModule(initAction Action) {
	initAddModuleActions = append(initAddModuleActions, initAction)
}

func RegisterShutdownAddModule(shutdownAction Action) {
	shutdownAddModuleActions = append(shutdownAddModuleActions, shutdownAction)
}

func GetCoreDependencies() []Action {
	return coreDependencies
}

func GetShutdownCoreDependencies() []Action {
	return shutdownCoreDependencies
}

func GetAddModule() []Action {
	return initAddModuleActions
}

func GetShutdownAddModule() []Action {
	return shutdownAddModuleActions
}
