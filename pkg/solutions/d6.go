package solutions

import (
	"fmt"
	"os"
)

func Day6Init(puzzle int) (day6 Day5) {
	day6.Puzzle = puzzle
	filepathFmt := "%s/pkg/inputs/day5/%s.txt"
	currDir, _ := os.Getwd()
	switch puzzle {
	case 0:
		day6.fileName = fmt.Sprintf(filepathFmt, currDir, "sample")
		day6.moveAllAtOnce = true
	case 2:
		day6.fileName = fmt.Sprintf(filepathFmt, currDir, "puzzle1")
		day6.moveAllAtOnce = true
	default:
		day6.fileName = fmt.Sprintf(filepathFmt, currDir, "puzzle1")
		day6.moveAllAtOnce = false
	}
	return
}
