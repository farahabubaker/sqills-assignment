package models

import (
	"time"
)

type Station struct {
	Id   int
	Name string
}

type Route struct {
	RouteId       int
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
	ServiceNo   int
	RouteId     int
	Date        time.Time
	CarriageIds []string
}

type Carriage struct {
	CarriageId string
	Seats      []Seats
}

type Seats struct {
	SeatNo int
	Type   string
}
