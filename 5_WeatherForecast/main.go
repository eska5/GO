// Package main provides CurrentCondition, CurrentLocation and
// has a function called Forecast which forecasts the weather in the city.
package main

import "fmt"

// CurrentCondition represents a certain weather condition in string.
var CurrentCondition string

// CurrentLocation represents a certain location in string.
var CurrentLocation string

// Forecast returns string of combined CurrentLocation and CurrentCondition based on variables city and condition in string format.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}

// main prints Forecast function result in console using fmt.
func main() {
	fmt.Println(Forecast("Gdansk", "foggy"))
}
