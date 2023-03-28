package model

import (
	"encoding/json"
	"io"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Flight struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	TakeOffLocation Location           `bson:"startLoc,omitempty" json:"startLoc"`
	LandingLocation Location           `bson:"endLoc,omitempty" json:"endLoc"`
	FlightStartTime time.Time          `bson:"startTime,omitempty" json:"startTime"`
	TicketPrice     float32            `bson:"price,omitempty" json:"price"`
	Capacity        int                `bson:"capacity,omitempty" json:"capacity"`
	TicketsSold     int                `bson:"sold,omitempty" json:"sold"`
	Passengers      []*User            `bson:"passengers,omitempty" json:"passengers"`
}

func (f *Flight) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(f)
}

func (f *Flight) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(f)
}
