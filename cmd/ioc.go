/*
  - Inversion of control.
    Basically just instantiate what we're going to use
*/
package cmd

import (
	"context"
	httpserver "guest-bag-authentication/pkg/infrastructure/http_server"
	"guest-bag-authentication/pkg/infrastructure/logger"
	"guest-bag-authentication/pkg/interface/http/handlers"
	"guest-bag-authentication/pkg/interface/http/presenters"
)

type Container struct {
	httpServer              httpserver.IHTTPServer
	authenticationPresenter presenters.IRoutes
}

func NewContainer(ctx context.Context) *Container {
	logger := logger.NewLogger()

	// instantiate http server
	httpServer := httpserver.NewHTTPServer(logger)

	// Instantiate handlers (controllers) and presenters (routes) for the authentication workflow
	authenticationHandler := handlers.NewAuthenticationHandler(logger)
	authenticationPresenter := presenters.NewAuthenticationPresenter(logger, authenticationHandler)

	return &Container{
		httpServer,
		authenticationPresenter,
	}

}
