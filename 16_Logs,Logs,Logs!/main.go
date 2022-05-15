package main

import (
	"fmt"
	"unicode/utf8"
)

// Task 1
func Application(log string) string {
	for _, runeInLog := range log {
		if runeInLog == '❗' {
			return "recommendation"
		} else if runeInLog == '🔍' {
			return "search"
		} else if runeInLog == '☀' {
			return "weather"
		}
	}
	return "default"
}

// Task 2
func Replace(log string, oldRune, newRune rune) string {
	result := []rune(log)
	for index, runeInLog := range result {
		if oldRune == runeInLog {
			result[index] = newRune
		}
	}
	return string(result)
}

// Task 3
func WithinLimit(log string, limit int) bool {
	numberOfRunes := utf8.RuneCountInString(log)
	if numberOfRunes <= limit {
		return true
	} else {
		return false
	}
}

//main
func main() {
	fmt.Println(Application("executed search 🔍"))
	fmt.Println(Replace("❗ recommended product", '❗', '?'))
	fmt.Println(WithinLimit("❗ recommended product", 10))
}
