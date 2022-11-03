package main

import (
	"encoding/json"
	"fmt"
	"github.com/hablullah/go-prayer"
	"net/http"
	"strconv"
	"time"
)

// data adhan
type dataadhan struct {
	Day     string
	Fajr    string
	Sunrise string
	Zuhr    string
	Asr     string
	Maghrib string
	Isha    string
}

// dataadhaninmonth
type dataadhaninmonth struct {
	Timings dataadhan
	Date    map[string]interface{}
}

func adhan(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		latitude, _ := strconv.ParseFloat(r.FormValue("latitude"), 32)
		longitude, _ := strconv.ParseFloat(r.FormValue("longitude"), 32)
		month, _ := strconv.Atoi(r.FormValue("month"))
		year, _ := strconv.Atoi(r.FormValue("year"))

		var err error
		now := time.Now()
		currentLocation := now.Location()

		//config
		cfg := prayer.Config{
			Latitude:          latitude,
			Longitude:         longitude,
			Elevation:         50,
			CalculationMethod: prayer.Kemenag,
			AsrConvention:     prayer.Shafii,
			PreciseToSeconds:  false,
			TimeCorrections: prayer.TimeCorrections{
				Fajr:    2 * time.Minute,
				Sunrise: -time.Minute,
				Zuhr:    2 * time.Minute,
				Asr:     time.Minute,
				Maghrib: time.Minute,
				Isha:    time.Minute,
			},
		}

		//get first day and last day of the month and the year
		firstOfMonth := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, currentLocation)
		startDate := firstOfMonth.Day()
		endDate := firstOfMonth.AddDate(0, 1, -1).Day()

		//zone
		zone := time.FixedZone("WIB", 7*3600)

		//tampungan data
		var dataadhansemua []interface{}

		for i := startDate; i <= endDate; i++ {
			date := time.Date(year, time.Month(month), i, 0, 0, 0, 0, zone)

			result, _ := prayer.Calculate(cfg, date)
			/*entry := map[string]interface{}{
				"timings": dataadhan{strconv.Itoa(i), result.Fajr.Format("15:04"), result.Sunrise.Format("15:04"), result.Zuhr.Format("15:04"), result.Asr.Format("15:04"), result.Maghrib.Format("15:04"), result.Isha.Format("15:04")},
				"date":    map[string]interface{}{"readable": date.String(), "timestamp": date.Unix()},
			}*/
			dataelemen := dataadhan{strconv.Itoa(i), result.Fajr.Format("15:04"), result.Sunrise.Format("15:04"), result.Zuhr.Format("15:04"), result.Asr.Format("15:04"), result.Maghrib.Format("15:04"), result.Isha.Format("15:04")}
			entry := dataadhaninmonth{
				dataelemen,
				map[string]interface{}{"readable": date.String(), "timestamp": date.Unix()}}
			dataadhansemua = append(dataadhansemua, entry)
		}

		responseOutput := map[string]interface{}{
			"code":      200,
			"status":    "OK",
			"latitude":  latitude,
			"longitude": longitude,
			"data":      dataadhansemua,
		}

		result2, err2 := json.Marshal(responseOutput) // data yang di-encode ke JSON

		if err2 != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result2) //r.Write() digunakan untuk mendaftarkan data sebagai response
		return
	}

	http.Error(w, "", http.StatusBadRequest)

}

func main() {
	fmt.Println("Service Adhan")
	http.HandleFunc("/adhan", adhan)

	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
