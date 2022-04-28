package main

import "fmt"

// Task 1
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return successRate * float64(productionRate) / 100
}

// Task 2
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(float64(productionRate/60) * successRate / 100)
}

// Task 3
func CalculateCost(carsCount int) uint {
	return uint((carsCount / 10 * 95000) + (carsCount % 10 * 10000))
}

func main() {
	fmt.Println(CalculateWorkingCarsPerHour(1920, 70))
	fmt.Println(CalculateWorkingCarsPerMinute(1920, 70))
	fmt.Println(CalculateCost(37))
}
