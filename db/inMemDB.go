package db

import "ticketing-service/models"

type SQLDatabase struct {

	// In memory variables acting as my 'Tables' in a normal SQL database
	Bookings     []*models.Booking
	Passengers   []*models.Passenger
	Reservations []*models.Reservations

	Stations  []*models.Station
	Routes    []*models.Route
	Services  []*models.Service
	Carriages []*models.Carriage
}

func (sql *SQLDatabase) Create(data string, tableName string) error {
	print("creating...")
	return nil
}
