package questions

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Problem struct {
	Q string
	A string
}

func ReadFile(filename string) ([]Problem, error) {
	csvFilename := flag.String("csv", filename, "a csv file in the format of 'question,answer'")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}
	problems := parseLines(lines)
	return problems, nil
}

func parseLines(lines [][]string) []Problem {
	ret := make([]Problem, len(lines))
	for i, line := range lines {
		ret[i] = Problem{
			Q: line[0],
			A: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func Present(i int, p Problem) bool {
	fmt.Printf("Problem #%d: %s = \n", i+1, p.Q)
	var answer string
	fmt.Scanf("%s\n", &answer)
	return answer == p.A
}
