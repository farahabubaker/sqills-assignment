package repositories

import (
	"errors"
	"ticketing-service/db"
	"ticketing-service/models"
)

type BookingRepository interface {
	Create(booking *models.Booking) (*models.Booking, error)
	Update(booking *models.Booking) (*models.Booking, error)
	Get(id string) (*models.Booking, error)
	Delete(id string) error
	List() ([]*models.Booking, error)
}

// Ideally you would have configuration setup somewhere instead of hardcoding these values
func NewBookingRespoitory(db db.Database) BookingRepository {
	return &bookingRepository{
		db:               db,
		bookingTableName: "booking",
	}
}

type bookingRepository struct {
	// ctx context.Context
	db               db.Database
	bookingTableName string
}

func (b *bookingRepository) Create(booking *models.Booking) (*models.Booking, error) {
	err := b.db.Create(booking, b.bookingTableName)
	if err != nil {
		// TODO: implement errors for logging
		return nil, errors.New("this is an error")
	}
	return nil, nil
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
	return nil, nil
}
