package db

import (
	"ticketing-service/models"
	"time"
)

var stations = []*models.Station{
	{Id: 1, Name: "London"},
	{Id: 2, Name: "Paris"},
	{Id: 3, Name: "Amsterdam"},
	{Id: 4, Name: "Berlin"},
	{Id: 5, Name: "Arras"},
	{Id: 6, Name: "Lille"},
	{Id: 7, Name: "Brussels"},
	{Id: 8, Name: "Antwerp"},
	{Id: 9, Name: "Hengelo"},
	{Id: 10, Name: "Hannover"},
	{Id: 99, Name: "X"},
}

var routes = []*models.Route{
	{RouteId: 1, Origin: "Paris", OriginId: 2, Destination: "London", DestinationId: 1,
		Stops: []models.Stops{
			{StationId: 5, Distance: 180, Position: 1},
			{StationId: 6, Distance: 50, Position: 2},
			{StationId: 1, Distance: 280, Position: 3}}},

	{RouteId: 2, Origin: "Paris", OriginId: 2, Destination: "Amsterdam", DestinationId: 3,
		Stops: []models.Stops{
			{StationId: 7, Distance: 300, Position: 1},
			{StationId: 8, Distance: 50, Position: 2},
			{StationId: 3, Distance: 160, Position: 3}}},

	{RouteId: 3, Origin: "Amsterdam", OriginId: 3, Destination: "Berlin", DestinationId: 4,
		Stops: []models.Stops{
			{StationId: 8, Distance: 150, Position: 1},
			{StationId: 10, Distance: 230, Position: 2},
			{StationId: 4, Distance: 285, Position: 3}}},

	{RouteId: 4, Origin: "London", OriginId: 1, Destination: "Paris", DestinationId: 2,
		Stops: []models.Stops{
			{StationId: 1, Distance: 280, Position: 1},
			{StationId: 6, Distance: 50, Position: 2},
			{StationId: 5, Distance: 180, Position: 3}}},
}

var services = []*models.Service{
	{ServiceNo: 5160, RouteId: 2, Date: parseTime("2025-04-01 15:30:00"), CarriageIds: []string{"A"}},
	{ServiceNo: 3215, RouteId: 4, Date: parseTime("2025-04-01 15:30:00"), CarriageIds: []string{"H", "N"}},
	{ServiceNo: 6821, RouteId: 2, Date: parseTime("2025-04-01 15:30:00"), CarriageIds: []string{"A", "T"}},
}

var carriages = []*models.Carriage{
	{CarriageId: "A", Seats: []models.Seats{{SeatNo: 11, Type: "F"}, {SeatNo: 12, Type: "F"}}},
	{CarriageId: "H", Seats: []models.Seats{{SeatNo: 1, Type: "F"}, {SeatNo: 2, Type: "F"}}},
	{CarriageId: "T", Seats: []models.Seats{{SeatNo: 7, Type: "S"}, {SeatNo: 12, Type: "F"}}},
	{CarriageId: "N", Seats: []models.Seats{{SeatNo: 5, Type: "S"}, {SeatNo: 15, Type: "F"}}},
}

func SeedDb(InMemDB *InMemDB) {
	// Seed the database with temp data to meet the scenario requirements

	InMemDB.Stations = append(InMemDB.Stations, stations...)
	InMemDB.Routes = append(InMemDB.Routes, routes...)
	InMemDB.Services = append(InMemDB.Services, services...)
	InMemDB.Carriages = append(InMemDB.Carriages, carriages...)

}

func parseTime(dateTime string) time.Time {
	timeLayout := "2006-01-02 15:04:05"

	dt, err := time.Parse(timeLayout, dateTime)
	if err != nil {
		// In a real-world setting, this would be handled differently
		return time.Now()
	}

	return dt
}
