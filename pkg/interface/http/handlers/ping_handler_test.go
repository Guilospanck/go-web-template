package handlers

import (
	usecases "go-web-template/pkg/application/usecases/ping"
	"go-web-template/pkg/domain/dtos"
	httpserver "go-web-template/pkg/infrastructure/http_server"
	"go-web-template/pkg/infrastructure/logger"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PingHandler_Ping(t *testing.T) {
	t.Run("Should return HTTP 200 (OK) response with right value", func(t *testing.T) {
		// arrange
		expectedValue := dtos.PingResponseDTO{
			Pong: 1,
		}
		expectedHTTPStatusCode := 200

		headers := http.Header{}
		headers.Add("Content-Type", "application/json")
		httpRequest := httpserver.HttpRequest{
			Body:    []byte{},
			Headers: headers,
		}
		logger := logger.NewLoggerSpy()
		postPingUsecase := usecases.NewPostPingUsecase()

		// act
		response := NewPingHandler(logger, postPingUsecase).Ping(httpRequest)

		// assert
		assert.Equal(t, response.Body, expectedValue)
		assert.Equal(t, response.StatusCode, expectedHTTPStatusCode)
	})
}
