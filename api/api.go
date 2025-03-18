package api

import (
	"strings"
	"ticketing-service/handlers"
	"ticketing-service/logging"
)

// Mock response interface
type Response interface {
	GetStatusCode() int
	GetBody() []byte
}

type ResponseDetails struct {
	statusCode int
	body       []byte
}

func (r *ResponseDetails) GetStatusCode() int {
	return r.statusCode
}

func (r *ResponseDetails) GetBody() []byte {
	return r.body
}

// Mock HTTP client
type HTTPClient interface {
	Post(path string, body []byte) Response
	Get(path string, body []byte) Response
}

func NewHTTPClient(logger logging.Logging) HTTPClient {
	return &httpClient{
		logger: logger,
		rh:     handlers.NewReservationHandler(logger),
	}
}

type httpClient struct {
	logger logging.Logging
	rh     handlers.ReservationHandler
}

func (c *httpClient) Post(path string, body []byte) Response {
	if strings.Compare("/reservation", path) == 0 {
		booking, err := c.rh.CreateReservation(body)
		if err != nil {
			c.logger.Error("could not create reservation", "api.go")
			return &ResponseDetails{
				statusCode: 400,
				body:       []byte("invalid reservation request"),
			}
		}
		return &ResponseDetails{
			statusCode: 200,
			body:       booking,
		}
	}
	return &ResponseDetails{
		statusCode: 400,
		body:       []byte("invalid reservation request"),
	}
}

func (c *httpClient) Get(path string, body []byte) Response {
	if strings.Compare("/reservation/validate", path) == 0 {
		valid, err := c.rh.ValidateReservation(body)
		if err != nil || !valid {
			c.logger.Error("could not validate incoming reservation", "api.go")
			return &ResponseDetails{
				statusCode: 400,
				body:       []byte("invalid reservation request"),
			}
		}

	}
	return &ResponseDetails{
		statusCode: 200,
		body:       []byte("reservation validated"),
	}
}
