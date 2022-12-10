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

type Day1 struct {
	Puzzle int
}

func (d Day1) Run() int {
	fmt.Println(fmt.Sprintf("Puzzle: %d", d.Puzzle))
	currDir, _ := os.Getwd()
	var inputBytes []byte
	filepathFmt := "%s/pkg/inputs/day1/%s.txt"
	total := 0
	switch puzzle := d.Puzzle; puzzle {
	case 1:
		fileName := fmt.Sprintf(filepathFmt, currDir, "puzzle1")
		inputBytes, _ = ioutil.ReadFile(fileName)
		caloriesSlice := getTotalCaloriesPerElf(inputBytes)
		total = caloriesSlice[len(caloriesSlice)-1]
	case 2:
		fileName := fmt.Sprintf(filepathFmt, currDir, "puzzle1")
		fmt.Println(fileName)
		inputBytes, _ = ioutil.ReadFile(fileName)
		// Need to sum the top 3
		caloriesSlice := getTotalCaloriesPerElf(inputBytes)
		total = caloriesSlice[len(caloriesSlice)-1] + caloriesSlice[len(caloriesSlice)-2] + caloriesSlice[len(caloriesSlice)-3]
	default:
		fileName := fmt.Sprintf(filepathFmt, currDir, "sample")
		inputBytes, _ = ioutil.ReadFile(fileName)
		caloriesSlice := getTotalCaloriesPerElf(inputBytes)
		total = caloriesSlice[len(caloriesSlice)-1]
	}
	return total
	// fmt.Println(input)
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
			fmt.Println("Got non int value")
			break
		}
		sliceLen := len(caloriesSlice)
		fmt.Println(fmt.Sprintf("slice length: %d elf: %d", sliceLen, elf))
		if sliceLen == elf {
			fmt.Println("adding new elf matching length and elf+1")
			fmt.Println(elf)
			caloriesSlice = append(caloriesSlice, calorie)
			continue
		}
		fmt.Println(fmt.Sprintf("Elf: %d, His current calories: %d, adding %d calories", elf, caloriesSlice[elf], calorie))
		caloriesSlice[elf] = caloriesSlice[elf] + calorie
	}
	sort.Ints(caloriesSlice)
	return caloriesSlice
}
