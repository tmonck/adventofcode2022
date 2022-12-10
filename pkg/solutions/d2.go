package solutions

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Day2 struct {
	Puzzle               int
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
	day2 := Day2{Puzzle: puzzle}
	switch puzzle {
	case 2:
		fmt.Println("case2")
		var Rock = Choice{name: "Rock", score: 1, beats: "Scissors", code: []string{"A"}}
		var Paper = Choice{name: "Paper", score: 2, beats: "Rock", code: []string{"B"}}
		var Scissors = Choice{name: "Scissors", score: 3, beats: "Paper", code: []string{"C"}}
		day2.choices = map[string]Choice{}
		day2.choices[Rock.name] = Rock
		day2.choices[Paper.name] = Paper
		day2.choices[Scissors.name] = Scissors
		day2.results = map[string]Result{}
		day2.results["X"] = Result{code: "X", score: 0}
		day2.results["Y"] = Result{code: "Y", score: 3}
		day2.results["Z"] = Result{code: "Z", score: 6}
		day2.secondColumnIsChoice = false
	default:
		fmt.Println("default")
		var Rock = Choice{name: "Rock", score: 1, beats: "Scissors", code: []string{"A", "X"}}
		var Paper = Choice{name: "Paper", score: 2, beats: "Rock", code: []string{"B", "Y"}}
		var Scissors = Choice{name: "Scissors", score: 3, beats: "Paper", code: []string{"C", "Z"}}
		day2.choices = map[string]Choice{}
		day2.choices[Rock.name] = Rock
		day2.choices[Paper.name] = Paper
		day2.choices[Scissors.name] = Scissors
		day2.results = map[string]Result{}
		day2.secondColumnIsChoice = true
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
	currDir, _ := os.Getwd()
	var inputBytes []byte
	filepathFmt := "%s/pkg/inputs/day2/%s.txt"
	totalScore := 0
	switch puzzle := d.Puzzle; puzzle {
	case 0:
		fileName := fmt.Sprintf(filepathFmt, currDir, "sample")
		inputBytes, _ = ioutil.ReadFile(fileName)
	default:
		fileName := fmt.Sprintf(filepathFmt, currDir, "puzzle1")
		inputBytes, _ = ioutil.ReadFile(fileName)
	}
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
