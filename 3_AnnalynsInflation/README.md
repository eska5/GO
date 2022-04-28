
# Excercise 3.Annalyn's Inflation

# Instructions
In this exercise, you'll be implementing the quest logic for a new RPG game a friend is developing. The game's main character is Annalyn, a brave girl with a fierce and loyal pet dog. Unfortunately, disaster strikes, as her best friend was kidnapped while searching for berries in the forest. Annalyn will try to find and free her best friend, optionally taking her dog with her on this quest.

After some time spent following her best friend's trail, she finds the camp in which her best friend is imprisoned. It turns out there are two kidnappers: a mighty knight and a cunning archer.

Having found the kidnappers, Annalyn considers which of the following actions she can engage in:

- Fast attack: a fast attack can be made if the knight is sleeping, as it takes time for him to get his armor on, so he will be vulnerable.
- Spy: the group can be spied upon if at least one of them is awake. Otherwise, spying is a waste of time.
- Signal prisoner: the prisoner can be signaled using bird sounds if the prisoner is awake and the archer is sleeping, as archers are trained in bird signaling so they could intercept the message.
- Free prisoner: Annalyn can try sneaking into the camp to free the prisoner. This is a risky thing to do and can only succeed in one of two ways:
	- If Annalyn has her pet dog with her she can rescue the prisoner if the archer is asleep. The knight is scared of the dog and the archer will not have time to get ready before Annalyn and the prisoner can escape.
	- If Annalyn does not have her dog then she and the prisoner must be very sneaky! Annalyn can free the prisoner if the prisoner is awake and the knight and archer are both sleeping, but if the prisoner is sleeping they can't be rescued: the prisoner would be startled by Annalyn's sudden appearance and wake up the knight and archer.

You have four tasks: to implement the logic for determining if the above actions are available based on the state of the three characters found in the forest and whether Annalyn's pet dog is present or not.

# Tasks
1. Check if a fast attack can be made

Define the `CanFastAttack()` function that takes a boolean value that indicates if the knight is awake. This function returns true if a fast attack can be made based on the state of the knight. Otherwise, returns false

2. Check if the group can be spied upon

Define the `CanSpy()` function that takes three boolean values, indicating if the knight, archer and the prisoner, respectively, are awake. The function returns true if the group can be spied upon, based on the state of the three characters. Otherwise, returns false

3. Check if the prisoner can be signaled

Define the `CanSignalPrisoner()` function that takes two boolean values, indicating if the archer and the prisoner, respectively, are awake. The function returns true if the prisoner can be signaled, based on the state of the two characters. Otherwise, returns false

4. Check if the prisoner can be freed

Define the `CanFreePrisoner()` function that takes four boolean values. The first three parameters indicate if the knight, archer and the prisoner, respectively, are awake. The last parameter indicates if Annalyn's pet dog is present. The function returns true if the prisoner can be freed based on the state of the three characters and Annalyn's pet dog presence. Otherwise, it returns false

# Solution
``` 
package main

import "fmt"

// Task 1
func CanFastAttack(knightIsAwake bool) bool {
	return !knightIsAwake
}

// Task 2
func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
	if knightIsAwake == true || archerIsAwake == true || prisonerIsAwake == true {
		return true
	} else {
		return false
	}
}

// Task 3
func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
	if archerIsAwake == false && prisonerIsAwake == true {
		return true
	} else {
		return false
	}
}

// Task 4
func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
	if petDogIsPresent == true && archerIsAwake == false {
		return true
	} else if petDogIsPresent == false && prisonerIsAwake == true && archerIsAwake == false && knightIsAwake == false {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(CanFastAttack(true))
	fmt.Println(CanSpy(true, false, true))
	fmt.Println(CanSignalPrisoner(true, false))
	fmt.Println(CanFreePrisoner(false, true, false, false))
}

```

# Booleans

Booleans in Go are represented by the predeclared boolean type `bool`, which values can be either `true` or `false`. It's a defined type.
Example:
```
var variable bool \\ this is false
variable1 := true \\ this is true
variable2 := false \\ this is false
```

# Operators

Go supports three logical operators that can evaluate expressions down to Boolean values, returning either true or false.

- `&&` -> `and`, true if both statements are true 
- `||` -> `or`, true if at least one statement is true 
- `!` -> `not`, true if statement is false 

# Knowledge Gained

Besides common operator knowledge I have gained knowledge in topics such as passing parameters of the same type. Example:
```
func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {}
```
Example shows us that we can define passed variable time only on the last variable if previous variables are of the same type.
Moreover we can see that bool is the type of returned variable.

I have also came across `if statement`. Example:
```
if petDogIsPresent == true && archerIsAwake == false {
		return true
	} else if petDogIsPresent == false && prisonerIsAwake == true && archerIsAwake == false && knightIsAwake == false {
		return true
	} else {
		return false
	}
```
We can see that if statements are written just like in python besides `{}` of course.
It is also worth mentioning that we have to write `} else {` in the same line, otherwise run will return an error. 
Else if is almost like in python written in 1 line: `} else if petDogIsPresent == false {`.

