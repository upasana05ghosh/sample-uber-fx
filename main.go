package main

import (
	"context"
	"net/http"

	"webserver/server"

	"webserver/loggerfx"

	"go.uber.org/fx"

	"go.uber.org/zap"
)

//01. Initailly
// func main() {
// 	mux := http.NewServeMux()
//  server.New(mux)
// http.ListenAndServe(":8080", mux)
// }

//02. Using fx DI
// 	fx.New(
// 		fx.Provide(http.NewServeMux()),
// 		fx.Invoke(server.New),
// 		fx.Invoke(registerHooks),
// 	).Run()

func main() {
	/*
		When created, the application immediately executes all the functions passed via Invoke options. To supply these functions with the parameters they need, the application looks for constructors that return the appropriate types; if constructors for any required types are missing or any invocations
	*/
	fx.New(
		fx.Provide(http.NewServeMux), //dependency -  add objects into the dependency graph
		fx.Invoke(server.New),        //initialize the application using any Provided dependencies
		fx.Invoke(registerHooks),     //Lifecycle - these manage the start and stop for any dependencies that need it
		loggerfx.Module,
	).Run()
}

func registerHooks(
	lifecycle fx.Lifecycle, mux *http.ServeMux, logger *zap.SugaredLogger,
) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				logger.Info("Listening on localhost:8080")
				go http.ListenAndServe(":8080", mux)
				return nil
			},
			OnStop: func(context.Context) error {
				return logger.Sync()
			},
		},
	)
}
