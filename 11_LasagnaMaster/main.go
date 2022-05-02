package main

import "fmt"

// Task 1
func PreparationTime(layers []string, avgPrepTime int) int {
	if avgPrepTime == 0 {
		return len(layers) * 2
	} else {
		return len(layers) * avgPrepTime
	}
}

// Task 2
func Quantities(layers []string) (int, float64) {
	var sauce float64 = 0
	noodles := 0
	for i := 0; i < len(layers); i++ {
		if layers[i] == "sauce" {
			sauce += 0.2
		} else if layers[i] == "noodles" {
			noodles += 50
		}
	}
	return noodles, sauce
}

// Task 3
func AddSecretIngredient(friendList, myList []string) {
	myList[len(myList)-1] = friendList[len(friendList)-1]
}

// Task 4
func ScaleRecipe(quantities []float64, portions int) []float64 {
	var amounts []float64
	for i := 0; i < len(quantities); i++ {
		amounts = append(amounts, quantities[i]*float64(portions)/2)
	}
	return amounts
}

//main
func main() {
	fmt.Println(PreparationTime([]string{"sauce", "noodles", "sauce", "meat", "mozzarella", "noodles"}, 3))
	fmt.Println(Quantities([]string{"sauce", "noodles", "sauce", "meat", "mozzarella", "noodles"}))

	friendsList := []string{"noodles", "sauce", "mozzarella", "kampot pepper"}
	myList := []string{"noodles", "meat", "sauce", "mozzarella", "?"}
	AddSecretIngredient(friendsList, myList)
	fmt.Println(myList)
	fmt.Println(ScaleRecipe([]float64{1.2, 3.6, 10.5}, 4))
}
