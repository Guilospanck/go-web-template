package httpserver

import (
	"context"
	"fmt"
	"go-web-template/pkg/application/interfaces"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type IHTTPServer interface {
	Setup(ctx context.Context)
	RegisterRoute(method string, endpoint string, handler ...gin.HandlerFunc) error
	Run() error
}

type HTTPServer struct {
	logger  interfaces.ILogger
	ctx     context.Context
	router  *gin.Engine
	address string
	server  *http.Server
}

func (server *HTTPServer) Setup(ctx context.Context) {
	server.ctx = ctx

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	server.address = fmt.Sprintf("%s:%s", host, port)

	server.router = gin.New()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{"*"}
	server.router.Use(cors.New(corsConfig))

	server.server = &http.Server{
		Addr:    server.address,
		Handler: server.router,
	}
}

func (server *HTTPServer) RegisterRoute(method string, endpoint string, handler ...gin.HandlerFunc) error {
	switch method {
	case "POST":
		server.router.POST(endpoint, handler...)
	case "GET":
		server.router.GET(endpoint, handler...)
	case "PUT": // Put updates the entire resource
		server.router.PUT(endpoint, handler...)
	case "PATCH": // Patch updates partially
		server.router.PATCH(endpoint, handler...)
	case "DELETE":
		server.router.DELETE(endpoint, handler...)
	default:
		return fmt.Errorf("Method now allowed")
	}

	return nil
}

func (server *HTTPServer) Run() error {
	server.logger.Info(fmt.Sprintf("Server running at http://%s", server.address))
	err := server.server.ListenAndServe()
	if err != nil {
		server.logger.Error(fmt.Sprintf("Error while trying to serve HTTP: %s", err.Error()))
		return err
	}
	return nil
}

func NewHTTPServer(logger interfaces.ILogger) *HTTPServer {
	return &HTTPServer{logger: logger}
}
