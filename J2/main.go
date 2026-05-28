package main

import "fmt"

func main() {

	jsonStations, _ := LoadFromJSON("weather_data.json")
	xmlStations, _ := LoadFromXML("weather_data.xml")

	jsonObsCount := 0

	for _, station := range jsonStations {
		jsonObsCount += len(station.Observations)
	}

	xmlObsCount := 0

	for _, station := range xmlStations {
		xmlObsCount += len(station.Observations)
	}

	fmt.Printf("JSON : %d stations, %d observations\n", len(jsonStations), jsonObsCount)

	fmt.Printf("XML : %d stations, %d observations\n", len(xmlStations), xmlObsCount)

	if len(jsonStations) == len(xmlStations) &&
		jsonObsCount == xmlObsCount {
		fmt.Println("Cohérence : OK")

	} else {

		fmt.Println("Cohérence : ERREUR")
	}
	fmt.Println()

	maxStation, maxWind := MaxWindGust(jsonStations)

	fmt.Printf("Station la plus ventée : %s (%.1f km/h)\n", maxStation.ID, maxWind)
	fmt.Println()

	var bordeaux Station

	for _, station := range jsonStations {

		if station.ID == "FR-BOR-001" {
			bordeaux = station
		}
	}

	avgTemp := AvgTemperature(bordeaux)

	fmt.Printf("Température moyenne %s : %.1f °C\n", bordeaux.Name, avgTemp)
	fmt.Println()

	counts := CountByCountry(jsonStations)

	fmt.Println("Stations par pays :")
	fmt.Println(counts)
}
