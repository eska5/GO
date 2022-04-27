package main

import "fmt"

// Task 1
const OvenTime = 40

// Task 2
func RemainingOvenTime(actualMinutesInOven int) int {
	return OvenTime - actualMinutesInOven
}

// Task 3
func PreparationTime(numberOfLayers int) int {
	return numberOfLayers * 2
}

// Task 4
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	return 2*numberOfLayers + actualMinutesInOven
}

func main() {
	fmt.Println(OvenTime)
	fmt.Println(RemainingOvenTime(20))
	fmt.Println(PreparationTime(4))
	fmt.Println(ElapsedTime(4, 20))
}
