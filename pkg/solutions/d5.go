package solutions

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// import "regexp"

type Day5 struct {
	Day
	moveAllAtOnce bool
}

const (
	// Capture groups are: num of columns ;)
	columnNumRegex = `^\s\d.*(\d)$`
	// Capture groups are: all words inside square brackets
	columnValuesRegex = `[(\w)]`
	// Capture groups are: num to move, src column, dest column
	movementRegex = `^\w.*?(\d.*?) .*?(\d.*?) .*?(\d.*?)` //`^\w.*?(\d).*?(\d).*?(\d)`gm
)

func Day5Init(puzzle int) (day5 Day5) {
	day5.Puzzle = puzzle
	filepathFmt := "%s/pkg/inputs/day5/%s.txt"
	currDir, _ := os.Getwd()
	switch puzzle {
	case 0:
		day5.fileName = fmt.Sprintf(filepathFmt, currDir, "sample")
		day5.moveAllAtOnce = true
	case 2:
		day5.fileName = fmt.Sprintf(filepathFmt, currDir, "puzzle1")
		day5.moveAllAtOnce = true
	default:
		day5.fileName = fmt.Sprintf(filepathFmt, currDir, "puzzle1")
		day5.moveAllAtOnce = false
	}
	return
}
func (d Day5) Run() (total string) {
	inputBytes, _ := ioutil.ReadFile(d.fileName)
	slices := BuildSlices(inputBytes)
	fmt.Println(slices)
	moveInstructions := BuildMovements(inputBytes)
	finalDestSlices := MoveSliceContents(slices, moveInstructions, d.moveAllAtOnce)
	fmt.Println(finalDestSlices)
	finalString := []string{}
	for _, tmpSlice := range finalDestSlices {
		if len(tmpSlice)-1 < 0 {
			finalString = append(finalString, " ")
		} else {
			finalString = append(finalString, tmpSlice[len(tmpSlice)-1]) // srcSlice[len(slices[instruction["src"]])-1]
		}
	}
	fmt.Println(strings.Join(finalString, ""))
	return
}

func MoveSliceContents(slices map[int][]string, instructions []map[string]int, moveAllAtOnce bool) [][]string {
	for _, instruction := range instructions {
		// Move the items one at a time
		srcSlice := slices[instruction["src"]]
		numToMove := instruction["numToMove"]
		if moveAllAtOnce {
			itemsToMove := []string{}
			if len(slices[instruction["src"]])-numToMove > 0 {
				itemsToMove = srcSlice[len(slices[instruction["src"]])-numToMove:]
				slices[instruction["src"]] = srcSlice[:len(slices[instruction["src"]])-numToMove]
			} else {
				itemsToMove = srcSlice
				slices[instruction["src"]] = []string{}
			}
			// Move Item
			slices[instruction["dest"]] = append(slices[instruction["dest"]], itemsToMove...)
		} else {
			for numToMove := instruction["numToMove"]; numToMove > 0; numToMove-- {
				srcSlice := slices[instruction["src"]]
				if len(srcSlice)-1 < 0 {
					break
				}
				itemToMove := srcSlice[len(slices[instruction["src"]])-1]
				// Move Item
				slices[instruction["dest"]] = append(slices[instruction["dest"]], itemToMove)
				// Remove Item
				slices[instruction["src"]] = srcSlice[:len(slices[instruction["src"]])-1]
			}
		}
	}

	returnSlice := [][]string{}
	for i := 0; i < len(slices); i++ {
		returnSlice = append(returnSlice, slices[i+1])
	}
	return returnSlice
}
func BuildMovements(fileContents []byte) []map[string]int {
	// Capture groups are: num to move, src column, dest column
	// movementRegex = `^\w.*?(\d).*?(\d).*?(\d)`
	moveInstructions := []map[string]int{}
	mvRegex := regexp.MustCompile(movementRegex)
	splStrFileContents := strings.Split(string(fileContents), "\n")
	for _, line := range splStrFileContents {
		foundMovement := mvRegex.FindAllStringSubmatch(line, -1)
		if len(foundMovement) > 0 {
			// There's only 1 slice of strings
			tempMap := map[string]int{}
			for mvI, movement := range foundMovement[0] {
				if mvI == 0 {
					continue
				}
				if mvI == 1 {
					i, _ := strconv.Atoi(movement)
					tempMap["numToMove"] = i
				}
				if mvI == 2 {
					i, _ := strconv.Atoi(movement)
					tempMap["src"] = i
				}
				if mvI == 3 {
					i, _ := strconv.Atoi(movement)
					tempMap["dest"] = i
				}
			}
			// fmt.Println(fmt.Sprintf("| %s | move %d from %d to %d |", line, tempMap["numToMove"], tempMap["src"], tempMap["dest"]))
			moveInstructions = append(moveInstructions, tempMap)
		}
	}
	return moveInstructions
}

func BuildSlices(fileContents []byte) map[int][]string {
	colRegex := regexp.MustCompile(columnNumRegex)
	colValueRegex := regexp.MustCompile(columnValuesRegex)
	columnContents := map[int][]string{}
	splStrFileContents := strings.Split(string(fileContents), "\n")
	for lineIndex, line := range splStrFileContents {
		foundColumnNums := colRegex.FindAllStringSubmatch(line, 100)
		if len(foundColumnNums) >= 1 {
			maxCol, _ := strconv.Atoi(foundColumnNums[0][1])
			for tempI := lineIndex; tempI >= 0; tempI-- {
				if tempI == lineIndex {
					continue
				}
				fmt.Println(splStrFileContents[tempI])
				foundColumnVals := colValueRegex.FindAllStringSubmatch(splStrFileContents[tempI], maxCol+1)
				if len(splStrFileContents[tempI]) == 3 {
					if splStrFileContents[tempI][1:2] != " " {
						columnContents[1] = append(columnContents[1], splStrFileContents[tempI][1:2])
					}
				} else if len(foundColumnVals) != maxCol {
					tempCharCount := len(splStrFileContents[tempI])
					colsToFill := (tempCharCount % maxCol) + 1
					for tempCharCount > 0 {
						if splStrFileContents[tempI][tempCharCount-2:tempCharCount-1] != " " {
							columnContents[colsToFill] = append(columnContents[colsToFill], splStrFileContents[tempI][tempCharCount-2:tempCharCount-1])
						}
						tempCharCount = tempCharCount - 4
						colsToFill--
					}
				} else {
					for i, value := range foundColumnVals {
						columnContents[i+1] = append(columnContents[i+1], value[0])
					}
				}
			}
		}
	}
	return columnContents
}
