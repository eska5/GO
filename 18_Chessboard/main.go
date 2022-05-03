package main

import "fmt"

type Rank []bool
type Chessboard map[string]Rank

// Task 1
func CountInRank(cb Chessboard, rank string) int {
	sum := 0
	for _, x := range cb[rank] {
		if x {
			sum++
		}
	}
	return sum
}

// Task 2
func CountInFile(cb Chessboard, file int) int {
	sum := 0
	if file < 1 || file > 8 {
		return sum
	} else {
		for _, x := range cb {
			if x[file-1] {
				sum++
			}
		}
		return sum
	}
}

// Task 3
func CountAll(cb Chessboard) int {
	return len(cb) * len(cb)
}

// Task 4
func CountOccupied(cb Chessboard) int {
	sum := 0
	for index := range cb {
		sum += CountInRank(cb, index)
	}
	return sum
}

//main
func main() {
	fmt.Println("No chessboard :O")
}
