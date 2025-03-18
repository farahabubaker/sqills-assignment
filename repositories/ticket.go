package repositories

import (
	"errors"
	"math/rand"
	"ticketing-service/db"
	"ticketing-service/logging"
	"ticketing-service/models"
)

type TicketRepository interface {
	Create(routes []models.Routes, orgId int, desId int) ([]*models.Ticket, error)
}

func NewTicketRepository(db db.Database, logger logging.Logging) TicketRepository {
	return &ticketRepository{
		db:              db,
		logger:          logger,
		ticketTableName: "ticket",
	}
}

type ticketRepository struct {
	db              db.Database
	logger          logging.Logging
	ticketTableName string
}

func (t *ticketRepository) Create(routes []models.Routes, orgId int, desId int) ([]*models.Ticket, error) {
	if routes == nil {
		return nil, errors.New("chosen routes for passenger cannot be empty")
	}

	tickets := make([]*models.Ticket, 0)
	for _, route := range routes {
		ticket := models.Ticket{
			TicketNo:      rand.Intn(10000),
			ServiceNo:     route.ServiceNo,
			SeatNo:        route.SeatNo,
			Carriage:      route.Carriage,
			SeatType:      route.SeatType,
			OriginId:      orgId,
			DestinationId: desId,
		}

		tickets = append(tickets, &ticket)
	}

	return tickets, nil
}
