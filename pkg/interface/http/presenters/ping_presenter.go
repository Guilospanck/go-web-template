/*
Presenters are routes.
*/
package presenters

import (
	"go-web-template/pkg/application/interfaces"
	"go-web-template/pkg/infrastructure/adapters"
	httpserver "go-web-template/pkg/infrastructure/http_server"
	"go-web-template/pkg/interface/http/handlers"
)

type pingPresenter struct {
	logger  interfaces.ILogger
	handler handlers.IPingHandler
}

func (presenter *pingPresenter) Register(httpServer httpserver.IHTTPServer) {
	httpServer.RegisterRoute(
		"POST",
		"/ping",
		adapters.HandlerAdapter(presenter.handler.Ping, presenter.logger),
	)
}

func NewPingPresenter(logger interfaces.ILogger, handler handlers.IPingHandler) *pingPresenter {
	return &pingPresenter{logger, handler}
}
