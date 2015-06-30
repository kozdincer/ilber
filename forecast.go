package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

var (
	weatherURL = "http://api.openweathermap.org/data/2.5/weather?q=%s&units=metric"
)

const (
	cityNum = 81
	burdur  = 15
)

func init() {
	register("/hava", forecast)
}

// openweathermap response
type Forecast struct {
	City    string `json:"name"`
	Weather []struct {
		ID          int    `json:"id"`
		Status      string `json:"main"`
		Description string
	}
	Temperature struct {
		Celsius float64 `json:"temp"`
	} `json:"main"`
}

func (f Forecast) String() string {
	var icon string
	now := time.Now()

	if len(f.Weather) == 0 {
		return ""
	}

	switch f.Weather[0].Status {
	case "Clear":
		if 6 < now.Hour() && now.Hour() < 18 { // for istanbul
			icon = "☀"
		} else {
			icon = "☽"
		}
	case "Clouds":
		icon = "☁"
	case "Rain":
		icon = "☔"
	case "Fog":
		icon = "▒"
	case "Mist":
		icon = "░"
	case "Haze":
		icon = "░"
	case "Snow":
		icon = "❄"
	case "Thunderstorm":
		icon = "⚡"
	default:
		icon = ""
	}

	return fmt.Sprintf("%v %v %.1f °C (%v)", icon, f.City, f.Temperature.Celsius, f.Weather[0].Description)
}

func forecast(locations ...string) string {
	var location string
	if locations == nil {
		location = "Istanbul"
	} else {
		location = strings.Join(locations, " ")
	}

	url := fmt.Sprintf(weatherURL, location)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("weather error: %v\n", err)
		return ""
	}
	defer resp.Body.Close()

	var forecast Forecast
	if err := json.NewDecoder(resp.Body).Decode(&forecast); err != nil {
		log.Printf("decode error: %v\n", err)
		return ""
	}

	if forecast.String() == "" {
		// burdur easter-egg.
		if rand.Intn(cityNum) == burdur {
			return fmt.Sprintf("%v bulunamadi ama Burdur'da hava cok guzel.", location)
		}

		return fmt.Sprintf("%v bulunamadi.", location)

	}

	return forecast.String()
}
