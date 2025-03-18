package repositories

import (
	"testing"
	"ticketing-service/logging"
	"ticketing-service/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock Database interface
type MockDatabase struct {
	mock.Mock
}

func (m *MockDatabase) Create(data any, tableName string) error {
	args := m.Called(data, tableName)
	return args.Error(0)
}

func (m *MockDatabase) Get(data any, tableName string) (any, error) {
	args := m.Called(data, tableName)
	return args.Get(0), args.Error(1)
}

func (m *MockDatabase) List(tableName string) (any, error) {
	args := m.Called(tableName)
	return args.Get(0), args.Error(1)
}

// Test_VerifyBooking tests the ValidateBooking method in bookingRepository
func Test_VerifyBooking(t *testing.T) {
	dbMock := new(MockDatabase)
	logger := &logging.Logs{}
	logging.SetDebugMode(logger)

	br := NewBookingRespoitory(dbMock, logger)

	route := models.Routes{
		ServiceNo: 5160,
		SeatNo:    11,
		Carriage:  "A",
		SeatType:  "F",
	}

	dbMock.On("Get", route.ServiceNo, "service").Return(true, nil).Once()
	dbMock.On("Get", models.SeatLocation{
		SeatNo:   route.SeatNo,
		Carriage: route.Carriage,
		SeatType: route.SeatType,
	}, "seat").Return(true, nil).Once()

	valid, err := br.ValidateBooking(&route)

	assert.NoError(t, err, "should not error")
	assert.True(t, valid, "should be valid")
}

// Test_Create tests the Create method in bookingRepository
func Test_Create(t *testing.T) {
	dbMock := new(MockDatabase)
	logger := &logging.Logs{}
	logging.SetDebugMode(logger)

	br := NewBookingRespoitory(dbMock, logger)

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

	expectedBooking := models.Booking{
		Id:            0, // set to 0 for comparison
		OriginId:      2,
		DestinationId: 3,
		Passengers:    passengers,
	}

	dbMock.On("Create", mock.Anything, "booking").Return(nil).Once()
	dbMock.On("Create", mock.Anything, "reservation").Return(nil).Once()

	bookRes, err := br.Create(passengers, 2, 3)

	// Reset dynamically generated ID for comparison
	bookRes.Id = 0

	// Assertions
	assert.NoError(t, err, "should not error")
	assert.Equal(t, expectedBooking, *bookRes, "should be equal to the expected booking")
}
