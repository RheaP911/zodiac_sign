package main

import (
	"net/http"

	"github.com/RheaP911/zodiac_sign/api"
	"github.com/RheaP911/zodiac_sign/models"
	"github.com/RheaP911/zodiac_sign/views"
	"github.com/uadmin/uadmin"
)

func main() {
	// Assign RootURL value as "/admin/" and Site Name as "Login System"
	// NOTE: This code works only if database does not exist yet.
	uadmin.RootURL = "/admin/"
	uadmin.SiteName = "Horoscope"

	uadmin.Register(
		models.ZodiacSigns{},
		models.Horoscope{},
		models.Colors{},
		models.Numbers{},
	)

	http.HandleFunc("/", uadmin.Handler(views.MainHandler))
	http.HandleFunc("/api/", uadmin.Handler(api.Handler))

	// Run the server
	uadmin.StartServer()
}
