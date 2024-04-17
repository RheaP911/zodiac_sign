package api

import (
	"net/http"

	"github.com/RheaP911/zodiac_sign/models"
	"github.com/uadmin/uadmin"
)

func HoroscopeAPIHandler(w http.ResponseWriter, r *http.Request) {
	horoscope := []models.Horoscope{}
	uadmin.All(&horoscope)

	for h := range horoscope {
		uadmin.Preload(&horoscope[h])
	}

	uadmin.ReturnJSON(w, r, horoscope)
}