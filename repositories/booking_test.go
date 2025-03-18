package repositories

import (
	"testing"
	"ticketing-service/db"
	"ticketing-service/logging"
	"ticketing-service/models"

	"github.com/stretchr/testify/assert"
)

func Test_VerifyBooking(t *testing.T) {
	br := newBookingRepo()
	route := models.Routes{
		ServiceNo: 5160,
		SeatNo:    11,
		Carriage:  "A",
		SeatType:  "F",
	}

	valid, err := br.ValidateBooking(&route)
	assert.NoError(t, err, "should not error")
	assert.Equal(t, valid, true, "should both be true")

}

func Test_Create(t *testing.T) {
	br := newBookingRepo()

	tic := []models.Ticket{
		{
			TicketNo:      12,
			ServiceNo:     3215,
			SeatNo:        1,
			Carriage:      "H",
			SeatType:      "F",
			OriginId:      1,
			DestinationId: 2,
		},
		{
			TicketNo:      13,
			ServiceNo:     6821,
			SeatNo:        1,
			Carriage:      "A",
			SeatType:      "F",
			OriginId:      2,
			DestinationId: 3,
		},
	}

	var tickets []*models.Ticket
	for i := range tic {
		tickets = append(tickets, &tic[i])
	}

	var passengers []*models.Passenger
	p := models.Passenger{PAXId: 1, Tickets: tickets}
	passengers = append(passengers, &p)

	booking := models.Booking{
		Id:            0,
		OriginId:      2,
		DestinationId: 3,
		Passengers:    passengers,
	}

	bookRes, err := br.Create(passengers, 2, 3)

	// reset id to 0 for testing purposes
	bookRes.Id = 0
	assert.NoError(t, err, "should not error")
	assert.Equal(t, *bookRes, booking, "should both be true")
}

func newBookingRepo() BookingRepository {
	db.Initialize()
	logger := logging.Logs{}
	logging.SetDebugMode(&logger)

	return (NewBookingRespoitory(db.GetMemDB(), &logger))

}
