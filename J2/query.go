package main

func FilterByCountry(stations []Station, iso string) []Station {
	var result []Station

	for _, station := range stations {
		if station.Country == iso {
			result = append(result, station)
		}
	}
	return result
}

func AvgTemperature(s Station) float64 {
	if len(s.Observations) == 0 {
		return 0
	}

	var sum float64

	for _, obs := range s.Observations {
		sum += obs.Temperature
	}
	return sum / float64(len(s.Observations))
}

func MaxWindGust(stations []Station) (Station, float64) {
	var maxStation Station
	var maxWind float64

	for _, station := range stations {
		for _, obs := range station.Observations {
			if obs.Wind.Speed > maxWind {
				maxWind = obs.Wind.Speed
				maxStation = station
			}
		}
	}
	return maxStation, maxWind
}

func CountByCountry(stations []Station) map[string]int {
	counts := make(map[string]int)

	for _, station := range stations {
		counts[station.Country]++
	}
	return counts
}
