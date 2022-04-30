package main

import "fmt"

// Task 1
func NeedsLicense(kind string) bool {
	if kind == "truck" || kind == "car" {
		return true
	} else {
		return false
	}
}

// Task 2
func ChooseVehicle(option1, option2 string) string {
	if option1 > option2 {
		return option2 + " is clearly the better choice."
	} else {
		return option1 + " is clearly the better choice."
	}
}

// Task 3
func CalculateResellPrice(originalPrice, age float64) float64 {
	if age <= 3 {
		return originalPrice * 0.8
	} else if age >= 10 {
		return originalPrice * 0.5
	} else {
		return originalPrice * 0.7
	}
}
func main() {
	fmt.Println(NeedsLicense("truck"))
	fmt.Println(ChooseVehicle("ferrari", "lamborghini"))
	fmt.Println(CalculateResellPrice(2100000.0, 4))
}
