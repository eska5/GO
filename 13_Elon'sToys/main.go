package main

import (
	"fmt"
	"strconv"
)

// Car implements a remote controlled car.
type Car struct {
	speed        int
	batteryDrain int

	battery  int
	distance int
}

// NewCar creates a new car with given specifications.
func NewCar(speed, batteryDrain int) *Car {
	return &Car{
		speed:        speed,
		batteryDrain: batteryDrain,
		battery:      100,
	}
}

// Task 1
func (car *Car) Drive() {
	if car.battery-car.batteryDrain > 0 {
		car.battery -= car.batteryDrain
		car.distance += car.speed
		fmt.Println(car)
	}
}

// Task 2
func (car Car) DisplayDistance() string {
	return "Driven " + strconv.Itoa(car.distance) + " meters"
}

// Task 3
func (car Car) DisplayBattery() string {
	return "Battery at " + strconv.Itoa(car.battery) + "%"
}

// Task 4
func (car Car) CanFinish(trackDistance int) bool {
	if trackDistance/car.speed*car.batteryDrain > car.battery {
		return false
	}
	return true
}
