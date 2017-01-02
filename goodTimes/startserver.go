package hello

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
)

func init() {
	http.HandleFunc("/getKaalam", getKaalam)
	http.HandleFunc("/", handler)

}

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Hello, The API usage is /getKaalam?latitude=39.934002&longitude=-74.89099879999998&date=2017-Jan-02")
}

func getKaalam(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := appengine.NewContext(r)
	client := urlfetch.Client(ctx)
	//39.948184, -74.902575
	latitude, err := strconv.ParseFloat(r.URL.Query().Get("latitude"), 64)
	if err != nil {
		log.Println("cannot parse latitude using default latitude 39.948184")
		latitude = float64(39.948184)
		log.Print(err)
	}
	longitude, err := strconv.ParseFloat((r.URL.Query().Get("longitude")), 64)
	if err != nil {
		log.Println("cannot parse longitude using default longitude -74.902575")
		longitude = float64(-74.902575)
		log.Print(err)
	}
	date, err := time.Parse("2006-Jan-2", (r.URL.Query().Get("date")))
	if err != nil {
		log.Println("cannot parse date using default date = today()")
		date = time.Now()
		log.Print(err)
	}

	sunrise, sunset := getSunriseAndSunset(latitude, longitude, date, client)
	rahuKaala := getKallas(sunrise, sunset, date)
	response := ResponseJSON{}
	response.RahuKaalEndTime = rahuKaala.endTime
	response.RahuKaalStartTime = rahuKaala.startTime

	jsonresponse, err := json.Marshal(&response)
	if err != nil {
		log.Println("Error trying to write json response ")
	}
	fmt.Fprintf(w, string(jsonresponse))

}
func getKallas(sunrise time.Time, sunset time.Time, date time.Time) kaalamType {

	rahuKaalStartAndEndTime := getRahuKaal(sunrise, sunset, date)
	return rahuKaalStartAndEndTime
}
