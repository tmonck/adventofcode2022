package solutions

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Day2 struct {
	Day
	choices              Choices
	results              Results
	secondColumnIsChoice bool
}

type Choice struct {
	score int
	name  string
	beats string
	code  []string
}

func Day2Init(puzzle int) Day2 {
	day2 := Day2{}
	day2.Puzzle = puzzle
	day2.choices = map[string]Choice{}
	day2.results = map[string]Result{}
	day2.secondColumnIsChoice = true
	currDir, _ := os.Getwd()
	filepathFmt := "%s/pkg/inputs/day2/%s.txt"

	Rock := Choice{name: "Rock", score: 1, beats: "Scissors", code: []string{"A", "X"}}
	Paper := Choice{name: "Paper", score: 2, beats: "Rock", code: []string{"B", "Y"}}
	Scissors := Choice{name: "Scissors", score: 3, beats: "Paper", code: []string{"C", "Z"}}
	switch puzzle {
	case 1:
		day2.choices[Rock.name] = Rock
		day2.choices[Paper.name] = Paper
		day2.choices[Scissors.name] = Scissors
		day2.fileName = fmt.Sprintf(filepathFmt, currDir, "puzzle1")
	case 2:
		Rock.code = []string{"A"}
		Paper.code = []string{"B"}
		Scissors.code = []string{"C"}
		day2.choices[Rock.name] = Rock
		day2.choices[Paper.name] = Paper
		day2.choices[Scissors.name] = Scissors
		day2.results["X"] = Result{code: "X", score: 0}
		day2.results["Y"] = Result{code: "Y", score: 3}
		day2.results["Z"] = Result{code: "Z", score: 6}
		day2.secondColumnIsChoice = false
		day2.fileName = fmt.Sprintf(filepathFmt, currDir, "puzzle1")
	default:
		day2.choices[Rock.name] = Rock
		day2.choices[Paper.name] = Paper
		day2.choices[Scissors.name] = Scissors
		day2.fileName = fmt.Sprintf(filepathFmt, currDir, "sample")
	}
	return day2
}

var Rock = Choice{name: "Rock", score: 1, beats: "Scissors", code: []string{"A", "X"}}
var Paper = Choice{name: "Paper", score: 2, beats: "Rock", code: []string{"B", "Y"}}
var Scissors = Choice{name: "Scissors", score: 3, beats: "Paper", code: []string{"C", "Z"}}

type Choices map[string]Choice

var choices = Choices{}

type Result struct {
	score int
	code  string
}

type Results map[string]Result

var results = Results{}

func (d Day2) Run() int {
	fmt.Println(fmt.Sprintf("Puzzle: %d", d.Puzzle))
	var inputBytes []byte
	totalScore := 0
	inputBytes, _ = ioutil.ReadFile(d.fileName)
	totalScore = d.CalculateScore(inputBytes, d.secondColumnIsChoice)
	return totalScore
}

func (m Choices) ConvertToChoice(s string) Choice {
	for _, c := range m {
		for _, v := range c.code {
			if v == s {
				return c
			}
		}
	}
	return Choice{}
}

func (m Results) ConvertResultToChoice(s string, opponentChoice Choice, choices Choices) Choice {
	for _, c := range m {
		if s == c.code {
			switch c.score {
			case 0:
				return choices[opponentChoice.beats]
			case 3:
				return opponentChoice
			default:
				for _, choice := range choices {
					if choice.beats == opponentChoice.name {
						return choice
					}
				}
			}
		}
	}
	return Choice{}
}
func (d Day2) CalculateScore(inputBytes []byte, secondColumnIsChoice bool) (totalScore int) {
	totalScore = 0
	for _, line := range strings.Split(string(inputBytes), "\n") {
		if len(line) == 0 {
			continue
		}
		var opponentChoice Choice
		var myChoice Choice
		for x, choice := range strings.Split(line, " ") {
			if x == 0 {
				opponentChoice = d.choices.ConvertToChoice(choice)
				continue
			}
			if secondColumnIsChoice {
				myChoice = d.choices.ConvertToChoice(choice)
			} else {
				myChoice = d.results.ConvertResultToChoice(choice, opponentChoice, d.choices)
			}
			totalScore = totalScore + myChoice.score
		}

		if myChoice.beats == opponentChoice.name {
			totalScore = totalScore + 6
		}
		if myChoice.name == opponentChoice.name {
			totalScore = totalScore + 3
		}
	}
	return
}
