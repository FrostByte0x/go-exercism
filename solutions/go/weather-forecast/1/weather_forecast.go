// Package weather will display the weather at a given location.
package weather

var (
	// CurrentCondition is the weather condition.
	CurrentCondition string
	// CurrentLocation is the location where the condition is occuring.
	CurrentLocation string
)

// Forecast will output the weather condition at a given city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
