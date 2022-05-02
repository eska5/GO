package main

import "fmt"

// Task 1
func TotalBirdCount(birdsPerDay []int) int {
	sum := 0
	for i := 0; i < len(birdsPerDay); i++ {
		sum += birdsPerDay[i]
	}
	return sum
}

// Task 2
func BirdsInWeek(birdsPerDay []int, week int) int {
	sum := 0
	for i := (week - 1) * 7; i < week*7; i++ {
		sum += birdsPerDay[i]
	}
	return sum
}

// Task 3
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i++ {
		if i%2 == 0 {
			birdsPerDay[i]++
		}
	}
	return birdsPerDay
}

//main
func main() {
	fmt.Println(TotalBirdCount([]int{2, 5, 0, 7, 4, 1}))
	fmt.Println(BirdsInWeek([]int{2, 5, 0, 7, 4, 1, 3, 0, 2, 5, 0, 1, 3, 1}, 2))
	fmt.Println(FixBirdCountLog([]int{2, 5, 0, 7, 4, 1}))
}
