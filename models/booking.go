package models

type Reservations struct {
	Bookings []Booking
}

type Booking struct {
	Id         string
	Passengers []Passenger
}

type Passenger struct {
	PAXId   int
	Tickets []Ticket
}

type Ticket struct {
	TicketNo    int
	SeatNo      string
	Origin      string
	Destination string
	ServiceNo   int
}
