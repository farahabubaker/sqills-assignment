package db

import (
	"errors"
	"strings"
	"ticketing-service/models"
)

type SQLDatabase struct {

	// In memory variables acting as my 'Tables' in a normal SQL database
	ReservationSystem []*models.Booking
	Passengers        []*models.Passenger

	Stations  []*models.Station
	Routes    []*models.Route
	Services  []*models.Service
	Carriages []*models.Carriage
}

func (sql *SQLDatabase) Create(data any, tableName string) error {

	if strings.Compare(tableName, "booking") == 0 {
		booking, ok := data.(models.Booking)
		if !ok {
			return errors.New("could not add booking")
		}

		sql.ReservationSystem = append(sql.ReservationSystem, &booking)
		println("adding booking to reservation system......")

	}
	return nil
}

func (sql *SQLDatabase) List(tableName string) (any, error) {
	if strings.Compare(tableName, "booking") == 0 {
		return sql.ReservationSystem, nil
	}

	return nil, nil
}

func (sql *SQLDatabase) Get(data any, tableName string) (any, error) {
	if strings.Compare(tableName, "service") == 0 {
		serviceNo, ok := data.(int)
		if !ok {
			return false, errors.New("could not add booking")
		}

		for _, s := range sql.Services {
			if s.ServiceNo == serviceNo {
				return true, nil
			}
		}
	}
	if strings.Compare(tableName, "seat") == 0 {
		seat, ok := data.(models.SeatLocation)
		if !ok {
			return nil, errors.New("could not add booking")
		}

		for _, booking := range sql.ReservationSystem {
			for _, pax := range booking.Passengers {
				if !checkSeat(seat, pax.Tickets) {
					return false, nil
				}
			}
		}

		return true, nil

	}
	return nil, nil
}

func checkSeat(seat models.SeatLocation, tickets []*models.Ticket) bool {
	for _, t := range tickets {
		print("checking tickets")
		if t.SeatNo == seat.SeatNo && t.Carriage == seat.Carriage && t.SeatType == seat.SeatType {
			println("oh no!")
			return false
		}
	}
	return true
}
