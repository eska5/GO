
# Excercise 10.GrossStore

# Instructions
A friend of yours has an old wholesale store called Gross Store. The name comes from the quantity of the item that the store sell: it's all in `gross unit`. Your friend asked you to implement a point of sale (POS) system for his store. First, you want to build a prototype for it. In your prototype, your system will only record the quantity. Your friend gave you a list of measurements to help you:

Unit  | Score
------------- | -------------
quarter_of_a_dozen  | 3
half_of_a_dozen	  | 6
dozen | 12
small_gross | 120
gross | 144
great_gross | 1728

# Tasks
1. Store the unit of measurement in your program
In order to use the measurement, you need to store the measurement in your program.

2. Create a new customer bill
You need to implement a function that create a new (empty) bill for the customer.

3. Add an item to the customer bill
To implement this, you'll need to:

- Return `false` if the given `unit` is not in the `units` map.
- Otherwise add the item to the customer `bill`, indexed by the item name, then return `true`.
- If the item is already present in the bill, increase its quantity by the amount that belongs to the provided `unit`.

4. Remove an item from the customer bill
To implement this, you'll need to:

- Return `false` if the given item is `not` in the bill
- Return `false` if the given `unit` is not in the `units` map.
- Return `false` if the new quantity would be less than 0.
- If the new quantity is 0, completely remove the item from the `bill` then return `true`.
- Otherwise, reduce the quantity of the item and return `true`.

5. Return the quantity of a specific item that is in the customer bill
To implement this, you'll need to:

- Return `0` and `false` if the `item` is not in the bill.
- Otherwise, return the quantity of the item in the `bill` and `true`.

# Solution
``` 
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
```

# Map

In Go, `map` is a built-in data type that maps keys to values. In other programming languages, you might be familiar with the concept of `map` as a dictionary, hash table, key/value store or an associative array.

Syntactically, `map` looks like this:
```
map[KeyType]ElementType
```
It is also important to know that each key is unique, meaning that assigning the same key twice will overwrite the value of the corresponding key.

To create a map, you can do:
```
 // With map literal
  foo := map[string]int{}
```
or
```
  // or with make function
  foo := make(map[string]int)
```

Here are some operations that you can do with a map
```
  // Add a value in a map with the `=` operator:
  foo["bar"] = 42
  // Here we update the element of `bar`
  foo["bar"] = 73
  // To retrieve a map value, you can use
  baz := foo["bar"]
  // To delete an item from a map, you can use
  delete(foo, "bar")
```

If you try to retrieve the value for a key which does not exist in the map, it will return the zero value of the value type. This can confuse you, especially if the default value of your `ElementType` (for example, 0 for an int), is a valid value. To check whether a key exists in your map, you can use

```
value, exists := foo["baz"]
  // If the key "baz" does not exist,
  // value: 0; exists: false
```

