package solutions

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Solution interface {
	Run() int
}

type Day struct {
	Puzzle   int
	fileName string
}
type Day1 struct {
	Day
}

func Day1Init(puzzle int) (day1 Day1) {
	day1.Puzzle = puzzle
	currDir, _ := os.Getwd()
	filepathFmt := "%s/pkg/inputs/day1/%s.txt"
	switch puzzle {
	case 0:
		day1.fileName = fmt.Sprintf(filepathFmt, currDir, "sample")
	default:
		day1.fileName = fmt.Sprintf(filepathFmt, currDir, "puzzle1")
	}
	return
}

func (d Day1) Run() int {
	fmt.Println(fmt.Sprintf("Puzzle: %d", d.Puzzle))
	var inputBytes []byte
	inputBytes, _ = ioutil.ReadFile(d.fileName)
	caloriesSlice := getTotalCaloriesPerElf(inputBytes)
	total := 0
	switch puzzle := d.Puzzle; puzzle {
	case 2:
		total = caloriesSlice[len(caloriesSlice)-1] + caloriesSlice[len(caloriesSlice)-2] + caloriesSlice[len(caloriesSlice)-3]
	default:
		total = caloriesSlice[len(caloriesSlice)-1]
	}
	return total
}

func getTotalCaloriesPerElf(inputBytes []byte) []int {
	input := string(inputBytes)
	stringSlice := strings.Split(input, "\n")
	elf := 0
	caloriesSlice := []int{}
	for _, mystring := range stringSlice {
		if mystring == "" {
			elf = elf + 1
			continue
		}
		calorie, err := strconv.Atoi(mystring)
		if err != nil {
			break
		}
		sliceLen := len(caloriesSlice)
		if sliceLen == elf {
			caloriesSlice = append(caloriesSlice, calorie)
			continue
		}
		caloriesSlice[elf] = caloriesSlice[elf] + calorie
	}
	sort.Ints(caloriesSlice)
	return caloriesSlice
}
