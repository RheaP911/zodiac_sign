package models

import (
	"github.com/uadmin/uadmin"
)

// PersonalIfno Model !
type Horoscope struct {
	uadmin.Model
	Name      string `uadmin:"default_value:Random Horoscope"`
	Horoscope string `uadmin:"display_name:Daily Horoscope"`
}

// Save function !
func (h *Horoscope) Save() {
	// Save the model to DB
	uadmin.Save(h)
	// Some other business Logic ...
}
