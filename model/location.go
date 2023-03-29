package model

import (
	"encoding/json"
	"io"
)

type Location struct {
	Country string `bson:"country,omitempty" json:"country"`
	City    string `bson:"country,omitempty" json:"city"`
	Airport string `bson:"airport,omitempty" json:"airport"`
}

type Locations []*Location

func (l *Locations) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(l)
}

func (l *Location) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(l)
}

func (l *Location) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(l)
}
