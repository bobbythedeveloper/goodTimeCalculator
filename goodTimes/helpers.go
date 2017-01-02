package hello

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func getRahuPosition(datetime time.Time) int64 {
	switch datetime.Weekday() {
	case time.Monday:
		return 0
	case time.Tuesday:
		return 5
	case time.Wednesday:
		return 3
	case time.Thursday:
		return 4
	case time.Friday:
		return 2
	case time.Saturday:
		return 1
	case time.Sunday:
		return 6
	}

	return 1
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

	return sunriseDateTime, sunsetDateTime

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
