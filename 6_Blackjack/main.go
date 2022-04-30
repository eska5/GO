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
