package main

import (
	"andrewbernier42/quizgame/questions"
	"fmt"
	"os"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {

	a := app.New()
	w := a.NewWindow("Hello Fyne")

	label := widget.NewLabel("Hello, World!")
	w.SetContent(container.NewVBox(label))

	w.ShowAndRun()
	problems, err := questions.ReadFile("problems.csv")
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}

	operation := questions.SelectOperation()

	correct := 0
	//for i, p := range problems {
	for i := range problems {
		if questions.Present(i, questions.GetProblem(operation, 999)) {
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
