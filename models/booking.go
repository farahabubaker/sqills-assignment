package models

type Reservations struct {
	Reservations []Reservation `json:"reservations"`
}

type Reservation struct {
	Passenger     string   `json:"pax"`
	OriginId      int      `json:"orgId"`
	DestinationId int      `json:"desId"`
	Routes        []Routes `json:"routes"`
}

type Routes struct {
	ServiceNo int    `json:"service"`
	SeatNo    int    `json:"seat"`
	Carriage  string `json:"carriage"`
	SeatType  string `json:"type"`
}

type SeatLocation struct {
	SeatNo   int
	Carriage string
	SeatType string
}

// type ReservationSystem struct {
// 	Bookings []Booking
// }

type Booking struct {
	Id            int
	OriginId      int
	DestinationId int
	Passengers    []*Passenger
}

type Passenger struct {
	PAXId   int
	Tickets []*Ticket
}

type Ticket struct {
	TicketNo      int
	ServiceNo     int
	SeatNo        int
	Carriage      string
	SeatType      string
	OriginId      int
	DestinationId int
}
