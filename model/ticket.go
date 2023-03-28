package model

import (
	"time"

	"github.com/google/uuid"
)

type Ticket struct {
	ID              uuid.UUID `json:"id,omitempty"`
	FlightStartTime time.Time `json:"flightStart,omitempty"`
	TakeOffLocation string    `json:"takeOffLoc,omitempty"`
	LandingLocation string    `json:"landingLoc,omitempty"`
	SeatNumber      int       `json:"seatNumber,omitempty"`
}
