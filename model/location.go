package model

import (
	"encoding/json"
	"io"
)

type Location struct {
	Country string `bson:"country,omitempty" json:"country"`
	City    string `bson:"country,omitempty" json:"city"`
	Street  string `bson:"street,omitempty" json:"street"`
	Number  string `bson:"number,omitempty" json:"number"`
}

func (l *Location) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(l)
}

func (l *Location) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(l)
}
