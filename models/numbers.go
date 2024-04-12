package models

import (
	"github.com/uadmin/uadmin"
)

// Info Model
type Numbers struct {
	uadmin.Model
	Name          string `uadmin:"display_name:Number"`
	Number string
	Description string
}

func (n *Numbers) Save() {
	// Save the model to DB
	uadmin.Save(n)
	// Some other business Logic ...
}
