package main

import (
	"net/http"
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

func Test_IntegrationScenario_1_2(t *testing.T) {

	httpClient := initialize()
	testDb := db.GetMemDB()

	// No Reservations should exist in the reservation system
	assert.Equal(t, 0, len(testDb.ReservationSystem), "no reservations should exist")

	for range 2 {
		// Simulate POST request to validate and create reservation
		res := httpClient.Post("/reservation", scenario_1_mockrequest)
		if res.GetStatusCode() == http.StatusOK {

			// Assert that reservation was created successfully
			assert.Equal(t, http.StatusOK, res.GetStatusCode(), "Reservation should be created successfully")
		} else {
			t.Log("Validation failed, reservation not created")
			assert.Equal(t, http.StatusInternalServerError, res.GetStatusCode(), "Validation failure should return status 500")
		}
		time.Sleep(5 * time.Second)

		// Only one Reservation should exist in the reservation system, after each mock request
		assert.Equal(t, 1, len(testDb.ReservationSystem), "one reservations should exist")
	}

}

func Test_IntegrationScenario_3_4(t *testing.T) {

	httpClient := initialize()
	testDb := db.GetMemDB()

	// No Reservations should exist in the reservation system
	assert.Equal(t, 0, len(testDb.ReservationSystem), "no reservations should exist")

	for range 2 {
		// Simulate POST request to validate and create reservation
		res := httpClient.Post("/reservation", scenario_2_mockrequest)
		if res.GetStatusCode() == http.StatusOK {
			// Assert that reservation was created successfully
			assert.Equal(t, http.StatusOK, res.GetStatusCode(), "Reservation should be created successfully")
		} else {
			t.Log("Validation failed, reservation not created")
			assert.Equal(t, http.StatusInternalServerError, res.GetStatusCode(), "Validation failure should return status 500")
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
