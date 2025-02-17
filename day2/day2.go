package day2

import (
	"fmt"
	"main/internal"
	"os"
	"strconv"
	"strings"
)

func Solution() {
	readings, err := ReadPuzzleInput()
	if err != nil {
		internal.ExitErr(err)
	}
	safeReadings := 0
	for _, reading := range readings {
		if isSafe(reading) {
			safeReadings++
		}
	}
	fmt.Println(safeReadings)
}

func isSafe(reading []int) bool {
	increasing := true

	if reading[1]-reading[0] == 0 {
		return false // Non increasing
	} else if reading[1]-reading[0] < 0 {
		increasing = false
	}

	for i := 1; i < len(reading); i++ {
		diff := reading[i] - reading[i-1]
		if increasing && (0 < diff && diff < 4) { // Increasing
			continue
		} else if !increasing && (0 < -diff && -diff < 4) { // Decreasing
			continue
		} else {
			return false
		}
	}
	return true
}

func ReadPuzzleInput() (readings [][]int, err error) {
	data, err := os.ReadFile("inputs/2.txt")
	if err != nil {
		return [][]int{}, err
	}
	for _, v := range strings.Split(string(data), "\n") {
		var levels []int
		for _, val := range strings.Split(v, " ") {
			level, err := strconv.Atoi(val)
			if err != nil {
				return [][]int{}, err
			} else {
				levels = append(levels, level)
			}
		}
		readings = append(readings, levels)
	}
	return
}
