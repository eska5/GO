
# Excercise 15.CardTricks

# Instructions
As a magician-to-be, Elyse needs to practice some basics. She has a stack of cards that she wants to manipulate.

To make things a bit easier she only uses the cards 1 to 10.

# Tasks
1. Create a slice with certain cards
When practicing with her cards, Elyse likes to start with her favorite three cards of the deck: 2, 6 and 9. Write a function `FavoriteCards` that returns a slice with those cards in that order.

2. Retrieve a card from a stack
Return the card at position index from the given stack.

If the index is out of bounds (ie. if it is negative or after the end of the stack), we want to return `-1`

3. Exchange a card in the stack
Exchange the card at position `index` with the new card provided and return the adjusted stack. Note that this will modify the input slice which is the expected behavior.

If the index is out of bounds (ie. if it is negative or after the end of the stack), we want to append the new card to the end of the stack

4. Add cards to the top of the stack
Add the card(s) specified in the `value` parameter at the top of the stack.

If no argument is given for the `value` parameter, then the result equals the original slice.

5. Remove a card from the stack
Remove the card at position `index` from the stack and return the stack. Note that this may modify the input slice which is ok.

If the index is out of bounds (ie. if it is negative or after the end of the stack), we want to leave the stack unchanged

# Solution
``` 
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

```

# Slices
Slices in Go are similar to lists or arrays in other languages. They hold a number of elements of a specific type (or interface).

Slices in Go are based on arrays. Arrays have a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array.

A slice is written like `[]T` with `T` being the type of the elements in the slice:
```
var empty []int                 // an empty slice
withData := []int{0, 1, 2, 3, 4, 5}  // a slice pre-filled with some data
```

You can get/set an element at a given zero-based index using square-bracket notation:
```
withData[1] = 5
x := withData[1] // x is now 5
```

You can create a new slice from an existing slice by getting a range of elements, once again using square-bracket notation, but specifying both a starting (inclusive) and ending (exclusive) index. If you don't specify a starting index, it defaults to 0. If you don't specify an ending index, it defaults to the length of the slice.
```
newSlice := withData[2:4] // newSlice == []int{2,3}
```

It is common to add new elements to the end of a slice, and so Go provides a built-in `append` function.
```
a := []int{1, 3}
b := append(a, 4, 2) // b == []int{1,3,4,2}
```
Note that append is a `variadic` function, see details below.

# Variadic Functions
Usually, functions in Go accept only a fixed number of arguments. However, it is also possible to write variadic functions in Go.

A variadic function is a function that accepts a variable number of arguments.

If the type of the last parameter in a function definition is prefixed by ellipsis `...`, then the function can accept any number of arguments for that parameter.
```
func find(a int, b ...int) {
    // ...
}
```
In the above function, parameter b is variadic and we can pass 0 or more arguments to b.
```
find(5, 6)
find(5, 6, 7)
find(5)
```
# CAUTION - Variadic parameter must be the last parameter of the function !!!

The way variadic functions work is by converting the variable number of arguments to a slice of the type of the variadic parameter.

Here is an example of an implementation of a variadic function.
```
func find(num int, nums ...int) {
    fmt.Printf("type of nums is %T\n", nums)

    for i, v := range nums {
        if v == num {
            fmt.Println(num, "found at index", i, "in", nums)
            return
        }
    }

    fmt.Println(num, "not found in ", nums)
}

func main() {
    find(89, 90, 91, 95)
    // Output:
    // type of nums is []int
    // 89 not found in  [90 91 95]

    find(45, 56, 67, 45, 90, 109)
    // Output:
    // type of nums is []int
    // 45 found at index 2 in [56 67 45 90 109]

    find(87)
    // Output:
    // type of nums is []int
    // 87 not found in  []
}
```

In line `find(89, 90, 91, 95)` of the program above, the variable number of arguments to the find function are `90`, `91` and `95`. The `find` function expects a variadic int parameter after `num`. Hence these three arguments will be converted by the compiler to a slice of type `int []int{90, 91, 95}` and then it will be passed to the find function as nums.

Sometimes you already have a slice and want to pass that to a variadic function. This can be achieved by passing the slice followed by `...`. That will tell the compiler to use the slice as is inside the variadic function. The step described above where a slice is created will simply be omitted in this case.
```
list := []int{1, 2, 3}
find(1, list...) // "find" defined as shown above
```
