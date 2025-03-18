package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"ticketing-service/db"
	"ticketing-service/logging"
	"ticketing-service/models"
	"ticketing-service/repositories"
)

type ReservationHandler struct {
	logger logging.Logging
	br     repositories.BookingRepository
	pr     repositories.PassengerRepository
	tr     repositories.TicketRepository
}

func NewReservationHandler(db db.Database, logger logging.Logging) ReservationHandler {
	return ReservationHandler{
		logger: logger,
		br:     repositories.NewBookingRespoitory(db, logger),
		pr:     repositories.NewPassengerRespoitory(db, logger),
		tr:     repositories.NewTicketRepository(db, logger),
	}
}

func (rh *ReservationHandler) ValidateReservation(request []byte) (bool, error) {
	rh.logger.Debug("starting validation process", "reservationHandler.go")

	var reservationObj models.Reservations
	err := json.Unmarshal(request, &reservationObj)
	if err != nil {
		rh.logger.Debug("could not unmarshal reservation request during validation", "reservationHandler.go")
		return false, err
	}
	for _, res := range reservationObj.Reservations {
		rh.logger.Debug(fmt.Sprintf("checking reservation for pax: %s", res.Passenger), "reservationHandler.go")

		validRes, err := rh.checkReservation(res)
		if !validRes || err != nil {
			rh.logger.Info("could not validate reservation: seats already resevered", "reservationHandler.go")
			return false, errors.New("could not create reservation")
		}
	}
	return true, nil
}

/*
**

	checkReservation:
	1. Check to see if the Service exists
	2. Verify the seat (Seat Number + Carriage + Seat Type) is available to book

**
*/
func (rh *ReservationHandler) checkReservation(reservation models.Reservation) (bool, error) {
	for _, route := range reservation.Routes {
		validRes, err := rh.br.ValidateBooking(&route)
		if !validRes || err != nil {
			rh.logger.Debug(fmt.Sprintf("could not validate reservation: seats already resevered for service: %d", route.ServiceNo), "reservationHandler.go")
			return false, errors.New("could not create reservation")
		}
	}
	return true, nil
}

func (rh *ReservationHandler) CreateReservation(request []byte) (*models.Booking, error) {
	var reservationObj models.Reservations
	println("Entering Create Reservation Handler Method....")
	err := json.Unmarshal(request, &reservationObj)
	if err != nil {
		// TODO: implement logging
		panic(err)
	}

	passengers := make([]*models.Passenger, 0)
	for _, res := range reservationObj.Reservations {
		println("Creating Reservation ....")

		tickets, err := rh.tr.Create(res.Routes, res.OriginId, res.DestinationId)
		if err != nil {
			return nil, errors.New("could not create ticket")
		}

		println("creating tickets........")

		passenger, err := rh.pr.Create(tickets)
		if err != nil {
			return nil, errors.New("could not craete passenger")
		}
		println("creating passenger....")

		passengers = append(passengers, passenger)
	}

	booking, err := rh.br.Create(passengers)
	if err != nil {
		return nil, errors.New("could not craete booking")
	}

	return booking, nil
}
