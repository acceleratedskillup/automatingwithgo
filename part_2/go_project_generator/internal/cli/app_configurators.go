package cli

import "goprojgen/internal/projgenerator"

var routerTypeConfigurator = func(selected int) projgenerator.AppConfig {
	appConfig := projgenerator.AppConfig{}
	switch selected {
	case 0:
		appConfig.UseRouter = true
		appConfig.RouterType = "Gorilla Mux"
		appConfig.RouterImportPath = "github.com/gorilla/mux"
		appConfig.RouterConstructor = "mux.NewRouter()"
	case 1:
		appConfig.UseRouter = true
		appConfig.RouterType = "HttpRouter"
		appConfig.RouterImportPath = "github.com/julienschmidt/httprouter"
		appConfig.RouterConstructor = "httprouter.New()"
	}
	return appConfig
}

var configTypeConfigurator = func(selected int) projgenerator.AppConfig {
	appConfig := projgenerator.AppConfig{}
	switch selected {
	case 0: // Command-line flags
		appConfig.ConfigSource = "flags"
	case 1: // Environment variables
		appConfig.ConfigSource = "env"
	}
	return appConfig
}
