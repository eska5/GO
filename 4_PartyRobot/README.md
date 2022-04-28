
# Excercise 4.Party Robot

# Instructions
Once there was an eccentric programmer living in a strange house with barred windows. One day he accepted a job from an online job board to build a party robot. The robot is supposed to greet people and help them to their seats. The first edition was very technical and showed the programmer's lack of human interaction. Some of which also made it into the next edition.

# Tasks
1. Welcome a new guest to the party

Implement the `Welcome` function to return a welcome message using the given name


2. Welcome a new guest to the party whose birthday is today

Implement the `HappyBirthday` function to return a birthday message using the given name and age of the person. Unfortunately the programmer is a bit of a show-off, so the robot has to demonstrate its knowledge of every guest's birthday.

3. Give directions

Implement the `AssignTable` function to give directions. It should accept 5 parameters.

- The name of the new guest to greet (`string`)
- The table number (`int`)
- The name of the seatmate (`string`)
- The direction where to find the table (`string`)
- The distance to the table (`float64`)

The exact result format can be seen in the example below.

The robot provides the table number in a 3 digits format. If the number is less than 3 digits it gets extra leading zeroes to become 3 digits (e.g. 3 becomes 003). The robot also mentions the distance of the table. Fortunately only with a precision that's limited to 1 digit.

# Solution
``` 
package main

import "fmt"

// Task 1
func Welcome(name string) string {
	Output := fmt.Sprintf("Welcome to my party, %s!", name)
	return Output
}

// Task 2
func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
}

// Task 3
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	return Welcome(name) + fmt.Sprintf("\nYou have been assigned to table %03d. Your table is %s, exactly %.1f meters from here.\nYou will be sitting next to %s.", table, direction, distance, neighbor)
}

func main() {
	fmt.Println(Welcome("Jakub"))
	fmt.Println(HappyBirthday("Jakub", 25))
	fmt.Println(AssignTable("Chihiro", 22, "Frank", "top left", 27.23456))
}
```

# Packages
Go application is organized in packages. A package is a collection of files in the same folder. Source files in folder
must have main package name(folder that they are in) on top of the package.

Go provides a bunch of useful libraries that you can import using `import` keyword. Standard libraries are imported by their names.
```
import "fmt"
```

We address imported package using import name and .
```
fmt.PrintLn("Hello!")
```

Go determines if the item such as variable or function can be accessed only if item name starts with capital letters.
-`public` -> Hello("hello")
-`private` -> hello("hello")

# String Formatting
`fmt` which stands for format is a in-build package which offers functions to manipulate the format of input as well as output.
One of the most common functions used is `Sprintf` which uses verbs like `%s` like in `C` to output string variable in output.
```
fmt.Sprintf("Hello %s", stringVar)
```

In Go floating points are formated with verbs such as:
- `%g` -> compact representation
- `%e` -> exponent
- `%f` -> non exponent
There is also a posibility with controll of numeric positions.
```
fmt.Sprintf("%.2f", floatVar) \ 9.2222 -> 9.22
fmt.Sprintf("%03d", intVar) \ 22 -> 022
```

Format contains more useful functions:
- `Println`
- `Printf`
- `Sprintf`