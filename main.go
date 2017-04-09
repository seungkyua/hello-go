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

 */

package main

import (
	"net/http"
	"encoding/json"
	"strings"
)

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