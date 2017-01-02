package hello

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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
	//latitude := float64(39.948184)
	//longitude := float64(-74.902575)

	sunrise, sunset := getSunriseAndSunset(latitude, longitude, date, client)
	//rahuKaala := getKallas(sunrise, sunset, time.Now())
	rahuKaala := getKallas(sunrise, sunset, date)

	response := ResponseJSON{}
	//response.RahuKaalEndTime = rahuKaala.endTime.Format("3:04:05 PM")
	//response.RahuKaalStartTime = rahuKaala.startTime.Format("3:04:05 PM")

	response.RahuKaalEndTime = rahuKaala.endTime
	response.RahuKaalStartTime = rahuKaala.startTime

	jsonresponse, err := json.Marshal(&response)
	if err != nil {
		log.Println("Error trying to write json response ")
	}
	os.Stdout.Write(jsonresponse)
	fmt.Fprintf(w, string(jsonresponse))

}
func getKallas(sunrise time.Time, sunset time.Time, date time.Time) kaalamType {
	/*	sunriseTime, err := time.Parse("3:04:05 PM", sunrise)
		if err != nil {
			log.Println("error parsing time sunriseTime")
			log.Fatal(err)
		}
		sunsetTime, err := time.Parse("3:04:05 PM", sunset)
		if err != nil {
			log.Println("error parsing time sunriseTime")
			log.Fatal(err)
		}*/
	rahuKaalStartAndEndTime := getRahuKaal(sunrise, sunset, date)
	return rahuKaalStartAndEndTime
}
func getRahuKaal(sunrise time.Time, sunset time.Time, date time.Time) kaalamType {
	rahuKaalStartAndEndTime := kaalamType{}
	parts := (sunset.Unix() - sunrise.Unix()) / 8

	firstpart := sunrise.Add(time.Duration(parts * 1000000000 * getRahuPosition(date)))
	//log.Println("first part --" + strconv.ParseInt(parts, 10, 64))
	rahuKaalStartAndEndTime.startTime = firstpart
	rahuKaalStartAndEndTime.endTime = rahuKaalStartAndEndTime.startTime.Add(time.Duration(parts * 1000000000))
	return rahuKaalStartAndEndTime
}

func getSunriseAndSunset(lat float64, longt float64, date time.Time, client *http.Client) (time.Time, time.Time) {
	latstr := fmt.Sprintf("%f", lat)
	longstr := fmt.Sprintf("%f", longt)
	dateString := fmt.Sprintf("%d-%d-%d", date.Year(), date.Month(), date.Day())
	//http://api.sunrise-sunset.org/json?lat=36.7201600&lng=-4.4203400&date=2016-12-26
	req, err := http.NewRequest("GET", "http://api.sunrise-sunset.org/json", nil)
	if err != nil {
		log.Print(err)
	}
	q := req.URL.Query()
	q.Add("lat", latstr)
	q.Add("lng", longstr)
	q.Add("date", dateString)
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
	sunriseDateTime, err := time.Parse("3:04:05 PM 2006-1-2", backendResponse.Results.Sunrise+" "+dateString)
	if err != nil {
		log.Println("Unable parse sun rise time from backend")
		log.Println(err)
	}
	sunsetDateTime, err := time.Parse("3:04:05 PM 2006-1-2", backendResponse.Results.Sunset+" "+dateString)
	if err != nil {
		log.Println("Unable parse sun rise time from backend")
		log.Println(err)
	}
	//loc, err := time.LoadLocation("America/New_York")
	//if err != nil {
	//		panic(err)
	//	}
	//sunriseDateTime.In(loc)
	//log.Println(sunriseDateTime.In(loc))
	//	return backendResponse.Results.Sunrise, backendResponse.Results.Sunset
	return sunriseDateTime, sunsetDateTime

}
