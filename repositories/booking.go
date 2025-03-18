package repositories

import (
	"errors"
	"fmt"
	"math/rand"
	"ticketing-service/db"
	"ticketing-service/logging"
	"ticketing-service/models"
)

type BookingRepository interface {
	Create(passengers []*models.Passenger, orgId int, desId int) (*models.Booking, error)
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

func (b *bookingRepository) Create(passengers []*models.Passenger, orgId int, desId int) (*models.Booking, error) {

	if passengers == nil {
		return nil, errors.New("passengers cannot be nil")
	}

	booking := models.Booking{
		Id:            rand.Intn(10000),
		OriginId:      orgId,
		DestinationId: desId,
		Passengers:    passengers,
	}

	err := b.db.Create(booking, b.bookingTableName)
	if err != nil {
		b.logger.Error("could not add booking", "booking.go")
		return nil, errors.New("could not add booking")
	}

	err = b.db.Create(booking, "reservation")
	if err != nil {
		b.logger.Error("could not add booking to reservation system", "booking.go")
		return nil, errors.New("could not add booking to reservation system")
	}

	b.logger.Debug(fmt.Sprintf("successfully created booking: %d", booking.Id), "booking.go")

	return &booking, nil
}

func (b *bookingRepository) ValidateBooking(route *models.Routes) (bool, error) {
	if route == nil {
		b.logger.Error("route cannot be null when verifying booking", "booking.go")
		return false, errors.New("route cannot be null when verifying booking")
	}

	// Verify if service is available
	validServiceData, err := b.db.Get(route.ServiceNo, "service")
	validService, ok := validServiceData.(bool)
	if !ok {
		b.logger.Error("could not process data from 'database'", "booking.go")
		return false, errors.New("could not process data from 'database")
	}

	if !validService || err != nil {
		b.logger.Error(fmt.Sprintf("could not validate booking, invalid Service: %d", route.ServiceNo), "booking.go")
		return false, errors.New("service cannot be invalid when booking")
	}

	// Verify if seat is available
	seatLocation := models.SeatLocation{
		SeatNo:   route.SeatNo,
		Carriage: route.Carriage,
		SeatType: route.SeatType,
	}

	validSeatData, err := b.db.Get(seatLocation, "seat")
	validSeat, ok := validSeatData.(bool)
	if !ok {
		b.logger.Error("could not process data from 'database'", "booking.go")
		return false, errors.New("could not process data from 'database")
	}

	if !validSeat || err != nil {
		b.logger.Error(fmt.Sprintf("could not validate booking, invalid Seat: %d, Carriage: %s, or SeatType: %s", route.SeatNo, route.Carriage, route.SeatType), "booking.go")
		return false, errors.New("service cannot be invalid when booking")
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
