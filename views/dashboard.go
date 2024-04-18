package views

import (
	"net/http"
	"time"

	"github.com/RheaP911/zodiac_sign/models"
	"github.com/uadmin/uadmin"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request, session *uadmin.Session) {
	// Call the custom struct and assign the full name in the User field under the context object.
	c := map[string]interface{}{}
	zodiac := []models.ZodiacSigns{}
	horoscope := []models.Horoscope{}
	numbers := []models.Numbers{}
	colors := []models.Colors{}

	currentTime := time.Now()
	// Format date in MM/DD/YYYY format
	dateStr := currentTime.Format("January 2, 2006")

	// and date into the map
	c["Date"] = dateStr

	uadmin.All(&zodiac)
	c["Zodiac"] = zodiac

	uadmin.All(&horoscope) // Retrieve all available horoscopes
	c["Horoscope"] = horoscope

	uadmin.All(&numbers) // Retrieve all available horoscopes
	c["Numbers"] = numbers

	uadmin.All(&colors)
	c["Colors"] = colors

	

	// Render the home filepath and pass the context data object to the HTML file.
	uadmin.RenderHTML(w, r, "templates/dashboard.html", c)
}
