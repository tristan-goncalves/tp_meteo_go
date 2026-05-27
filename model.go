package main

import "time"

type Station struct {
	ID           string
	Name         string
	Country      string
	Altitude     int
	Location     Location
	Device       Device
	Observations []Observation
}

type Location struct {
	Latitude  float64
	Longitude float64
}

type Device struct {
	Type         string
	Manufacturer string
	InstalledOn  time.Time
}

type Observation struct {
	Timestamp     time.Time
	Temperature   float64 // °C
	Humidity      int     // %
	Pressure      float64 // hPa
	Wind          Wind
	Precipitation float64 // mm
	AirQuality    AirQuality
	Conditions    string
	Notes         *string
}

type Wind struct {
	Speed     float64 // km/h
	Direction int     // degrees
}

type AirQuality struct {
	PM25 float64
	PM10 float64
	NO2  float64
}
