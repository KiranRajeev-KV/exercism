// Package weather provides tools to get forcasts.
package weather



var (
    // CurrentCondition tells current condition.
	CurrentCondition string
    // CurrentLocation tells current location.
	CurrentLocation  string
)

// Forecast returns forcast and the condition of the city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
