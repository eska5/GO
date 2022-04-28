
# Excercise 5.Weather Forecast

# Instructions
Being hired by a big weather forecast company, your job is to maintain their code base. Scrolling through the code, you find it hard to follow what the code is actually doing. Good thing you know just the right thing to do!

# Tasks
1. Document package weather

Write a package comment for `package weather` that describes its contents. The package comment should introduce the package and provide information relevant to the package as a whole.

2. Document the CurrentCondition variable

Write a comment for the package variable `CurrentCondition`.

This should tell any user of the package what information the variable stores, and what they can do with it.

3. Document the CurrentLocation variable

Just like the previous step, write a comment for `CurrentLocation`.

4. Document the Forecast() function

Write a function comment for `Forecast()`. Your comments must describe what the function does, but not how it does it.

# Solution
``` 
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

```

# Documentation comments
Comments in go play an important role in documentation. `go doc` command is used to extract comments to create documentation about go packages.
Documentation should be a complete sentence and start with name of described thing with period at the end of the sentence.

Every comment should be exactly above the described thing such as function, package, variable.

# Package comments
package comments should be written above a package declaration for example `package main` and begin with Package main with the capital letter.
Example: 
```
// Package main provides function to convert
// string to int.
package main
```
 
# Function comments
Function comment should be written directly above the function declaration. It should be a full sentence that starts with the function name.
It should also explain what arguments the function takes, what it does with them, and what its return values mean, ending in a period)
Example:
```
// Add returns a sum of values integer a and b and returns sum of them in the integer format.
func Add(a, b int) int {
	return a + b
}
```