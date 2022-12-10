package presenters

import httpserver "go-web-template/pkg/infrastructure/http_server"

type IRoutes interface {
	Register(httpServer httpserver.IHTTPServer)
}
