
# Excercise 8.Cars Assemble

# Instructions
In this exercise you'll be writing code to analyze the production in a car factory.


# Tasks
1. Calculate the number of working cars produced per hour

The cars are produced on an assembly line. The assembly line has a certain speed, that can be changed. The faster the assembly line speed is, the more cars are produced. However, changing the speed of the assembly line also changes the number of cars that are produced successfully, that is cars without any errors in their production.

Implement a function that takes in the number of cars produced per hour and the success rate and calculates the number of successful cars made per hour. The success rate is given as a percentage, from `0` to `100`

2. Calculate the number of working cars produced per minute

Implement a function that takes in the number of cars produced per hour and the success rate and calculates how many cars are successfully produced each minute

3. Calculate the cost of production

Each car normally costs $10,000 to produce individually, regardless of whether it is successful or not. But with a bit of planning, 10 cars can be produced together for $95,000.

For example, 37 cars can be produced in the following way: 37 = 3 x groups of ten + 7 individual cars

So the cost for 37 cars is: 3*95,000+7*10,000=355,000

Implement the function `CalculateCost` that calculates the cost of producing a number of cars, regardless of whether they are successful
# Solution
``` 
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
```

# Numbers
types of numeric types:
- `int`
- `float64`
- `uint`
just like in c++ we can also declare if we want for example 32 bit of float for example `float32`.

# Arithmetic Operators
Arithmetic operators:
- `+`
- `-`
- `*`
- `/`
- `%`

Division of integers drops the reminder for example `8/3 = 2`.

Go also supports shorthand assignments such as:
- `x+=y`
- `x-=y`
- `x++`
- `x--`

# Converting between types
To convert already declared variable such as `int x` to `float32`:
```
var x int = 24
f := float32(x)
```


# Arithmetic operations
