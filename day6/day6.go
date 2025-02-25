package day6

import (
	"fmt"
	"os"
	"strings"

	"github.com/jammutkarsh/AdventOfCode2024/internal"
)

func Solution() {
	box, err := readPuzzleInput()
	if err != nil {
		internal.ExitErr(err)
	}
	x, y := findGuard(box)
	move(box, x, y)
	fmt.Println(countMappedArea(box))
}

func readPuzzleInput() ([][]rune, error) {
	data, err := os.ReadFile("inputs/6.txt")
	if err != nil {
		return [][]rune{}, err
	}
	guardBox := [][]rune{}
	for _, row := range strings.Split(string(data), "\n") {
		runeRow := []rune{}
		for i := range row {
			runeRow = append(runeRow, rune(row[i]))
		}
		guardBox = append(guardBox, runeRow)
	}
	return guardBox, nil
}

func move(box [][]rune, x, y int) {
	coordinates := [][]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	count := 0
	i, j := coordinates[count%4][0], coordinates[count%4][1]
	for 0 <= x+i && x+i < len(box[x]) && 0 <= y+j && y+j < len(box) {
		i, j = coordinates[count%4][0], coordinates[count%4][1]
		if y+j < len(box) && x+i < len(box[x]) && box[y+j][x+i] == '#' {
			count++
		} else {
			box[y][x] = 'X'
			y += j
			x += i
		}
	}
	box[y][x] = 'X' // Mark the border
}

func findGuard(box [][]rune) (i, j int) {
	for y := range box {
		for x := range box[y] {
			if box[y][x] == '^' {
				return x, y // Returning as per the  Cartesian 2D plane
			}
		}
	}
	return
}

func countMappedArea(box [][]rune) int {
	count := 0
	for x := range box {
		for y := range box[x] {
			if box[x][y] == 'X' {
				count++
			}
		}
	}
	return count
}
