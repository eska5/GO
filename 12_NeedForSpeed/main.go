package main

import "fmt"

// Task 1
type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

// Task 1
func NewCar(speed, batteryDrain int) Car {
	return Car{
		battery:      100,
		batteryDrain: batteryDrain,
		speed:        speed,
		distance:     0, // either , or } cannot do \n }
	}
}

// Task 2
type Track struct {
	distance int
}

// Task 2
func NewTrack(distance int) Track {
	return Track{distance: distance}
}

// Task 3
func Drive(car Car) Car {
	if car.battery-car.batteryDrain > 0 {
		car.battery -= car.batteryDrain
		car.distance += car.speed
		return car
	} else {
		return car
	}
}

// Task 4
func CanFinish(car Car, track Track) bool {
	laps := track.distance / car.speed
	if track.distance%car.speed != 0 {
		laps += 1
	}
	if laps*car.batteryDrain > car.battery {
		return false
	} else {
		return true
	}
}

//main
func main() {
	fmt.Println(NewCar(10, 12))
	fmt.Println(NewTrack(200))
	fmt.Println(Drive(NewCar(10, 12)))
	fmt.Println(CanFinish(NewCar(10, 12), NewTrack(80)))
}
