package main

import "fmt"

func castVote(votes map[string]string, code, performer string) {
	if _, exists := votes[code]; !exists {
		votes[code] = performer
	}
}

func tallyVotes(votes map[string]string) map[string]int {
	tally := make(map[string]int)
	for _, performer := range votes {
		tally[performer]++
	}
	return tally
}

func main() {
	votes := make(map[string]string)

	castVote(votes, "A123", "Singer1")
	castVote(votes, "B456", "Singer2")
	castVote(votes, "A123", "Singer3") // This vote should not be counted

	results := tallyVotes(votes)
	fmt.Println("Talent Show Results:", results)
}
