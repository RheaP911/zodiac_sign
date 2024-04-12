package views

import (
	"math/rand"
	"net/http"

	"github.com/RheaP911/zodiac_sign/models"
	"github.com/uadmin/uadmin"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request, session *uadmin.Session) {
	// Call the custom struct and assign the full name in the User field under the context object.
	c := map[string]interface{}{}

	zodiac := []models.ZodiacSigns{}
	uadmin.All(&zodiac)
	c["Zodiac"] = zodiac

	// horoscope := []models.Horoscope{}
	// uadmin.All(&horoscope)
	// c["Horoscope"] = horoscope

	// colors := []models.Colors{}
	// uadmin.All(&colors)
	// c["Colors"] = colors

	// Render the home filepath and pass the context data object to the HTML file.
	uadmin.RenderHTML(w, r, "templates/dashboard.html", c)
}

func randomizeNumber(number int) int {
	// Generate a random number between 0 and the given number
	randomized := rand.Intn(number) + 1

	return randomized
}
