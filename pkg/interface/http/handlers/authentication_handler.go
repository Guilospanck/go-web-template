/*
Handlers are like controllers
*/
package handlers

import (
	"fmt"
	"guest-bag-authentication/pkg/application/interfaces"
	httpserver "guest-bag-authentication/pkg/infrastructure/http_server"
	"guest-bag-authentication/pkg/interface/http/factories"
)

type IAuthenticationHandler interface {
	Ping(httpRequest httpserver.HttpRequest) httpserver.HttpResponse
}

type authenticationHandler struct {
	logger              interfaces.ILogger
	httpResponseFactory factories.HttpResponseFactory
}

func (handler *authenticationHandler) Ping(httpRequest httpserver.HttpRequest) httpserver.HttpResponse {
	test := map[string]int{
		"pong": 1,
	}

	handler.logger.Info(fmt.Sprintf("Will send test: %v", test))
	return handler.httpResponseFactory.Ok(test, nil)
}

func NewAuthenticationHandler(logger interfaces.ILogger) *authenticationHandler {
	httpResponseFactory := factories.NewHttpResponseFactory()

	return &authenticationHandler{
		logger,
		httpResponseFactory,
	}
}
