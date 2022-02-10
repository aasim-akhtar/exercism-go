// Package weather provides this and that.
package weather
// CurrentCondition stores info about the current weather forecast of city.
var CurrentCondition string
// CurrentLocation stores the location of city.
var CurrentLocation string
// Forecast return the current weather condition of the city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
