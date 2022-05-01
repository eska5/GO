
# Excercise 9.Vehicle Purchase

# Instructions
In this exercise you'll be organizing races between various types of remote controlled cars. Each car has its own speed and battery drain characteristics.

Cars start with full (100%) batteries. Each time you drive the car using the remote control, it covers the car's speed in meters and decreases the remaining battery percentage by its battery drain.

If a car's battery is below its battery drain percentage, you can't drive the car anymore.

Each race track has its own distance. Cars are tested by checking if they can finish the track without running out of battery.


# Tasks
1. Creating a remote controlled car

Define a `Car` struct with the following `int` type fields:
- battery
- batteryDrain
- speed
- distance

Allow creating a remote controlled car by defining a function `NewCar` that takes the speed of the car in meters, and the battery drain percentage as its two parameters (both of type `int`) and returns a `Car` instance

2. Creating a race track

Define another struct type called `Track` with the field `distance` of type integer. Allow creating a race track by defining a function `NewTrack` that takes the track's distance in meters as its sole parameter (which is of type `int`)

3. Drive the car

Implement the `Drive` function that updates the number of meters driven based on the car's speed, and reduces the battery according to the battery drainage

4. Check if a remote controlled car can finish a race

To finish a race, a car has to be able to drive the race's distance. This means not draining its battery before having crossed the finish line. Implement the `CanFinish` function that takes a `Car` and a `Track` instance as its parameter and returns `true` if the car can finish the race; otherwise, return `false`. Assume the car is just starting the race

# Solution
``` 
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

```

# Struct
In Go, a `struct` is a sequence of named elements called fields, each field has a name and type. The name of a field must be unique within the struct. Structs can be compared with classes in the Object-Oriented Programming paradigm.

You create a new struct by using the `type` keyword and the built-in type `struct`, and explicitly define the name and type of the fields. For example, to define a `Car` struct:
```
type Car struct{
    brand string
    age int
}
```
Once you have defined the `struct`, you need to create a new instance defining the fields using their field name in any order:
```
car := Car{
	brand: "ferrari",
	age: 3,
}
```
To read or modify instance fields, use the `.` notation:
```
car.age = 5
```

