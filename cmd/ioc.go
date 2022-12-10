/*
  - Inversion of control.
    Basically just instantiate what we're going to use
*/
package cmd

import (
	"context"
	usecases "go-web-template/pkg/application/usecases/ping"
	httpserver "go-web-template/pkg/infrastructure/http_server"
	"go-web-template/pkg/infrastructure/logger"
	"go-web-template/pkg/interface/http/handlers"
	"go-web-template/pkg/interface/http/presenters"
)

type Container struct {
	httpServer    httpserver.IHTTPServer
	pingPresenter presenters.IRoutes
}

func NewContainer(ctx context.Context) *Container {
	logger := logger.NewLogger()

	// instantiate http server
	httpServer := httpserver.NewHTTPServer(logger)

	// instantiate usecases
	postPingUsecase := usecases.NewPostPingUsecase()

	// Instantiate handlers (controllers) and presenters (routes) for the ping workflow
	pingHandler := handlers.NewPingHandler(logger, postPingUsecase)
	pingPresenter := presenters.NewPingPresenter(logger, pingHandler)

	return &Container{
		httpServer,
		pingPresenter,
	}

}
