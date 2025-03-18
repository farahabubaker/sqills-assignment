package repositories

import (
	"testing"
	"ticketing-service/data"
	"ticketing-service/db"
	"ticketing-service/logging"
	"ticketing-service/models"

	"github.com/stretchr/testify/assert"
)

func Test_VerifyBooking(t *testing.T) {
	SQLDB := db.SQLDatabase{}
	data.SeedDb(&SQLDB)
	logger := logging.Logs{}
	logging.SetDebugMode(&logger)

	b := NewBookingRespoitory(&SQLDB, &logger)

	route := models.Routes{
		ServiceNo: 5160,
		SeatNo:    11,
		Carriage:  "A",
		SeatType:  "F",
	}

	valid, err := b.ValidateBooking(&route)
	assert.NoError(t, err, "should not error")
	assert.Equal(t, valid, true, "should both be true")

}
