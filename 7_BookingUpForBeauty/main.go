package main

import (
	"fmt"
	"time"
)

// Task 1
func Schedule(date string) time.Time {
	// we can use `_` as a placeholder that will never be used.
	timeVar, _ := time.Parse("1/2/2006 15:04:05", date) // can't be 01 because for number "7" it is "7" not "07".
	return timeVar
}

// Task 2
func HasPassed(date string) bool {
	timeVar, _ := time.Parse("January 2, 2006 15:04:05", date)
	isAfter := timeVar.After(time.Now())
	return !isAfter
}

// Task 3
func IsAfternoonAppointment(date string) bool {
	timeVar, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
	if timeVar.Hour() < 18 && timeVar.Hour() >= 12 {
		return true
	} else {
		return false
	}
}

// Task 4
func Description(date string) string {
	timeVar := Schedule(date)
	return "You have an appointment on " + timeVar.Format("Monday, January 2, 2006, at 15:04.")
}

// Task 5
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), 9, 15, 00, 00, 00, 00, time.UTC)
}

//main
func main() {
	fmt.Println(Schedule("7/25/2019 13:45:23"))
	fmt.Println(HasPassed("December 9, 2002 11:59:59"))
	fmt.Println(IsAfternoonAppointment("Thursday, July 25, 2019 13:45:00"))
	fmt.Println(Description("6/6/2005 10:30:00"))
	fmt.Println(AnniversaryDate())
}
