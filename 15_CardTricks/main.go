package main

import "fmt"

// Task 1
func FavoriteCards() []int {
	return []int{2, 6, 9}
}

// Task 2
func GetItem(slice []int, index int) int {
	if len(slice)-1 < index || index < 0 {
		return -1
	} else {
		return slice[index]
	}
}

// Task 3
func SetItem(slice []int, index, value int) []int {
	if len(slice)-1 < index || index < 0 {
		slice = append(slice, value)
		return slice
	} else {
		slice[index] = value
		return slice
	}
}

// Task 4
func PrependItems(slice []int, value ...int) []int {
	slice = append(value, slice...) // i am not sure why ...
	return slice
}

// Task 5
func RemoveItem(slice []int, index int) []int {
	if len(slice)-1 < index || index < 0 {
		return slice
	} else {
		//slice[:index] -> 0 .. index-1
		//adding dots after the second slice helps
		slice = append(slice[:index], slice[index+1:]...)
		return slice
	}
}

//main
func main() {
	fmt.Println(FavoriteCards())
	fmt.Println(GetItem(FavoriteCards(), 1))
	fmt.Println(SetItem(FavoriteCards(), 1, 27))
	fmt.Println(PrependItems(FavoriteCards(), 4))
	fmt.Println(RemoveItem(FavoriteCards(), 1))
}
