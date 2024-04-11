package views

import (
	"net/http"
	"time"
	"math/rand"

	"github.com/RheaP911/zodiac_sign/models"
	"github.com/uadmin/uadmin"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request, session *uadmin.Session) {

	type Context struct {
		User               string
		ZodiacSigns        models.ZodiacSigns
		Horoscope models.Horoscope
		Date               string
		Aries models.Aries
		Taurus models.Taurus
		Gemini models.Gemini
	}

	// Call the custom struct and assign the full name in the User field under the context object.
	c := Context{}



	// Getting the user firstname and lastname
	// c.User = session.User.FirstName + " " + session.User.LastName

	// getting the current time
	currentTime := time.Now()

	// Format date in MM/DD/YYYY format
	c.Date = currentTime.Format("January 02, 2006")



	// For Daily Horoscope
	horoscope := []models.Horoscope{}
	uadmin.All(&horoscope)
	horoscopeCount := uadmin.Count(&horoscope, "id > 0")
	// uadmin.Trail(uadmin.DEBUG, "Horscope Count: %v\n", horoscopeCount)
	uadmin.Get(&c.Horoscope, "id = ?", randomizeNumber(horoscopeCount))

	// Call Aries Data
	aries := []models.Aries{}
	uadmin.All(&aries)
	ariesCount := uadmin.Count(&aries, "id > 0")
	// uadmin.Trail(uadmin.DEBUG, "Horscope Count: %v\n", horoscopeCount)
	uadmin.Get(&c.Aries, "id = ?", randomizeNumber(ariesCount))

	// Call Gemini Data
	gemini := []models.Gemini{}
	uadmin.All(&gemini)
	geminiCount := uadmin.Count(&gemini, "id > 0")
	// uadmin.Trail(uadmin.DEBUG, "Horscope Count: %v\n", horoscopeCount)
	uadmin.Get(&c.Gemini, "id = ?", randomizeNumber(geminiCount))

	// Call Taurus Data
	taurus := []models.Taurus{}
	uadmin.All(&taurus)
	taurusCount := uadmin.Count(&taurus, "id > 0")
	// uadmin.Trail(uadmin.DEBUG, "Horscope Count: %v\n", horoscopeCount)
	uadmin.Get(&c.Taurus, "id = ?", randomizeNumber(taurusCount))



	// for x := range aries {
	// 	uadmin.Preload(&aries[x])
	// }


	// Render the home filepath and pass the context data object to the HTML file.
	uadmin.RenderHTML(w, r, "templates/dashboard.html", c)
}

func randomizeNumber(number int) int {
	// Generate a random number between 0 and the given number
	randomized := rand.Intn(number) + 1

	return randomized
}