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
