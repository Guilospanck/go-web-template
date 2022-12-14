package factories

import (
	"go-web-template/pkg/application/errors"
	httpserver "go-web-template/pkg/infrastructure/http_server"
	"net/http"
)

type HttpResponseFactory struct{}

type ErrorMessage struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func (HttpResponseFactory) Ok(body interface{}, headers http.Header) httpserver.HttpResponse {
	return httpserver.HttpResponse{
		StatusCode: 200,
		Body:       body,
		Headers:    headers,
	}
}

func (HttpResponseFactory) Created(body interface{}, headers http.Header) httpserver.HttpResponse {
	return httpserver.HttpResponse{
		StatusCode: 201,
		Body:       body,
		Headers:    headers,
	}
}

func (HttpResponseFactory) NoContent(headers http.Header) httpserver.HttpResponse {
	return httpserver.HttpResponse{
		StatusCode: 204,
		Headers:    headers,
	}
}

func (HttpResponseFactory) BadRequest(msg string, headers http.Header) httpserver.HttpResponse {
	return httpserver.HttpResponse{
		StatusCode: 400,
		Body: ErrorMessage{
			StatusCode: 400,
			Message:    msg,
		},
		Headers: headers,
	}
}

func (HttpResponseFactory) Unauthorized(msg string, headers http.Header) httpserver.HttpResponse {
	return httpserver.HttpResponse{
		StatusCode: 401,
		Body: ErrorMessage{
			StatusCode: 401,
			Message:    msg,
		},
		Headers: headers,
	}
}

func (HttpResponseFactory) Forbidden(msg string, headers http.Header) httpserver.HttpResponse {
	return httpserver.HttpResponse{
		StatusCode: 403,
		Body: ErrorMessage{
			StatusCode: 403,
			Message:    msg,
		},
		Headers: headers,
	}
}

func (HttpResponseFactory) NotFound(msg string, headers http.Header) httpserver.HttpResponse {
	return httpserver.HttpResponse{
		StatusCode: 404,
		Body: ErrorMessage{
			StatusCode: 404,
			Message:    msg,
		},
		Headers: headers,
	}
}

func (HttpResponseFactory) Conflict(msg string, headers http.Header) httpserver.HttpResponse {
	return httpserver.HttpResponse{
		StatusCode: 409,
		Body: ErrorMessage{
			StatusCode: 409,
			Message:    msg,
		},
		Headers: headers,
	}
}

func (HttpResponseFactory) InternalServerError(msg string, headers http.Header) httpserver.HttpResponse {
	return httpserver.HttpResponse{
		StatusCode: 500,
		Body: ErrorMessage{
			StatusCode: 500,
			Message:    msg,
		},
		Headers: headers,
	}
}

func (HttpResponseFactory) GenericResponse(statusCode int, body interface{}, headers http.Header) httpserver.HttpResponse {
	return httpserver.HttpResponse{
		StatusCode: statusCode,
		Body:       body,
		Headers:    headers,
	}
}

func (factory HttpResponseFactory) ErrorResponseMapper(err error, headers http.Header) httpserver.HttpResponse {
	switch err.(type) {
	case errors.BadRequestError:
		return factory.BadRequest(err.Error(), headers)
	case errors.NotFoundError:
		return factory.NotFound(err.Error(), headers)
	case errors.UnauthorizedError:
		return factory.Unauthorized(err.Error(), headers)
	case errors.InternalError:
		return factory.InternalServerError(err.Error(), headers)
	default:
		return factory.InternalServerError(err.Error(), headers)
	}
}

func NewHttpResponseFactory() HttpResponseFactory {
	return HttpResponseFactory{}
}
