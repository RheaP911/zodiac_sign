package models

import (
	"time"

	"github.com/uadmin/uadmin"
)

// Info Model
type Info struct {
	uadmin.Model
	Name          uadmin.User
	NameID        uint
	Birthdate     time.Time
	Nickname      string
	ZodiacSigns   ZodiacSigns
	ZodiacSignsID uint
}

func (i *Info) String() string {
	// user := uadmin.User{}
	// uadmin.Get(&user, "id=? ", i.Name)
	return i.Nickname
}

