package models

import (
	"github.com/uadmin/uadmin"
)

// Info Model
type Colors struct {
	uadmin.Model
	Name          string `uadmin:"display_name:Color"`
	Hex string
	Description string
}

func (co *Colors) Save() {
	// Save the model to DB
	uadmin.Save(co)
	// Some other business Logic ...
}
