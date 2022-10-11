// Package weather : Since Goblinocus takes its weather forecast very seriously, We create this function to forecast the weather condition of various cities in Goblinocus.
package weather

// CurrentCondition is the current weather condition of the city in Goblinocus.
var CurrentCondition string

// CurrentLocation is the current location of the weather forecast.
var CurrentLocation string

// Forecast is the function to forecast the weather condition of the cities in Goblinocus.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
