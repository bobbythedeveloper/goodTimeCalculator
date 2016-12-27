package hello

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
)

func init() {
	http.HandleFunc("/getKaalam", getKaalam)
}

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Hello, world!")
}

func getKaalam(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	client := urlfetch.Client(ctx)

	latitude := float64(36.7201600)
	longitude := float64(-4.4203400)

	sunrise, sunset := getSunriseAndSunset(latitude, longitude, client)
	guliKaala, yamaGantaKaala, rahuKaala := getKallas(sunrise, sunset, time.Now())

	log.Println("------------" + sunrise)
	log.Println("------------" + sunset)

	log.Println(guliKaala)
	log.Println(yamaGantaKaala)
	log.Println(rahuKaala)

}
func getKallas(sunrise string, sunset string, date time.Time) (string, string, string) {
	sunriseTime, err := time.Parse("3:04:05 PM", sunrise)
	if err != nil {
		log.Println("error parsing time sunriseTime")
		log.Fatal(err)
	}
	sunsetTime, err := time.Parse("3:04:05 PM", sunset)
	if err != nil {
		log.Println("error parsing time sunriseTime")
		log.Fatal(err)
	}
	log.Println(sunsetTime.String())
	log.Println(sunriseTime.String())
	parts := (sunsetTime.Unix() - sunriseTime.Unix()) / 8
	log.Println("first part -- before" + sunriseTime.String())

	firstpart := sunriseTime.Add(time.Duration(parts * 1000000000 * getRahuPosition(date)))
	log.Println("first part --" + firstpart.String())
	log.Println(date.Weekday())
	return "", "", ""
}
func getRahuPosition(datetime time.Time) int64 {
	switch datetime.Weekday() {
	case time.Monday:
		return 2
	case time.Tuesday:
		return 7
	case time.Wednesday:
		return 5
	case time.Thursday:
		return 6
	case time.Friday:
		return 4
	case time.Saturday:
		return 3
	case time.Sunday:
		return 8
	}

	return 1
}

func getSunriseAndSunset(lat float64, longt float64, client *http.Client) (string, string) {
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
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Issue calling backend webservice")
		log.Fatal(err)

	}
	body1, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Cannot read http body")
	}
	var backendResponse *sunRiseSunSet
	err = json.Unmarshal(body1, &backendResponse)
	if err != nil {
		log.Println("Error understanding backend json")
		log.Fatal(err)
	}
	return backendResponse.Results.Sunrise, backendResponse.Results.Sunset
}

type sunRiseSunSet struct {
	Results struct {
		Sunrise                   string `json:"sunrise"`
		Sunset                    string `json:"sunset"`
		SolarNoon                 string `json:"solar_noon"`
		DayLength                 string `json:"day_length"`
		CivilTwilightBegin        string `json:"civil_twilight_begin"`
		CivilTwilightEnd          string `json:"civil_twilight_end"`
		NauticalTwilightBegin     string `json:"nautical_twilight_begin"`
		NauticalTwilightEnd       string `json:"nautical_twilight_end"`
		AstronomicalTwilightBegin string `json:"astronomical_twilight_begin"`
		AstronomicalTwilightEnd   string `json:"astronomical_twilight_end"`
	} `json:"results"`
	Status string `json:"status"`
}
