package models

type Reservations struct {
	Reservations  []Reservation `json:"reservations"`
	OriginId      int           `json:"orgId"`
	DestinationId int           `json:"desId"`
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

type Booking struct {
	Id            int          `json:"id"`
	OriginId      int          `json:"orgId"`
	DestinationId int          `json:"desId"`
	Passengers    []*Passenger `json:"passengers"`
}

type Passenger struct {
	PAXId   int       `json:"paxId"`
	Tickets []*Ticket `json:"tickets"`
}

type Ticket struct {
	TicketNo      int    `json:"ticketno"`
	ServiceNo     int    `json:"service"`
	SeatNo        int    `json:"seat"`
	Carriage      string `json:"carriage"`
	SeatType      string `json:"type"`
	OriginId      int    `json:"orgId"`
	DestinationId int    `json:"desId"`
}

type SeatLocation struct {
	SeatNo   int
	Carriage string
	SeatType string
}
