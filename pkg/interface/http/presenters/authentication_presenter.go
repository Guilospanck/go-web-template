/*
Presenters are routes.
*/
package presenters

import (
	"guest-bag-authentication/pkg/application/interfaces"
	"guest-bag-authentication/pkg/infrastructure/adapters"
	httpserver "guest-bag-authentication/pkg/infrastructure/http_server"
	"guest-bag-authentication/pkg/interface/http/handlers"
)

type authenticationPresenter struct {
	logger  interfaces.ILogger
	handler handlers.IAuthenticationHandler
}

func (presenter *authenticationPresenter) Register(httpServer httpserver.IHTTPServer) {
	httpServer.RegisterRoute(
		"POST",
		"/ping",
		adapters.HandlerAdapter(presenter.handler.Ping, presenter.logger),
	)
}

func NewAuthenticationPresenter(logger interfaces.ILogger, handler handlers.IAuthenticationHandler) *authenticationPresenter {
	return &authenticationPresenter{logger, handler}
}
