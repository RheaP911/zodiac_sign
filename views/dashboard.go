package views

import (
	"math/rand"
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
	horoscopeCount := len(horoscope)
	randomHoroscopeIndex := randomizeNumber(horoscopeCount)
	randomHoroscope := horoscope[randomHoroscopeIndex]
	c["Horoscope"] = []models.Horoscope{randomHoroscope}

	uadmin.All(&numbers) // Retrieve all available horoscopes
	numbersCount := len(numbers)
	randomNumberIndex := randomizeNumber(numbersCount)
	randomNumber := numbers[randomNumberIndex]
	c["Numbers"] = []models.Numbers{randomNumber}


	uadmin.All(&colors)
	colorsCount := len(colors)
	randomColorIndex := randomizeNumber(colorsCount)
	randomColor := colors[randomColorIndex]
	c["Colors"] = []models.Colors{randomColor}

	// Render the home filepath and pass the context data object to the HTML file.
	uadmin.RenderHTML(w, r, "templates/dashboard.html", c)
}

func randomizeNumber(number int) int {
	// Generate a random number between 0 and the given number
	randomized := rand.Intn(number) + 1

	return randomized
}
