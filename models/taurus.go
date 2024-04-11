package models

import (
	"github.com/uadmin/uadmin"
)

// ZodiacSigns Model !
type Taurus struct {
	uadmin.Model
	Name             string `uadmin:"display_name: Zodiac Sign"`
	Date             string 
	LooksRate        string `uadmin:"list_exclude"`
	LooksDesc	string `uadmin:"list_exclude"`
	PersonalityRate  string `uadmin:"list_exclude"`
	PersonDesc	string  `uadmin:"list_exclude"`
	IntelligenceRate string `uadmin:"list_exclude"`
	IntelDesc	string `uadmin:"list_exclude"`
	Description      string `uadmin:"html"`

	Symbol   string `uadmin:"image"`
	StarSign string `uadmin:"image"`

	Compatible string
	Compatible1 string `uadmin:"image"`
	Compatible2 string `uadmin:"image"`
	Compatible3 string `uadmin:"image"`
}

// Save function !
func (t *Taurus) Save() {
	// Save the model to DB
	uadmin.Save(t)
	// Some other business Logic ...
}
