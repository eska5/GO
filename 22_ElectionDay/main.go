package main

import (
	"fmt"
	"strconv"
)

type ElectionResult struct {
	Name  string
	Votes int
}

// Task 1
func NewVoteCounter(initialVotes int) *int {
	var initialVotesCounter *int = &initialVotes
	return initialVotesCounter
}

// Task 2
func VoteCount(counter *int) int {
	if counter == nil {
		return 0
	}
	return *counter
}

// Task 3
func IncrementVoteCount(counter *int, increment int) {
	*counter = *counter + increment
}

// Task4
func NewElectionResult(candidateName string, votes int) *ElectionResult {
	var result *ElectionResult
	result = &ElectionResult{Name: candidateName, Votes: votes}
	return result
}

// Task 5
func DisplayResult(result *ElectionResult) string {
	return result.Name + " (" + strconv.Itoa(result.Votes) + ")"
}

// Task 6
func DecrementVotesOfCandidate(results map[string]int, candidate string) {
	results[candidate] -= 1
}
func main() {
	var init int = 4
	var initPointer *int = &init

	var elecResult = ElectionResult{
		Name:  "string",
		Votes: 1234,
	}
	var finalResults = map[string]int{
		"Mary": 10,
		"John": 51,
	}
	fmt.Println(*NewVoteCounter(init))
	fmt.Println(VoteCount(initPointer))
	IncrementVoteCount(initPointer, 3)
	fmt.Println(VoteCount(initPointer))
	fmt.Println(NewElectionResult("Ada", 4))
	fmt.Println(DisplayResult(&elecResult))
	DecrementVotesOfCandidate(finalResults, "John")

}
