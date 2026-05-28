package main

import (
	"encoding/xml"
	"os"
	"strconv"
	"time"
)

type xmlRoot struct {
	Stations []xmlStation `xml:"station"`
}

type xmlCoordinates struct {
	Latitude  float64 `xml:"lat,attr"`
	Longitude float64 `xml:"lon,attr"`
	Altitude  int     `xml:"altitude,attr"`
}

type xmlHardware struct {
	Model  string `xml:"model,attr"`
	Vendor string `xml:"vendor,attr"`
	Since  string `xml:"since,attr"`
}

type xmlWind struct {
	Speed     float64 `xml:"speed"`
	Direction int     `xml:"direction"`
}

type xmlAirQuality struct {
	Pollutants []xmlPollutant `xml:"pollutant"`
}

type xmlPollutant struct {
	Name  string `xml:"name,attr"`
	Value string `xml:",chardata"`
}

type xmlMeasure struct {
	Type  string `xml:"type,attr"`
	Value string `xml:",chardata"`
}

type xmlObservation struct {
	At         string        `xml:"at,attr"`
	Sky        string        `xml:"sky,attr"`
	Measures   []xmlMeasure  `xml:"measure"`
	Wind       xmlWind       `xml:"wind"`
	AirQuality xmlAirQuality `xml:"air_quality"`
	Note       *string       `xml:"note"`
}

type xmlStation struct {
	ID           string           `xml:"id,attr"`
	Name         string           `xml:"name"`
	Country      string           `xml:"country,attr"`
	Coordinates  xmlCoordinates   `xml:"coordinates"`
	Hardware     xmlHardware      `xml:"hardware"`
	Observations []xmlObservation `xml:"observations>observation"`
}

func LoadFromXML(path string) ([]Station, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var root xmlRoot

	err = xml.Unmarshal(data, &root)
	if err != nil {
		return nil, err
	}

	stations := make([]Station, 0, len(root.Stations))

	for _, s := range root.Stations {
		station := convertXML(s)

		stations = append(stations, station)
	}

	return stations, nil
}

func convertXML(station xmlStation) Station {
	installedAt, err := time.Parse(time.DateOnly, station.Hardware.Since)
	if err != nil {
	}

	observations := make([]Observation, 0, len(station.Observations))

	for _, obs := range station.Observations {
		newObservation := convertXMLObservation(obs)
		if err != nil {
		}

		observations = append(observations, newObservation)
	}

	return Station{
		ID:       station.ID,
		Name:     station.Name,
		Country:  station.Country,
		Altitude: station.Coordinates.Altitude,

		Coordinates: Coordinates{
			Latitude:  station.Coordinates.Latitude,
			Longitude: station.Coordinates.Longitude,
		},

		Sensor: Sensor{
			Model:        station.Hardware.Model,
			Manufacturer: station.Hardware.Vendor,
			InstalledAt:  installedAt,
		},

		Observations: observations,
	}
}

func convertXMLObservation(obs xmlObservation) (out Observation) {
	timestampParsed, err := time.Parse(time.RFC3339, obs.At)
	if err != nil {
	}

	for _, m := range obs.Measures {
		v, _ := strconv.ParseFloat(m.Value, 64)
		switch m.Type {
		case "temperature":
			out.Temperature = v
		case "humidity":
			out.Humidity = int(v)
		case "pressure":
			out.Pressure = v
		case "precipitation":
			out.Precipitation = v
		}
	}

	for _, m := range obs.AirQuality.Pollutants {
		v, _ := strconv.ParseFloat(m.Value, 64)
		switch m.Name {
		case "PM2.5":
			out.AirQuality.PM25 = v
		case "PM10":
			out.AirQuality.PM10 = v
		case "NO2":
			out.AirQuality.NO2 = v
		}
	}

	out.Timestamp = timestampParsed
	out.Wind.Speed = obs.Wind.Speed
	out.Wind.Direction = obs.Wind.Direction
	out.Conditions = obs.Sky
	out.Notes = obs.Note

	return out
}
