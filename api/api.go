package api

import (
	"net/http"
	"strings"

	"github.com/uadmin/uadmin"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/api/")
	path := strings.TrimPrefix(r.URL.Path, "/")

	uadmin.Trail(uadmin.DEBUG, "path", path)

	switch path {
	case "horoscope":
        HoroscopeAPIHandler(w, r)
	default:
		w.Write([]byte("Not Found"))
	}

	// if strings.HasPrefix(r.URL.Path, "/horoscope") {
    //     return
	// }
	
}