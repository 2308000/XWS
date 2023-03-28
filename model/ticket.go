package model

import (
	"encoding/json"
	"io"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Ticket struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	User       User               `bson:"user,omitempty" json:"user"`
	Flight     Flight             `bson:"flight,omitempty" json:"flight"`
	SeatNumber int                `bson:"seatNumber,omitempty" json:"seatNumber"`
}

func (t *Ticket) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(t)
}

func (t *Ticket) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(t)
}
