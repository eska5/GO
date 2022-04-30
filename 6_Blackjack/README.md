
# Excercise 6.Blackjack

# Instructions
In this exercise we will simulate the first turn of a Blackjack game.

You will receive two cards and will be able to see the face up card of the dealer. All cards are represented using a string such as "ace", "king", "three", "two", etc. The values of each card are:
- two -> 2
- three -> 3
- four -> 4
- five -> 5
- six -> 6
- seven -> 7
- eight -> 8
- nine -> 9
- ten -> 10
- jack -> 10
- queen -> 10
- king -> 10
- ace -> 11
- other -> 0
`Note:` Commonly, aces can take the value of 1 or 11 but for simplicity we will assume that they can only take the value of 11.

Depending on your two cards and the card of the dealer, there is a strategy for the first turn of the game, in which you have the following options:
- Stand (S)
- Hit (H)
- Split (P)
- Automatically win (W)

Although not optimal yet, you will follow the strategy your friend Alex has been developing, which is as follows:
- If you have a pair of aces you must always split them.
- If you have a Blackjack (two cards that sum up to a value of 21), and the dealer does not have an ace, a figure or a ten then you automatically win. If the dealer does have any of those cards then you'll have to stand and wait for the reveal of the other card.
- If your cards sum up to a value within the range [17, 20] you should always stand.
- If your cards sum up to a value within the range [12, 16] you should always stand unless the dealer has a 7 or higher, in which case you should always hit.
- If your cards sum up to 11 or lower you should always hit.

# Tasks
1. Calculate the value of any given card.

Implement a function to calculate the numerical value of a card:

2. Implement the decision logic for the first turn.

Write a function that implements the decision logic as described above:

# Solution
``` 
package main

import "fmt"

// Task 1
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "ten", "jack", "queen", "king":
		return 10
	case "nine":
		return 9
	case "eight":
		return 8
	case "seven":
		return 7
	case "six":
		return 6
	case "five":
		return 5
	case "four":
		return 4
	case "three":
		return 3
	case "two":
		return 2
	default:
		return 0
	}
}

// Task 2
func FirstTurn(card1, card2, dealerCard string) string {
	switch {
	case ParseCard(card1)+ParseCard(card2) == 22:
		return "P"
	case ParseCard(card1)+ParseCard(card2) == 21 && ParseCard(dealerCard) < 10:
		return "W"
	case ParseCard(card1)+ParseCard(card2) == 21 && ParseCard(dealerCard) >= 10:
		return "S"
	case ParseCard(card1)+ParseCard(card2) < 21 && ParseCard(card1)+ParseCard(card2) > 16:
		return "S"
	case ParseCard(card1)+ParseCard(card2) < 17 && ParseCard(card1)+ParseCard(card2) > 11 && ParseCard(dealerCard) < 7:
		return "S"
	case ParseCard(card1)+ParseCard(card2) < 17 && ParseCard(card1)+ParseCard(card2) > 11 && ParseCard(dealerCard) >= 7:
		return "H"
	default:
		return "H"
	}
}

func main() {
	fmt.Println(ParseCard("ace"))
	fmt.Println(FirstTurn("ace", "king", "nine"))
}
```

# Switch case
Like other languages, Go also provides a switch statement. Switch statements are a shorter way to write long if ... else if statements. To make a switch, we start by using the keyword switch followed by a value or expression. We then declare each one of the conditions with the case keyword. We can also declare a default case, that will run when none of the previous case conditions matched:
```
operatingSystem := "windows"

switch operatingSystem {
case "windows":
    // do something if the operating system is windows
case "linux":
    // do something if the operating system is linux
case "macos":
    // do something if the operating system is macos
default:
    // do something if the operating system is none of the above
} 
```

One interesting thing about switch statements, is that the value after the switch keyword can be omitted, and we can have boolean conditions for each case:
```
age := 21

switch {
case age > 20 && age < 30:
    // do something if age is between 20 and 30
case age == 10:
    // do something if age is equal to 10
default:
    // do something else for every other case
}
```


