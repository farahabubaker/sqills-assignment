package main

import (
	"os"
	"testing"
	"ticketing-service/api"
	"ticketing-service/db"
	"ticketing-service/logging"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_VerifyBooking(t *testing.T) {

	httpclient = initialize()

	// Create mock reservation obj from file
	v1_incoming, err := os.ReadFile("Scenario_1_IncomingReq.json")
	if err != nil {
		logger.Error("could not open file: Scenario_2_IncomingReq.json", "main.go")
	}
	assert.NoError(t, err, "should not error")
	assert.Equal(t, valid, true, "should both be true")

}

func initialize() api.HTTPClient {
	db.Initialize()

	// Initialize logging
	logger := logging.Logs{}
	logging.SetDebugMode(&logger)

	// Implement mock Rest interface to accept incoming reservations
	return (api.NewHTTPClient(&logger))
}

func run_scenarios(httpClient api.HTTPClient, logger logging.Logs) {
	// Create mock reservation obj from file
	v1_incoming, err := os.ReadFile("Scenario_1_IncomingReq.json")
	if err != nil {
		logger.Error("could not open file: Scenario_2_IncomingReq.json", "main.go")
	}
	v2_incoming, err := os.ReadFile("Scenario_2_IncomingReq.json")
	if err != nil {
		logger.Error("could not open file: Scenario_2_IncomingReq.json", "main.go")
	}

	// Call the POST method to handle an incoming reservation
	// Repeat twice to show the reservation failing the second time, due to the seats already being booked
	for range 2 {
		res := httpClient.Get("/reservation/validate", v1_incoming)
		if res.GetStatusCode() == 200 {
			logger.Info("validation complete: success", "main.go")
			httpClient.Post("/reservation", v1_incoming)
		} else {
			logger.Debug("validation complete: unsuccessful", "main.go")
		}
		logger.Info("===============================================\n", "")
		time.Sleep(5 * time.Second)
	}

	// Repeat twice to show the reservation failing the second time, due to the seats already being booked
	for range 2 {
		res := httpClient.Get("/reservation/validate", v2_incoming)
		if res.GetStatusCode() == 200 {
			logger.Info("validation complete: success", "main.go")
			httpClient.Post("/reservation", v2_incoming)
		} else {
			logger.Debug("validation complete: unsuccessful", "main.go")
		}
		logger.Info("===============================================\n", "")

		time.Sleep(5 * time.Second)
	}
}
