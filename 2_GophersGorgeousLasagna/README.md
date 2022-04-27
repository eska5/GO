
# Excercise 2.Gophers Georgeous Lasagna

# Instructions
In this exercise you're going to write some code to help you cook a brilliant lasagna from your favorite cooking book.

You have four tasks, all related to the time spent cooking the lasagna.

The objectives are simple:

- Define the `OvenTime` constant with how many minutes the lasagna should be in the oven. According to the cooking book, the expected oven time in minutes is 40
- Define the `RemainingOvenTime()` function that takes the actual minutes the lasagna has been in the oven as a parameter and returns how many minutes the lasagna still has to remain in the oven, based on the expected oven time in minutes from the previous task.
- Define the `PreparationTime` function that takes the number of layers you added to the lasagna as a parameter and returns how many minutes you spent preparing the lasagna, assuming each layer takes you 2 minutes to prepare.
- Define the `ElapsedTime` function that takes two parameters: the first parameter is the number of layers you added to the lasagna, and the second parameter is the number of minutes the lasagna has been in the oven. The function should return how many minutes in total you've worked on cooking the lasagna, which is the sum of the preparation time in minutes, and the time in minutes the lasagna has spent in the oven at the moment.

# Solution
``` 
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

```

# Basics

Go is a statically typed, compiled programming language. Go syntax is simillar to C and Python.
Go is also known as Golang.

In this readme we will cover topics such as: Packages, Variables, Constants and Functions.

# Packages
Go applications are organized in packages. Packages are collections of files in the same directory.
if we want to use package from the directory for example `SomePackage/subPackage/ThePack`, we should declare `package ThePack`.
imports from package such as (functions, types, variables, constants) that start with a capital letter are `public and visible`, 
on the other hand when we won't declare them with a capital letter, they are `private and inaccessible`.
- camelCase -> `private and inaccessible`
- PascalCase -> `public and visible` 

Example:
```
package main
```
# Variables
Go is a statically-typed language, which means that all variables `must have a defined type` at compilation.

Variables can be defined by explicitly specifying a type:
```
var variableName int
```

We can also initialize our variable without giving a type, compiler will assing the variable type for us to match the given value.
```
variableName := 10
```

Assigning values are declared with `=` operator. 
Variables once declared are not capable of changing its type.

# Constants
Constant values are unchangeable. To define constant variable we use the `const` keyword.
```
const variableName = 1234
```

# Functions
Go functions can have 0 or more parameters. Parameters must have type declared when being passed for example:
```
func start(variable string){}
```

Values are returned from functions with a keyword `return`

To invoke a function such as start given above we simply type
```
start(variable)
```

Comments are declared in the same way as in c++ which means:
- single line -> `// code`
- multiple lines -> `/* code */`

Complete example of Go function:
```
// public function
func PublicFunction (variable string) string {
    return variable
}

// private function
func publicFunction (variable string) string {
    return variable
}
```