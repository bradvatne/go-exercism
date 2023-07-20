
//Package weather contains forecast 
package weather

//CurrentCondition: tells us about the current weather for the location
var CurrentCondition string
//CurrentLocation: is the current location for the forecast
var CurrentLocation string

//Forecast : is a function that returns a string with the weather forecast details
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
