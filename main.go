package main

import (
	"ticketing-service/api"
	"ticketing-service/logging"
)

func main() {
	// Initialize logging
	logger := logging.Logs{}
	logging.SetDebugMode(&logger)

	// Implement mock Rest interface to accept incoming reservations
	_ = api.NewHTTPClient(&logger)

	logger.Info("Service is starting...", "")
}
