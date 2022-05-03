package main

import "fmt"

// Task 1
func Units() map[string]int {
	store := map[string]int{}
	store["quarter_of_a_dozen"] = 3
	store["half_of_a_dozen"] = 6
	store["dozen"] = 12
	store["small_gross"] = 120
	store["gross"] = 144
	store["great_gross"] = 1728
	return store
}

// Task 2
func NewBill() map[string]int {
	bill := map[string]int{}
	return bill
}

// Task 3
func AddItem(bill, units map[string]int, item, unit string) bool {
	valueOfUnit, exists := units[unit]
	if !exists {
		return false
	} else {
		valueOfBill, exists := bill[item]
		if !exists {
			bill[item] = valueOfUnit
			return true
		} else {
			bill[item] = valueOfBill + valueOfUnit
			return true
		}
	}
}

// Task 4
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	_, exists := bill[item]
	if !exists {
		return false
	} else {
		_, exists := units[unit]
		if !exists {
			return false
		} else if bill[item] < units[unit] {
			return false
		} else if bill[item] == units[unit] {
			delete(bill, item)
			return true
		} else {
			bill[item] = bill[item] - units[unit]
			return true
		}
	}
}

// Task 5
func GetItem(bill map[string]int, item string) (int, bool) {
	_, exists := bill[item]
	if !exists {
		return 0, false
	} else {
		return bill[item], true
	}
}

//main
func main() {
	fmt.Println(Units())
	fmt.Println(NewBill())
	fmt.Println(AddItem(NewBill(), Units(), "carrot", "dozen"))
	fmt.Println(RemoveItem(NewBill(), Units(), "carrot", "dozen"))
	fmt.Println(GetItem(NewBill(), "carrot"))
}
