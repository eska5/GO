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
