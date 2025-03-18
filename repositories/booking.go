package repositories

import (
	"errors"
	"math/rand"
	"ticketing-service/db"
	"ticketing-service/logging"
	"ticketing-service/models"
)

type BookingRepository interface {
	Create(booking []*models.Passenger) (*models.Booking, error)
	Update(booking *models.Booking) (*models.Booking, error)
	Get(id string) (*models.Booking, error)
	Delete(id string) error
	List() ([]*models.Booking, error)
	ValidateBooking(route *models.Routes) (bool, error)
}

// Ideally you would have configuration setup somewhere instead of hardcoding these values
func NewBookingRespoitory(db db.Database, logger logging.Logging) BookingRepository {
	return &bookingRepository{
		db:               db,
		logger:           logger,
		bookingTableName: "booking",
	}
}

type bookingRepository struct {
	// ctx context.Context
	db               db.Database
	logger           logging.Logging
	bookingTableName string
}

func (b *bookingRepository) Create(passengers []*models.Passenger) (*models.Booking, error) {
	println("creating booking.....")
	if passengers == nil {
		return nil, errors.New("passengers cannot be nil")
	}

	booking := models.Booking{
		Id:         rand.Intn(10000),
		Passengers: passengers,
	}

	err := b.db.Create(booking, b.bookingTableName)
	if err != nil {
		// TODO: implement errors for logging
		return nil, errors.New("this is an error")
	}

	err = b.db.Create(booking, "reservation")
	if err != nil {
		// TODO: implement errors for logging
		return nil, errors.New("this is an error")
	}

	return nil, nil
}

func (b *bookingRepository) List() ([]*models.Booking, error) {
	bookingsData, err := b.db.List(b.bookingTableName)
	if err != nil {
		return nil, errors.New("could not retrieve all bookings")
	}

	bookings, ok := bookingsData.([]*models.Booking)
	if !ok {
		return nil, errors.New("could not retrieve all bookings")
	}

	return bookings, nil
}

func (b *bookingRepository) ValidateBooking(route *models.Routes) (bool, error) {
	if route == nil {
		return false, errors.New("route cannot be null when verifying booking")
	}

	// Verify if service is available
	validServiceData, err := b.db.Get(route.ServiceNo, "service")

	validService, ok := validServiceData.(bool)
	if !ok {
		return false, errors.New("failed to veirfy booking")
	}

	if !validService || err != nil {
		return false, errors.New("could not verify booking")
	}

	seatLocation := models.SeatLocation{
		SeatNo:   route.SeatNo,
		Carriage: route.Carriage,
		SeatType: route.SeatType,
	}

	validSeatData, err := b.db.Get(seatLocation, "seat")

	validSeat, ok := validSeatData.(bool)
	if !ok {
		return false, errors.New("failed to veirfy booking")
	}

	if !validSeat || err != nil {
		return false, errors.New("could not verify booking")
	}
	return true, nil
}

// These would be other methods I would create
func (b *bookingRepository) Update(booking *models.Booking) (*models.Booking, error) {
	return nil, nil
}
func (b *bookingRepository) Get(id string) (*models.Booking, error) {
	return nil, nil
}
func (b *bookingRepository) Delete(id string) error {
	return nil
}
