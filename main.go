package main

import (
	"andrewbernier42/quizgame/questions"
	"fmt"
	"os"
)

func main() {

	problems, err := questions.ReadFile("problems.csv")
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}
	correct := 0
	//for i, p := range problems {
	for i := range problems {
		if questions.Present(i, questions.GetProblem("subtraction", 999)) {
			correct++
		}
		/*if questions.Present(i, p) {
			correct++
		}*/
	}

	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
