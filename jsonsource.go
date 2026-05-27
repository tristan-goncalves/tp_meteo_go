package main

type jsonRoot struct {
	Stations []jsonStation `json:"stations"`
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

type jsonLocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type jsonDevice struct {
	Type         string `json:"type"`
	Manufacturer string `json:"manufacturer"`
	InstalledOn  string `json:"installed_on"`
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

type jsonWind struct {
	Speed     float64 `json:"speed_kmh"`
	Direction int     `json:"direction_deg"`
}

type jsonAirQuality struct {
	PM25 float64 `json:"pm25"`
	PM10 float64 `json:"pm10"`
	NO2  float64 `json:"no2"`
}
