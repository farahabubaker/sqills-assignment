package main

import (
	"testing"
	"ticketing-service/api"
	"ticketing-service/db"
	"ticketing-service/logging"
	"time"

	"github.com/stretchr/testify/assert"
)

var scenario_1_mockrequest = []byte(`{
    "orgId" : 2,
    "desId" : 3,
    "reservations" : [
        {
            "pax" : "John",
            "orgId" : 2,
            "desId" : 3,
            "routes" : [
                {
                    "service" : 5160,
                    "seat" : 11,
                    "carriage" : "A",
                    "type" : "F"
                }
            ]
        },
        {
            "pax" : "Linda",
            "orgId" : 2,
            "desId" : 3,
            "routes" : [
                {
                    "service" : 5160,
                    "seat" : 12,
                    "carriage" : "A",
                    "type" : "F"
                }
            ]
        }
    ]
}`)

var scenario_2_mockrequest = []byte(`{
    "orgId" : 2,
    "desId" : 3,
    "reservations" : [
        {
            "pax" : "Mimi",
            "orgId" : 1,
            "desId" : 3,
            "routes" : [
                {
                    "service" : 3215,
                    "seat" : 1,
                    "carriage" : "H",
                    "type" : "F"
                },
                {
                    "service" : 6821,
                    "seat" : 1,
                    "carriage" : "A",
                    "type" : "F"
                }
            ]
        },
        {
            "pax" : "Riley",
            "orgId" : 1,
            "desId" : 3,
            "routes" : [
                {
                    "service" : 3215,
                    "seat" : 5,
                    "carriage" : "N",
                    "type" : "S"
                },
                {
                    "service" : 6821,
                    "seat" : 7,
                    "carriage" : "T",
                    "type" : "S"
                }
            ]
        }
    ]
}`)

func Test_IntegrationScenario1(t *testing.T) {

	httpClient := initialize()
	testDb := db.GetMemDB()

	// No Reservations should exist in the reservation system
	assert.Equal(t, 0, len(testDb.ReservationSystem), "no reservations should exist")

	for range 2 {
		// Simulate GET request to validate reservation
		res := httpClient.Get("/reservation/validate", scenario_1_mockrequest)
		if res.GetStatusCode() == 200 {
			// Validation successful, create the reservation
			t.Log("Validation successful, creating reservation")
			res = httpClient.Post("/reservation", scenario_1_mockrequest)

			// Assert that reservation was created successfully
			assert.Equal(t, 200, res.GetStatusCode(), "Reservation should be created successfully")
		} else {
			t.Log("Validation failed, reservation not created")
			assert.Equal(t, 400, res.GetStatusCode(), "Validation failure should return status 400")
		}
		time.Sleep(5 * time.Second)

		// Only one Reservation should exist in the reservation system, after each mock request
		assert.Equal(t, 1, len(testDb.ReservationSystem), "one reservations should exist")
	}

}

func Test_IntegrationScenario2(t *testing.T) {

	httpClient := initialize()
	testDb := db.GetMemDB()

	// No Reservations should exist in the reservation system
	assert.Equal(t, 0, len(testDb.ReservationSystem), "no reservations should exist")

	for range 2 {
		// Simulate GET request to validate reservation
		res := httpClient.Get("/reservation/validate", scenario_2_mockrequest)
		if res.GetStatusCode() == 200 {
			// Validation successful, create the reservation
			t.Log("Validation successful, creating reservation")
			res = httpClient.Post("/reservation", scenario_2_mockrequest)

			// Assert that reservation was created successfully
			assert.Equal(t, 200, res.GetStatusCode(), "Reservation should be created successfully")
		} else {
			t.Log("Validation failed, reservation not created")
			assert.Equal(t, 400, res.GetStatusCode(), "Validation failure should return status 400")
		}
		time.Sleep(5 * time.Second)

		// Only one Reservation should exist in the reservation system, after each mock request
		assert.Equal(t, 1, len(testDb.ReservationSystem), "one reservations should exist")
	}
}

func initialize() api.HTTPClient {
	db.Initialize()

	// Initialize logging
	logger := logging.Logs{}
	logging.SetDebugMode(&logger)

	// Implement mock Rest interface to accept incoming reservations
	return (api.NewHTTPClient(&logger))
}
