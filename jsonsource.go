package main

import (
	"encoding/json"
	"os"
	"strings"
	"time"
)

type jsonRoot struct {
	Stations []jsonStation `json:"stations"`
}

type jsonLocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type jsonDevice struct {
	Type         string `json:"type"`
	Manufacturer string `json:"manufacturer"`
	InstalledOn  string `json:"installed_on"`
}

type jsonWind struct {
	Speed     float64 `json:"speed_kmh"`
	Direction int     `json:"direction_deg"`
}

type jsonAirQuality struct {
	PM25 float64 `json:"pm25"`
	PM10 float64 `json:"pm10"`
	NO2  float64 `json:"no2"`
}

type jsonObservation struct {
	Timestamp     string         `json:"timestamp"`
	Temperature   float64        `json:"temperature_celsius"`
	Humidity      int            `json:"humidity_percent"`
	Pressure      float64        `json:"pressure_hpa"`
	Wind          jsonWind       `json:"wind"`
	Precipitation float64        `json:"precipitation_mm"`
	AirQuality    jsonAirQuality `json:"air_quality"`
	Conditions    string         `json:"conditions"`
	Notes         *string        `json:"notes"`
}

type jsonStation struct {
	ID           string            `json:"id"`
	Name         string            `json:"name"`
	Country      string            `json:"country"`
	Altitude     int               `json:"altitude_m"`
	Location     jsonLocation      `json:"location"`
	Device       jsonDevice        `json:"device"`
	Observations []jsonObservation `json:"observations"`
}

var countryListISO = map[string]string{
	"france":    "FR",
	"espagne":   "ES",
	"portugal":  "PT",
	"italie":    "IT",
	"allemagne": "DE",
	"belgique":  "BE",
	"pays-bas":  "NL",
	"autriche":  "AT",
	"suisse":    "CH",
	"danemark":  "DK",
	"suède":     "SE",
	"norvège":   "NO",
	"pologne":   "PL",
	"tchéquie":  "CZ",
}

func main() {
	LoadFromJSON("weather_data.json")
}

func countryToISO(country string) string {
	normalized := strings.ToLower(country)
	countryISO := countryListISO[normalized]

	return countryISO
}

func LoadFromJSON(path string) ([]Station, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var root jsonRoot

	err = json.Unmarshal(data, &root)
	if err != nil {
		return nil, err
	}

	stations := make([]Station, 0, len(root.Stations))

	for _, s := range root.Stations {
		stations = append(stations, convertJSON(s))
	}

	println(len(stations))
	println(len(stations[0].Observations))

	return stations, nil
}

func convertJSON(station jsonStation) Station {
	// Sensor
	installedOnParsed, err := time.Parse(time.DateOnly, station.Device.InstalledOn)
	if err != nil {
	}

	// Observations
	var observations []Observation

	for _, obs := range station.Observations {
		timestampParsed, err := time.Parse(time.RFC3339, obs.Timestamp)
		if err != nil {
		}

		newObservation := Observation{
			Timestamp:     timestampParsed,
			Temperature:   obs.Temperature,
			Humidity:      obs.Humidity,
			Pressure:      obs.Pressure,
			Conditions:    obs.Conditions,
			Precipitation: obs.Precipitation,
			Notes:         obs.Notes,
			Wind: Wind{
				Speed:     obs.Wind.Speed,
				Direction: obs.Wind.Direction,
			},
			AirQuality: AirQuality{
				PM25: obs.AirQuality.PM25,
				PM10: obs.AirQuality.PM10,
				NO2:  obs.AirQuality.NO2,
			},
		}

		observations = append(observations, newObservation)
	}

	return Station{
		ID:       station.ID,
		Name:     station.Name,
		Country:  countryToISO(station.Country),
		Altitude: station.Altitude,
		Coordinates: Coordinates{
			Latitude:  station.Location.Latitude,
			Longitude: station.Location.Longitude,
		},
		Sensor: Sensor{
			Model:        station.Device.Type,
			Manufacturer: station.Device.Manufacturer,
			InstalledAt:  installedOnParsed,
		},
		Observations: observations,
	}
}
