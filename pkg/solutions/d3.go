package solutions

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Day3 struct {
	Day
	itemPriorities map[string]int
}

var itemsSlice = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

func Day3Init(puzzle int) Day3 {
	day3 := Day3{itemPriorities: map[string]int{}}
	day3.Puzzle = puzzle
	filepathFmt := "%s/pkg/inputs/day3/%s.txt"
	currDir, _ := os.Getwd()

	switch puzzle {
	case 0:
		day3.fileName = fmt.Sprintf(filepathFmt, currDir, "sample")
	default:
		day3.fileName = fmt.Sprintf(filepathFmt, currDir, "puzzle1")
	}
	for i, k := range itemsSlice {
		day3.itemPriorities[k] = i + 1
	}
	return day3
}
func (d Day3) Run() string {
	fmt.Println(fmt.Sprintf("Puzzle: %d", d.Puzzle))
	var inputBytes []byte
	runningTotal := 0
	switch d.Puzzle {
	case 2:
		inputBytes, _ = ioutil.ReadFile(d.fileName)
		runningTotal = d.GetGroupCodePriority(inputBytes)
	default:
		inputBytes, _ = ioutil.ReadFile(d.fileName)
		for _, line := range strings.Split(string(inputBytes), "\n") {
			runningTotal = runningTotal + d.FindItemThatIsInBothCompartments(line)
		}
	}
	return strconv.Itoa(runningTotal)
}
func (d Day3) FindItemThatIsInBothCompartments(ruckContents string) (priorityTotal int) {
	compartment1 := ruckContents[0 : len(ruckContents)/2]
	compartment2 := ruckContents[len(ruckContents)/2:]
	priorityTotal = 0
	found := map[string]int{}
	for _, a := range compartment1 {
		for _, b := range compartment2 {
			if string(a) == string(b) {
				if _, ok := found[string(a)]; !ok {
					found[string(a)] = 0
					priorityTotal = priorityTotal + d.itemPriorities[string(a)]
				}
			}
		}
	}
	return
}

func (d Day3) GetGroupCodePriority(inputBytes []byte) int {
	priorityTotal := 0
	groups := []map[string][]int{}
	tempMap := map[string][]int{}
	for i, line := range strings.Split(string(inputBytes), "\n") {
		currentLineMapEntry := i % 3
		for _, char := range line {
			if _, ok := tempMap[string(char)]; ok {
				for l, s := range tempMap[string(char)] {
					if s == currentLineMapEntry {
						break
					} else {
						if l == len(tempMap[string(char)])-1 {
							tempMap[string(char)] = append(tempMap[string(char)], currentLineMapEntry)
						}
					}
				}
			} else {
				tempMap[string(char)] = []int{currentLineMapEntry}
			}
		}
		if currentLineMapEntry == 2 {
			groups = append(groups, tempMap)
			tempMap = map[string][]int{}
		}
	}
	for _, tempMap := range groups {
		for k, v := range tempMap {
			if len(v) > 2 {
				priorityTotal = priorityTotal + d.itemPriorities[k]
			}
		}
	}
	return priorityTotal
}
