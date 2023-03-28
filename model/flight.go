package model

import (
	"time"

	"github.com/google/uuid"
)

type Flight struct {
	ID              uuid.UUID   `json:"id,omitempty"`
	TakeOffLocation string      `json:"takeOffLoc,omitempty"`
	LandingLocation string      `json:"landingLoc,omitempty"`
	FlightStartTime time.Time   `json:"flightStart,omitempty"`
	TicketPrice     float32     `json:"price,omitempty"`
	Capacity        int         `json:"capacity,omitempty"`
	TicketsSold     int         `json:"sold,omitempty"`
	Passengers      []uuid.UUID `json:"therapy,omitempty"`
}
