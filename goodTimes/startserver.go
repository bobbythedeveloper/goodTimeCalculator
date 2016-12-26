package hello

import (
	"fmt"
	"log"
	"net/http"
)

func init() {
	http.HandleFunc("/getKaalam", getKaalam)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}

func getKaalam(w http.ResponseWriter, r *http.Request) {
	latitude := float64(36.7201600)
	longitude := float64(-4.4203400)

	sunrise, sunset := getSunriseAndSunset(latitude, longitude)
	log.Println(sunrise)
	log.Println(sunset)
}
func getSunriseAndSunset(lat float64, longt float64) (string, string) {
	latstr := fmt.Sprintf("%f", lat)
	longstr := fmt.Sprintf("%f", longt)

	//http://api.sunrise-sunset.org/json?lat=36.7201600&lng=-4.4203400&date=2016-12-26
	req, err := http.NewRequest("GET", "http://api.sunrise-sunset.org/json", nil)
	if err != nil {
		log.Print(err)
	}
	q := req.URL.Query()
	q.Add("lat", latstr)
	q.Add("lng", longstr)

	req.URL.RawQuery = q.Encode()

	fmt.Println(req.URL.String())

	// call backend and pass the latigude and logitude
	return "", ""
}
