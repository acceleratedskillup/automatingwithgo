package main

import (
	"fmt"
	"math/rand"
)

func getRandomWrongAnswers(wrongAnswers *[]string, countries []string, correctAnswer string) {
	for len(*wrongAnswers) < numberOfQuestionOptions-1 {
		wrong := capitals[countries[rand.Intn(len(countries))]]
		if wrong != correctAnswer && !contains(wrongAnswers, wrong) {
			*wrongAnswers = append(*wrongAnswers, wrong)
		}
	}
	fmt.Printf("util) s-addr:%p a-addr:%p\n", wrongAnswers, &(*wrongAnswers)[0])

}

func shuffleSlice(inputSlice []string) {
	rand.Shuffle(len(inputSlice), func(i, j int) {
		inputSlice[i], inputSlice[j] = inputSlice[j], inputSlice[i]
	})
}

func contains(slice *[]string, item string) bool {
	for _, a := range *slice {
		if a == item {
			return true
		}
	}
	return false
}

func index(slice []string, item string) int {
	for i, a := range slice {
		if a == item {
			return i
		}
	}
	return -1
}
