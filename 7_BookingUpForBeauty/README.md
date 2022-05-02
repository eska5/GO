
# Excercise 7.BookingUpForBeauty

# Instructions
In this exercise you'll be working on an appointment scheduler for a beauty salon that opened on September 15th in 2012.

You have five tasks, which will all involve appointment dates.

# Tasks
1. Parse appointment date
Implement the `Schedule` function to parse a textual representation of an appointment date into the corresponding `time.Time` format.

2. Check if an appointment has already passed
Implement the `HasPassed` function that takes an appointment date and checks if the appointment was somewhere in the past.

3. Check if appointment is in the afternoon
Implement the IsAfternoonAppointment function that takes an appointment date and checks if the appointment is in the afternoon (>= 12:00 and < 18:00).

4. Describe the time and date of the appointment
Implement the `Description` function that takes an appointment date and returns a description of that date and time.

5. Return the anniversary date of the salon's opening
Implement the `AnniversaryDate` function that returns the anniversary date of the salon's opening for the current year in UTC.

Assuming the current year is 2020.

# Solution
``` 
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

```

# Time
In Go, functionality for working with times is provided by the `time` package. The types and methods in this package allow us to manipulate times, get the current time, determine elapsed time, parse times from strings, and more.

To work with time, you will usually call a method on a `Time` instance, but there are also some functions called on the `time` package itself.

Parsing strings works a little differently in Go. Many languages use a format string with various codes like YYYY for four-digit year, MM for two-digit month, etc. Go instead uses an exact date, `Mon Jan 2 15:04:05 -0700 MST 2006`, to parse strings. You can rearrange these components or omit them as necessary, and you can spell out month and day names, or use their number values, but in order for Go to understand the parts of the string you are parsing, the corresponding parts of the format string must be from this exact date.

# Time import functions
- `timeVar, _ := time.Parse("1/2/2006 15:04:05", date)`
- `time.After()`
- `time.Now()`
- `time.Hour()`
- `time.Year()`
- `time.Format("Monday, January 2, 2006, at 15:04.")`
- `time.Date(2000, 9, 15, 00, 00, 00, 00, time.UTC)`