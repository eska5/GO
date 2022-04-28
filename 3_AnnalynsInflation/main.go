package main

import "fmt"

// Task 1
func CanFastAttack(knightIsAwake bool) bool {
	return !knightIsAwake
}

// Task 2
func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
	if knightIsAwake || archerIsAwake || prisonerIsAwake {
		return true
	} else {
		return false
	}
}

// Task 3
func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
	if !archerIsAwake && prisonerIsAwake {
		return true
	} else {
		return false
	}
}

// Task 4
func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
	if petDogIsPresent && !archerIsAwake {
		return true
	} else if !petDogIsPresent && prisonerIsAwake && !archerIsAwake && !knightIsAwake {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(CanFastAttack(true))
	fmt.Println(CanSpy(true, false, true))
	fmt.Println(CanSignalPrisoner(true, false))
	fmt.Println(CanFreePrisoner(false, true, false, false))
}
