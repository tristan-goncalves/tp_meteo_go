package main

import "time"

type Station struct {
	ID           string
	Name         string
	Country      string
	Altitude     int
	Coordinates  Coordinates
	Sensor       Sensor
	Observations []Observation
}

type Coordinates struct {
	Latitude  float64
	Longitude float64
}

type Sensor struct {
	Model        string
	Manufacturer string
	InstalledAt  time.Time
}

type Observation struct {
	Timestamp     time.Time
	Temperature   float64
	Humidity      int
	Pressure      float64
	Wind          Wind
	AirQuality    AirQuality
	Conditions    string
	Precipitation float64
	Notes         *string
}

type Wind struct {
	Speed     float64
	Direction int
}

type AirQuality struct {
	PM25 float64
	PM10 float64
	NO2  float64
}
