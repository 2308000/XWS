package model

import (
	"encoding/json"
	"io"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name,omitempty" json:"name"`
	Surname   string             `bson:"surname,omitempty" json:"surname"`
	Email     string             `bson:"email,omitempty" json:"email"`
	Password  string             `bson:"password,omitempty" json:"password"`
	Salt      string             `bson:"salt,omitempty" json:"salt"`
	BirthDate time.Time          `bson:"birthDate,omitempty" json:"birthDate"`
	Role      string             `bson:"role,omitempty" json:"role"`
}

type Users []*User

func (u *Users) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
}

func (u *User) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
}

func (u *User) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(u)
}
