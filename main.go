package main

import (
	"ticketing-service/api"
	"ticketing-service/db"
	"ticketing-service/logging"
)

func main() {

	// Initialize the in memory database and seed it
	db.Initialize()

	// Initialize logging
	logger := logging.Logs{}
	logging.SetDebugMode(&logger)

	// Implement mock Rest interface to accept incoming reservations
	_ = api.NewHTTPClient(&logger)

	// run_scenarios(httpClient, logger)
}
