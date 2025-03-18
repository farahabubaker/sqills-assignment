package repositories

import (
	"errors"
	"math/rand"
	"ticketing-service/db"
	"ticketing-service/logging"
	"ticketing-service/models"
)

type PassengerRepository interface {
	Create(tickets []*models.Ticket) (*models.Passenger, error)
}

func NewPassengerRespoitory(db db.Database, logger logging.Logging) PassengerRepository {
	return &passengerRepository{
		db:                 db,
		logger:             logger,
		passangerTableName: "passenger",
	}
}

type passengerRepository struct {
	db                 db.Database
	logger             logging.Logging
	passangerTableName string
}

func (p *passengerRepository) Create(tickets []*models.Ticket) (*models.Passenger, error) {
	if tickets == nil {
		return nil, errors.New("tickets cannot be empty")
	}

	passenger := &models.Passenger{
		PAXId:   rand.Intn(100),
		Tickets: tickets,
	}

	return passenger, nil
}
