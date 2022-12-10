/*
Handlers are like controllers
*/
package handlers

import (
	"fmt"
	"go-web-template/pkg/application/interfaces"
	"go-web-template/pkg/domain/usecases"
	httpserver "go-web-template/pkg/infrastructure/http_server"
	"go-web-template/pkg/interface/http/factories"
)

type IPingHandler interface {
	Ping(httpRequest httpserver.HttpRequest) httpserver.HttpResponse
}

type pingHandler struct {
	logger              interfaces.ILogger
	httpResponseFactory factories.HttpResponseFactory

	postPingUsecase usecases.IPostPingUsecase
}

func (handler *pingHandler) Ping(httpRequest httpserver.HttpRequest) httpserver.HttpResponse {
	res, err := handler.postPingUsecase.Perform()
	if err != nil {
		handler.logger.Error(err.Error())
		return handler.httpResponseFactory.InternalServerError(err.Error(), nil)
	}

	handler.logger.Info(fmt.Sprintf("Will send test: %v", res))
	return handler.httpResponseFactory.Ok(res, nil)
}

func NewPingHandler(logger interfaces.ILogger, postPingUsecase usecases.IPostPingUsecase) *pingHandler {
	httpResponseFactory := factories.NewHttpResponseFactory()

	return &pingHandler{
		logger,
		httpResponseFactory,
		postPingUsecase,
	}
}
