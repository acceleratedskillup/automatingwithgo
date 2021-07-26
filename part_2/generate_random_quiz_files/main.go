package main

import (
	"fmt"
	"os"
)

func main() {

	for quizNum := 1; quizNum <= totalQuizzes; quizNum++ {
		quizFile, err := os.Create(fmt.Sprintf("quiz%d.txt", quizNum))
		if err != nil {
			fmt.Println("Failed to create quiz file:", err)
			return
		}
		defer quizFile.Close()

		answerFile, err := os.Create(fmt.Sprintf("quiz_answers%d.txt", quizNum))
		if err != nil {
			fmt.Println("Failed to create answer file:", err)
			return
		}
		defer answerFile.Close()

		quizFile.WriteString("Name:\n\nDate:\n\nPeriod:\n\n")
		quizFile.WriteString("World Capitals Quiz\n\n")

		countries := make([]string, 0, len(capitals))
		for country := range capitals {
			countries = append(countries, country)
		}
		shuffleSlice(countries)

		for questionNum := 1; questionNum <= numberOfQuestions; questionNum++ {
			correctAnswer := capitals[countries[questionNum-1]]
			wrongAnswers := make([]string, 0, numberOfQuestionOptions-1)
			fmt.Printf("main) s-addr:%p a-addr:%p\n", &wrongAnswers, wrongAnswers)
			getRandomWrongAnswers(&wrongAnswers, countries, correctAnswer)
			answerOptions := append(wrongAnswers, correctAnswer)
			shuffleSlice(answerOptions)

			quizFile.WriteString(fmt.Sprintf("Q%d: What is the capital of %s?\n", questionNum, countries[questionNum-1]))
			for i, option := range answerOptions {
				quizFile.WriteString(fmt.Sprintf(" %c. %s\n", 'A'+i, option))
			}
			quizFile.WriteString("\n")

			answerFile.WriteString(fmt.Sprintf("Q%d: %c\n", questionNum, 'A'+index(answerOptions, correctAnswer)))
		}
	}
}
