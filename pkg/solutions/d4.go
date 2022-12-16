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

type Sections []Section
type Section []int

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
	inputBytes, _ = ioutil.ReadFile(d.fileName)
	overlappingSectionsCount := 0
	for _, line := range strings.Split(string(inputBytes), "\n") {
		if len(line) == 0 {
			continue
		}
		sections := d.ParseSections(line)
		switch d.Puzzle {
		case 2:
			if sections.HasAnyOverlap() {
				overlappingSectionsCount++
			}
		default:
			if sections.HasCompleteOverlap() {
				overlappingSectionsCount++
			}
		}
	}
	return strconv.Itoa(overlappingSectionsCount)
}

func (d Day4) ParseSections(line string) (results Sections) {
	sections := strings.Split(line, ",")
	for _, section := range sections {
		tempStrSlice := strings.Split(section, "-")
		tempIntSlice := []int{}
		for _, i := range tempStrSlice {
			tempInt, _ := strconv.Atoi(i)
			tempIntSlice = append(tempIntSlice, tempInt)
		}
		results = append(results, tempIntSlice)
	}
	return
}

func (s Sections) HasAnyOverlap() bool {
	if s[0][0] >= s[1][0] && s[0][1] <= s[1][1] ||
		s[0][0] <= s[1][0] && s[0][1] >= s[1][1] {
		fmt.Println("first block")
		return true
	}

	// a >= c && a <= d || b >= c && b <= d
	if s[0][0] >= s[1][0] && s[0][0] <= s[1][1] ||
		s[0][1] >= s[1][0] && s[0][1] <= s[1][1] {
		fmt.Println("second block")
		return true
	}
	return false
}

func (s Sections) HasCompleteOverlap() bool {
	if s[0][0] >= s[1][0] && s[0][1] <= s[1][1] ||
		s[0][0] <= s[1][0] && s[0][1] >= s[1][1] {
		return true
	}
	return false
}
