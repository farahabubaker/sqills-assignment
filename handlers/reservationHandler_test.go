package handlers

import (
	"testing"

	"ticketing-service/logging"
	"ticketing-service/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock BookingRepository
type MockBookingRepository struct {
	mock.Mock
}

func (m *MockBookingRepository) Create(passengers []*models.Passenger, orgId, desId int) (*models.Booking, error) {
	args := m.Called(passengers, orgId, desId)
	if args.Get(0) != nil {
		return args.Get(0).(*models.Booking), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockBookingRepository) Update(booking *models.Booking) (*models.Booking, error) {
	args := m.Called(booking)
	if args.Get(0) != nil {
		return args.Get(0).(*models.Booking), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockBookingRepository) Get(id string) (*models.Booking, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*models.Booking), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockBookingRepository) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockBookingRepository) List() ([]*models.Booking, error) {
	args := m.Called()
	if args.Get(0) != nil {
		return args.Get(0).([]*models.Booking), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockBookingRepository) ValidateBooking(route *models.Routes) (bool, error) {
	args := m.Called(route)
	return args.Bool(0), args.Error(1)
}

// Mock PassengerRepository
type MockPassengerRepository struct {
	mock.Mock
}

func (m *MockPassengerRepository) Create(tickets []*models.Ticket) (*models.Passenger, error) {
	args := m.Called(tickets)
	if args.Get(0) != nil {
		return args.Get(0).(*models.Passenger), args.Error(1)
	}
	return nil, args.Error(1)
}

// Mock TicketRepository
type MockTicketRepository struct {
	mock.Mock
}

func (m *MockTicketRepository) Create(routes []models.Routes, orgId, desId int) ([]*models.Ticket, error) {
	args := m.Called(routes, orgId, desId)
	if args.Get(0) != nil {
		return args.Get(0).([]*models.Ticket), args.Error(1)
	}
	return nil, args.Error(1)
}

func TestValidateReservation(t *testing.T) {
	mockLogger := &logging.Logs{}
	logging.SetDebugMode(mockLogger)

	mockBookingRepo := new(MockBookingRepository)
	mockPassengerRepo := new(MockPassengerRepository)
	mockTicketRepo := new(MockTicketRepository)

	handler := ReservationHandler{
		logger: mockLogger,
		br:     mockBookingRepo,
		pr:     mockPassengerRepo,
		tr:     mockTicketRepo,
	}

	// Test request JSON
	request := []byte(`{
		"reservations": [
			{
				"pax": "John",
				"orgId": 2,
				"desId": 3,
				"routes": [
					{
						"service": 5160,
						"seat": 11,
						"carriage": "A",
						"type": "F"
					}
				]
			}
		]
	}`)

	// Mock ValidateBooking
	mockBookingRepo.On("ValidateBooking", mock.AnythingOfType("*models.Routes")).Return(true, nil)

	valid, err := handler.ValidateReservation(request)

	assert.NoError(t, err)
	assert.True(t, valid)
	mockBookingRepo.AssertExpectations(t)
}

func TestCreateReservation(t *testing.T) {
	mockLogger := &logging.Logs{}
	logging.SetDebugMode(mockLogger)

	mockBookingRepo := new(MockBookingRepository)
	mockPassengerRepo := new(MockPassengerRepository)
	mockTicketRepo := new(MockTicketRepository)

	handler := ReservationHandler{
		logger: mockLogger,
		br:     mockBookingRepo,
		pr:     mockPassengerRepo,
		tr:     mockTicketRepo,
	}

	// Test request JSON
	request := []byte(`{
		"reservations": [
			{
				"pax": "John",
				"orgId": 2,
				"desId": 3,
				"routes": [
					{
						"service": 5160,
						"seat": 11,
						"carriage": "A",
						"type": "F"
					}
				]
			}
		]
	}`)

	// Mock Repo behavior
	mockTicketRepo.On("Create", mock.Anything, mock.Anything, mock.Anything).Return([]*models.Ticket{{}}, nil)
	mockPassengerRepo.On("Create", mock.Anything).Return(&models.Passenger{}, nil)
	mockBookingRepo.On("Create", mock.Anything, mock.Anything, mock.Anything).Return(&models.Booking{}, nil)

	response, err := handler.CreateReservation(request)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	mockTicketRepo.AssertExpectations(t)
	mockPassengerRepo.AssertExpectations(t)
	mockBookingRepo.AssertExpectations(t)
}
