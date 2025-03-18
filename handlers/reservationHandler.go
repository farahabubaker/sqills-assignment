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

func NewReservationHandler(logger logging.Logging) ReservationHandler {
	return ReservationHandler{
		logger: logger,
		br:     repositories.NewBookingRespoitory(db.GetMemDB(), logger),
		pr:     repositories.NewPassengerRespoitory(db.GetMemDB(), logger),
		tr:     repositories.NewTicketRepository(db.GetMemDB(), logger),
	}
}

func (rh *ReservationHandler) CreateReservation(request []byte) ([]byte, error) {
	rh.logger.Debug("starting reservation creation", "reservationHandler.go")

	var reservationObj models.Reservations
	err := json.Unmarshal(request, &reservationObj)
	if err != nil {
		rh.logger.Debug("could not unmarshal reservation request during validation", "reservationHandler.go")
		return nil, err
	}

	// First Validate the reservation
	rh.logger.Debug("starting validation process", "reservationHandler.go")
	for _, res := range reservationObj.Reservations {
		rh.logger.Debug(fmt.Sprintf("checking reservation for pax: %s", res.Passenger), "reservationHandler.go")

		validRes, err := rh.checkReservation(res)
		if !validRes || err != nil {
			rh.logger.Info("could not validate reservation: seats already resevered", "reservationHandler.go")
			return nil, errors.New("could not create reservation")
		}
	}

	// Then, create the reservation
	passengers := make([]*models.Passenger, 0)
	for _, res := range reservationObj.Reservations {
		rh.logger.Debug(fmt.Sprintf("creating reservation for pax: %s", res.Passenger), "reservationHandler.go")

		tickets, err := rh.tr.Create(res.Routes, res.OriginId, res.DestinationId)
		if err != nil {
			rh.logger.Error("issues during ticket creation", "reservationHandler.go")
			return nil, errors.New("could not create tickets")
		}

		passenger, err := rh.pr.Create(tickets)
		if err != nil {
			rh.logger.Error("issues during passenger creation", "reservationHandler.go")
			return nil, errors.New("ccould not craete passenger")
		}

		passengers = append(passengers, passenger)
	}

	booking, err := rh.br.Create(passengers, reservationObj.OriginId, reservationObj.DestinationId)
	if err != nil {
		rh.logger.Info("issues during ticket creation", "reservationHandler.go")
		return nil, errors.New("could not create booking")
	}

	bookingResp, err := json.Marshal(booking)
	if err != nil {
		rh.logger.Error("could not marshal booking", "reservationHandler.go")
		return nil, errors.New("internal service error")
	}

	return bookingResp, nil
}

func (rh *ReservationHandler) checkReservation(reservation models.Reservation) (bool, error) {
	for _, route := range reservation.Routes {
		validRes, err := rh.br.ValidateBooking(&route)
		if !validRes || err != nil {
			rh.logger.Error(fmt.Sprintf("could not validate reservation: seats already resevered for service: %d", route.ServiceNo), "reservationHandler.go")
			return false, errors.New("could not create reservation")
		}
	}
	return true, nil
}
