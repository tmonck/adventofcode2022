package solutions

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Day4 struct {
	Day
}

func Day4Init(puzzle int) Day4 {
	day4 := Day4{}
	day4.Puzzle = puzzle
	filepathFmt := "%s/pkg/inputs/day4/%s.txt"
	currDir, _ := os.Getwd()
	switch puzzle {
	case 0:
		day4.fileName = fmt.Sprintf(filepathFmt, currDir, "sample")
	default:
		day4.fileName = fmt.Sprintf(filepathFmt, currDir, "puzzle1")
	}
	return day4
}

func (d Day4) Run() string {
	inputBytes := []byte{}
	overlappingSectionsCount := 0
	switch d.Puzzle {
	case 2:
		inputBytes, _ = ioutil.ReadFile(d.fileName)
		for _, line := range strings.Split(string(inputBytes), "\n") {
			if len(line) == 0 {
				continue
			}
			if d.DoesAnySectionOverlap(line) {
				overlappingSectionsCount++
			}
		}
	default:
		inputBytes, _ = ioutil.ReadFile(d.fileName)
		for _, line := range strings.Split(string(inputBytes), "\n") {
			if len(line) == 0 {
				continue
			}
			if d.IsOneSectionCompleteOverlap(line) {
				overlappingSectionsCount++
			}
		}
	}
	return strconv.Itoa(overlappingSectionsCount)
}

func (d Day4) IsOneSectionCompleteOverlap(line string) bool {
	sections := strings.Split(line, ",")
	results := [][]int{}
	for _, section := range sections {
		tempStrSlice := strings.Split(section, "-")
		tempIntSlice := []int{}
		for _, i := range tempStrSlice {
			tempInt, _ := strconv.Atoi(i)
			tempIntSlice = append(tempIntSlice, tempInt)
		}
		results = append(results, tempIntSlice)
	}
	for i, result := range results {
		if i == 1 {
			continue
		}
		for a := range result {
			if a == 1 {
				continue
			}
			if result[a] >= results[i+1][a] && result[a+1] <= results[i+1][a+1] {
				return true
			}
			if result[a] <= results[i+1][a] && result[a+1] >= results[i+1][a+1] {
				return true
			}
			return false
		}
	}
	return false
}

func (d Day4) DoesAnySectionOverlap(line string) bool {
	sections := strings.Split(line, ",")
	results := [][]int{}
	for _, section := range sections {
		tempStrSlice := strings.Split(section, "-")
		tempIntSlice := []int{}
		for _, i := range tempStrSlice {
			tempInt, _ := strconv.Atoi(i)
			tempIntSlice = append(tempIntSlice, tempInt)
		}
		results = append(results, tempIntSlice)
	}
	for i, result := range results {
		if i == 1 {
			continue
		}
		for a := range result {
			if a == 1 {
				continue
			}
			//  i          i+1
			// [a, b] [c, d]
			// 36-57,7-36
			b := a + 1
			c := a
			d := b
			if result[a] >= results[i+1][a] && result[a+1] <= results[i+1][a+1] ||
				result[a] <= results[i+1][a] && result[a+1] >= results[i+1][a+1] {
				return true
			}

			if result[a] >= results[i+1][c] && result[a] <= results[i+1][d] ||
				result[b] >= results[i+1][c] && result[b] <= results[i+1][d] {
				return true
			}
			return false
		}
	}
	return false
}
