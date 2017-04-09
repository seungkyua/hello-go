/*
 *  http://howistart.org/posts/go/1/index.html
 *  https://home.openweathermap.org/users/sign_up
 *         id : hello-go / hello-go
 *        key : 92740b1767b3711a012f06ca7a9f0013
 *  http://api.openweathermap.org/data/2.5/weather?APPID=92740b1767b3711a012f06ca7a9f0013&q=seoul

{
  "coord": {
    "lon": 126.98,
    "lat": 37.57
  },
  "weather": [
    {
      "id": 801,
      "main": "Clouds",
      "description": "few clouds",
      "icon": "02d"
    }
  ],
  "base": "stations",
  "main": {
    "temp": 289.72,
    "pressure": 1015,
    "humidity": 36,
    "temp_min": 287.15,
    "temp_max": 291.15
  },
  "visibility": 10000,
  "wind": {
    "speed": 2.1,
    "deg": 110
  },
  "clouds": {
    "all": 20
  },
  "dt": 1491729600,
  "sys": {
    "type": 1,
    "id": 8514,
    "message": 0.007,
    "country": "KR",
    "sunrise": 1491685536,
    "sunset": 1491732131
  },
  "id": 1835848,
  "name": "Seoul",
  "cod": 200
}



 *  https://www.wunderground.com/weather/api
           id : skanddh@gmail.com / hello-go
          key : a8440a8f8703d4d8
 * http://api.wunderground.com/api/a8440a8f8703d4d8/conditions/q/seoul.json

{
  response: {
    version: "0.1",
    termsofService: "http://www.wunderground.com/weather/api/d/terms.html",
    features: {
      conditions: 1
    }
  },
  current_observation: {
    image: {
      url: "http://icons.wxug.com/graphics/wu2/logo_130x80.png",
      title: "Weather Underground",
      link: "http://www.wunderground.com"
    },
    display_location: {
      full: "Seoul, South Korea",
      city: "Seoul",
      state: "11",
      state_name: "South Korea",
      country: "KO",
      country_iso3166: "KR",
      zip: "00000",
      magic: "1",
      wmo: "47108",
      latitude: "37.56999969",
      longitude: "126.97000122",
      elevation: "86.0"
    },
    observation_location: {
      full: "Seoul, ",
      city: "Seoul",
      state: "",
      country: "Korea, South",
      country_iso3166: "KR",
      latitude: "37.57138824",
      longitude: "126.96583557",
      elevation: "282 ft"
    },
    estimated: {

    },
    station_id: "RKSL",
    observation_time: "Last Updated on April 9, 6:00 PM KST",
    observation_time_rfc822: "Sun, 09 Apr 2017 18:00:00 +0900",
    observation_epoch: "1491728400",
    local_time_rfc822: "Sun, 09 Apr 2017 19:56:17 +0900",
    local_epoch: "1491735377",
    local_tz_short: "KST",
    local_tz_long: "Asia/Seoul",
    local_tz_offset: "+0900",
    weather: "Clear",
    temperature_string: "63 F (17 C)",
    temp_f: 63,
    temp_c: 17,
    relative_humidity: "28%",
    wind_string: "From the East at 6 MPH",
    wind_dir: "East",
    wind_degrees: 90,
    wind_mph: 6,
    wind_gust_mph: 0,
    wind_kph: 9,
    wind_gust_kph: 0,
    pressure_mb: "1014",
    pressure_in: "29.93",
    pressure_trend: "-",
    dewpoint_string: "38 F (4 C)",
    dewpoint_f: 38,
    dewpoint_c: 4,
    heat_index_string: "NA",
    heat_index_f: "NA",
    heat_index_c: "NA",
    windchill_string: "NA",
    windchill_f: "NA",
    windchill_c: "NA",
    feelslike_string: "63 F (17 C)",
    feelslike_f: "63",
    feelslike_c: "17",
    visibility_mi: "12.0",
    visibility_km: "20.0",
    solarradiation: "--",
    UV: "-1",
    precip_1hr_string: " in ( mm)",
    precip_1hr_in: "",
    precip_1hr_metric: "--",
    precip_today_string: " in ( mm)",
    precip_today_in: "",
    precip_today_metric: "--",
    icon: "clear",
    icon_url: "http://icons.wxug.com/i/c/k/nt_clear.gif",
    forecast_url: "http://www.wunderground.com/global/stations/47108.html",
    history_url: "http://www.wunderground.com/history/airport/RKSL/2017/4/9/DailyHistory.html",
    ob_url: "http://www.wunderground.com/cgi-bin/findweather/getForecast?query=37.57138824,126.96583557",
    nowcast: ""
  }
}current_observation.display_location+-Viewsourceoptions


 */

package main

import (
	"net/http"
	"encoding/json"
	"strings"
	"log"
	"time"
)

/*
type weatherData struct {
	Name string `json:"name"`
	Main struct {
			 Kelvin float64 `json:"temp"`
		 } `json:"main"`
}


func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/weather/", func(w http.ResponseWriter, r *http.Request) {
		city := strings.SplitN(r.URL.Path, "/", 3)[2]

		data, err := query(city)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(data)
	})
	http.ListenAndServe(":9090", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello!"))
}

func query(city string) (weatherData, error) {
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?APPID=92740b1767b3711a012f06ca7a9f0013&q=" + city)
	if err != nil {
		return weatherData{}, err
	}

	defer resp.Body.Close()

	var d weatherData

	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return weatherData{}, err
	}

	return d, nil
}
*/

func main() {
	mw := multiWeatherProvider{
		openWeatherMap{},
		weatherUnderground{apiKey: "a8440a8f8703d4d8"},
	}

	http.HandleFunc("/weather/", func(w http.ResponseWriter, r *http.Request) {
		begin := time.Now()
		city := strings.SplitN(r.URL.Path, "/", 3)[2]

		temp, err := mw.temperature(city)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"city": city,
			"tmep": temp,
			"took": time.Since(begin).String(),
		})
	})
	http.ListenAndServe(":9090", nil)
}

type weatherProvider interface {
	temperature(city string) (float64, error)
}

type openWeatherMap struct{}

func (w openWeatherMap) temperature(city string) (float64, error) {
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?APPID=92740b1767b3711a012f06ca7a9f0013&q=" + city)
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	var d struct {
		Main struct {
				 Kelvin float64 `json:"temp"`
			 } `json:"main"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return 0, err
	}

	log.Printf("openWeatherMap: %s: %.2f", city, d.Main.Kelvin)

	return d.Main.Kelvin, nil
}

type weatherUnderground struct {
	apiKey string
}

func (w weatherUnderground) temperature(city string) (float64, error) {
	resp, err := http.Get("http://api.wunderground.com/api/" + w.apiKey + "/conditions/q/" + city + ".json")
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	var d struct {
		Observation struct {
						Celsius float64 `json:"temp_c"`
					} `json:"current_observation"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return 0, err
	}

	kelvin := d.Observation.Celsius + 273.15
	log.Printf("weatherUnderground: %s: %.2f", city, kelvin)
	return kelvin, nil
}

type multiWeatherProvider []weatherProvider

func (w multiWeatherProvider) temperature(city string) (float64, error) {
	sum := 0.0

	for _, provider := range w {
		k, err := provider.temperature(city)
		if err != nil {
			return 0, err
		}

		sum += k
	}

	return sum / float64(len(w)), nil
}

/*
func temperature(city string, providers ...weatherProvider) (float64, error) {
	sum := 0.0

	for _, provider := range providers {
		k, err := provider.temperature(city)
		if err != nil {
			return 0, err
		}

		sum += k
	}

	return sum / float64(len(providers)), nil
}
*/