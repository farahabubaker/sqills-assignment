package models

import "time"

type Station struct {
	Id   int
	Name string
}

type Route struct {
	Origin        string
	OriginId      int
	Destination   string
	DestinationId int
	Stops         []Stops
}

type Stops struct {
	StationId int
	Distance  float64
	Position  int
}

type Service struct {
	ServiceNo int
	Route     Route
	Time      time.Time
	Carriages []Carriage
}

type Carriage struct {
	CarriageId int
	Seats      []Seats
}

type Seats struct {
	SeatNo int
	Type   string
}
