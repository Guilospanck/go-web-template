package presenters

import httpserver "guest-bag-authentication/pkg/infrastructure/http_server"

type IRoutes interface {
	Register(httpServer httpserver.IHTTPServer)
}
