package main

import (
	"os"
	"ticketing-service/api"
	"ticketing-service/data"
	"ticketing-service/db"
	"ticketing-service/logging"
	"time"
)

func main() {

	// Step 1: Need to setup db and seed it
	SQLDB := db.SQLDatabase{}
	data.SeedDb(&SQLDB)

	// Step 2: setup logging
	logger := logging.Logs{}
	logging.SetDebugMode(&logger)

	// Step 2: Create mock reservation obj from file
	v1_incoming, err := os.ReadFile("Scenario_1_IncomingReq.json")
	if err != nil {
		logger.Error("could not open file: Scenario_2_IncomingReq.json", "main.go")
	}

	// v2_incoming, err := os.ReadFile("Scenario_2_IncomingReq.json")
	// if err != nil {
	// 	logger.Error("could not open file: Scenario_2_IncomingReq.json", "main.go")
	// }

	// Step 3: Implement Rest interface to accept incoming reservations
	httpClient := api.NewHTTPClient(&SQLDB, &logger)

	// Step 4: Call the POST method to handle an incoming reservation
	// Repeat twice to show the reservation failing the second time, due to the seats already being booked
	for range 2 {
		res := httpClient.Get("/reservation/validate", v1_incoming)
		if res.GetStatusCode() == 200 {
			logger.Info("validation complete: success", "main.go")
			httpClient.Post("/reservation", v1_incoming)
		} else {
			logger.Debug("validation complete: unsuccessful", "main.go")
		}
		time.Sleep(5 * time.Second)
	}

	// Repeat twice to show the reservation failing the second time, due to the seats already being booked
	// for range 2 {
	// 	res := httpClient.Get("/reservation/validate", v2_incoming)
	// 	if res.GetStatusCode() == 200 {
	// 		logger.Info("validation complete: success", "main.go")
	// 		httpClient.Post("/reservation", v2_incoming)
	// 	} else {
	// 		logger.Debug("validation complete: unsuccessful", "main.go")
	// 	}
	// 	time.Sleep(5 * time.Second)
	// }
}
