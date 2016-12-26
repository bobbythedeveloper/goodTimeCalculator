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
	latitude := float64(5000)
	longitude := float64(6000)

	sunrise, sunset := getSunriseAndSunset(latitude, longitude)
	log.Println(sunrise)
	log.Println(sunset)
}
func getSunriseAndSunset(lat float64, longt float64) (string, string) {

	// call backend and pass the latigude and logitude
	return "", ""
}
