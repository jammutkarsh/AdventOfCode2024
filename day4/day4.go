package day4

import (
	"fmt"
	"os"
	"strings"

	"github.com/jammutkarsh/AdventOfCode2024/internal"
)

func Solution() {
	matrix, err := readPuzzleInput()
	if err != nil {
		internal.ExitErr(err)
	}
	// These coordinates cover every possible direction where the word needs to be searched.
	// It eliminates the need to add special condition for reverse i.e XMAS -> SAMX
	searchCoordinates := getSearchCoordinates()
	var count int

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			for _, v := range searchCoordinates {
				if hasXMAS(i, j, v, matrix) {
					count++
				}
			}
		}
	}

	fmt.Println(count)
}

func hasXMAS(i, j int, coordinates []int, matrix []string) bool {
	x := coordinates[0]
	y := coordinates[1]
	for k, v := range "XMAS" {
		ii := i + k*x
		jj := j + k*y
		if !(0 <= ii && ii < len(matrix) && 0 <= jj && jj < len(matrix)) { // Should be within the grid
			return false
		}
		if matrix[ii][jj] != byte(v) {
			return false
		}
	}
	return true
}

func getSearchCoordinates() (coordinates [][]int) {
	for _, v := range []int{-1, 0, 1} {
		for _, b := range []int{-1, 0, 1} {
			if b == 0 && v == 0 { // The position of element itself
				continue
			}
			coordinates = append(coordinates, []int{v, b})
		}
	}
	/*
	[-1, 1]      [0, 1]       [1, 1]
	[-1, 0]    Element   [1, 0]
	[-1, -1]    [0, -1]     [1, -1]
	*/
	return
}

func readPuzzleInput() (matrix []string, err error) {
	data, err := os.ReadFile("inputs/4.txt")
	if err != nil {
		return []string{}, err
	}
	return strings.Split(string(data), "\n"), nil
}
