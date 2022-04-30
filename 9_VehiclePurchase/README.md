
# Excercise 9.Vehicle Purchase

# Instructions
In this exercise, you are going to write some code to help you prepare to buy a vehicle.

You have three tasks, one to determine if you need a license, one to help you choose between two vehicles and one to estimate the acceptable price for a used vehicle.


# Tasks
1. Determine if you will need a driver's license

Some vehicle kinds require a driver's license to operate them. Assume only the kinds `"car"` and `"truck"` require a license, everything else can be operated without a license.

Implement the `NeedsLicense(kind)` function that takes the kind of vehicle and returns a boolean indicating whether you need a license for that kind of vehicle

2. Choose between two potential vehicles to buy

You evaluate your options of available vehicles. You manage to narrow it down to two options but you need help making the final decision. For that, implement the function `ChooseVehicle(option1, option2)` that takes two vehicles as arguments and returns a decision that includes the option that comes first in lexicographical order.

3. Calculate an estimation for the price of a used vehicle

Now that you made a decision, you want to make sure you get a fair price at the dealership. Since you are interested in buying a used vehicle, the price depends on how old the vehicle is. For a rough estimate, assume if the vehicle is less than 3 years old, it costs 80% of the original price it had when it was brand new. If it is at least 10 years old, it costs 50%. If the vehicle is at least 3 years old but less than 10 years, it costs 70% of the original price.

Implement the `CalculateResellPrice(originalPrice, age)` function that applies this logic using `if`, `else if` and `else` (there are other ways if you want to practice). It takes the original price and the age of the vehicle as arguments and returns the estimated price in the dealership.

# Solution
``` 
package main

import "fmt"

// Task 1
func NeedsLicense(kind string) bool {
	if kind == "truck" || kind == "car" {
		return true
	} else {
		return false
	}
}

// Task 2
func ChooseVehicle(option1, option2 string) string {
	if option1 > option2 {
		return option2 + " is clearly the better choice."
	} else {
		return option1 + " is clearly the better choice."
	}
}

// Task 3
func CalculateResellPrice(originalPrice, age float64) float64 {
	if age <= 3 {
		return originalPrice * 0.8
	} else if age >= 10 {
		return originalPrice * 0.5
	} else {
		return originalPrice * 0.7
	}
}
func main() {
	fmt.Println(NeedsLicense("truck"))
	fmt.Println(ChooseVehicle("ferrari", "lamborghini"))
	fmt.Println(CalculateResellPrice(2100000.0, 4))
}

```

# Comparison
In Go numbers can be compared using the following relational and equality operators:
- `==`
- `!=`
- `<`
- `<=`
- `>`
- `>=`

The result of this comparison is always `true` or `false` 
Comparison operators can also be applied for `string` type.

# If Statements
Conditionals in Go are similar to conditionals in other languages. The underlying type of any conditional operation is the `bool` type, which can have the value of `true` or `false`. Conditionals are often used as flow control mechanisms to check for various conditions.

 For checking a particular case an `if` statement can be used, which executes its code if the underlying condition is `true`.

In scenarios involving more than one case many `if` statements can be chained together using the `else if` and `else` statements.

If statements can also include a short initialization statement that can be used to initialize one or more variables for the if statement. For example:

```
num := 7
if v := 2 * num; v > 10 {
    fmt.Println(v)
} 
```

