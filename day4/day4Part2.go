package day4

import (
	"fmt"
	"strings"

	"github.com/jammutkarsh/AdventOfCode2024/internal"
)

func SolutionPart2() {
	matrix, err := readPuzzleInput()
	if err != nil {
		internal.ExitErr(err)
	}
	var count int

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if hasMAS(i, j, matrix) {
				count++
			}
		}
	}

	fmt.Println(count)
}

func hasMAS(i, j int, matrix []string) bool {
	if !(0 < i && i < len(matrix)-1 && 0 < j && j < len(matrix)-1) { // Should be within the grid
		return false
	}
	/*
		M . S
		.  A  .
		M . S
	*/
	if matrix[i][j] != 'A' { // centre should be 'A'
		return false
	}

	// Diagonally the char should be M or S
	d1 := string(matrix[i+1][j+1]) + string(matrix[i-1][j-1])
	d2 := string(matrix[i+1][j-1]) + string(matrix[i-1][j+1])

	// Checking if d1,d2 is either MS or SM
	return strings.Contains("MSM", d1) && strings.Contains("MSM", d2)
}
