package questions

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func GetProblem(operation string, limit int) Problem {
	rand.Seed(time.Now().UnixNano()) // Seed to ensure different results

	num1 := rand.Intn(limit) // Generate a number between 0 and 999
	num2 := rand.Intn(limit)
	switch operation {
	case "addition":
		return CreateAdditionProblem(num1, num2)
	case "subtraction":
		return CreateSubtractionProblem(num1, num2)
	case "multiplication":
		return CreateMultiplicationProblem(num1, num2)
	case "division":
		if num2 == 0 {
			fmt.Println("Error: Division by zero is not allowed")
			os.Exit(1)
		}
		return CreateDivisionProblem(num1, num2)
	default:
		return Problem{}
	}

}

func CreateAdditionProblem(num1 int, num2 int) Problem {
	question := fmt.Sprintf("%d + %d = ?", num1, num2)
	answer := fmt.Sprintf("%d", num1+num2)

	return Problem{
		Q: question,
		A: answer,
	}
}

func CreateSubtractionProblem(num1 int, num2 int) Problem {
	question := fmt.Sprintf("%d - %d = ?", num1, num2)
	answer := fmt.Sprintf("%d", num1-num2)

	return Problem{
		Q: question,
		A: answer,
	}
}

func CreateMultiplicationProblem(num1 int, num2 int) Problem {
	question := fmt.Sprintf("%d * %d = ?", num1, num2)
	answer := fmt.Sprintf("%d", num1*num2)

	return Problem{
		Q: question,
		A: answer,
	}
}

func CreateDivisionProblem(num1 int, num2 int) Problem {
	question := fmt.Sprintf("%d / %d = ?", num1, num2)
	answer := fmt.Sprintf("%d", num1/num2)

	return Problem{
		Q: question,
		A: answer,
	}
}
