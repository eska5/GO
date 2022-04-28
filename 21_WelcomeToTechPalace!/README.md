
# Excercise 21.Welcome to tech Palace !

# Instructions
There is an appliance store called "Tech Palace" nearby. The owner of the store recently installed a big display to use for marketing messages and to show a special greeting when customers scan their loyalty cards at the entrance. The display consists of lots of small LED lights and can show multiple lines of text.

The store owner needs your help with the code that is used to generate the text for the new display.


# Tasks
1. Create the welcome message

For most customers who scan their loyalty cards, the store owner wants to see `Welcome to the Tech Palace,` followed by the name of the customer in capital letters on the display.

Implement the function `WelcomeMessage` that accepts the name of the customer as a `string` argument and returns the desired message as a `string`.

2. Add a fancy border

For loyal customers that buy a lot at the store, the owner wants the welcome display to be more fancy by adding a line of stars before and after the welcome message. They are not sure yet how many stars should be in the lines so they want that to be configurable.

Write a function `AddBorder` that accepts a welcome message (a `string`) and the number of stars per line (type `int`) as arguments. It should return a `string` that consists of 3 lines, a line with the desired number of stars, then the welcome message as it was passed in, then another line of stars.

3. Clean up old marketing messages

Before installing this new display, the store had a similar display that could only show non-configurable, static lines. The owner would like to reuse some of the old marketing messages on the new display. However, the data already includes a star border and some unfortunate whitespaces. Your task is to clean up the messages so they can be re-used.

Implement a function `CleanUpMessage` that accepts the old marketing message as a string. The function should first remove all stars from the text and afterwards remove the leading and trailing whitespaces from the remaining text. The function should then return the cleaned up message.

# Solution
``` 
package main

import (
	"fmt"
	"strings"
)

// Task 1
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// Task 2
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	return strings.Repeat("*", numStarsPerLine) + "\n" + welcomeMsg + "\n" + strings.Repeat("*", numStarsPerLine)
}

// Task 3
func CleanupMessage(oldMsg string) string {
	return strings.TrimSpace(strings.ReplaceAll(oldMsg, "*", ""))
}

func main() {
	fmt.Println(WelcomeMessage("Judy"))
	fmt.Println(AddBorder("Judy", 10))
	fmt.Println(CleanupMessage("*******\n\t*	* sale *	*\n\t*******"))
}
```

# String
String in Go is `immutable sequence of bytes`

String is defined like this `"example"`

We can also connect strings with `+` operator
```
fullText := "fra" + "nce" 
```

There are also special characters such as:
- `\n` -> new line
- `\t` -> tab

We are going to use `strings` package that contains many string related functionality.

# Strings import functions
- `strings.TrimSpace("a b c")` -> removes whitespace(" ")
- `strings.ToLower("Xyz")` -> makes string a lowercase
- `strings.ToUpper("Xyz")` -> makes string a uppercase
- `strings.Repeat("xyz",3)`  -> makes string repeat n times    
- `strings.trim("strrr","r")` -> trims string with the second parameter
- `strings.trimSpace("\t\n   strrr \t\n    ","r")` -> trims whitespaces from given string