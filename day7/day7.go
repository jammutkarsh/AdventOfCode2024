package day7

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jammutkarsh/AdventOfCode2024/internal"
)

func Solution() {
	puzzle, err := readPuzzleInput()
	if err != nil {
		internal.ExitErr(err)
	}
	var sum int = 0
	for result, arr := range puzzle {
		if !recursive(arr[0], result, arr, 1) { // i is 1 because we are already passing arr[0]
			delete(puzzle, result)
		} else {
			sum += result
		}
	}
	fmt.Println(sum)
}

func recursive(currentVal, maxVal int, arr []int, i int) bool {
	if i == len(arr) { // All numbers should be used
		return currentVal == maxVal
	}
	// Try Multiplication
	if recursive(currentVal*arr[i], maxVal, arr, i+1) {
		return true
	}
	// Try Addition
	return recursive(currentVal+arr[i], maxVal, arr, i+1)
}

func readPuzzleInput() (map[int][]int, error) {
	data, err := os.ReadFile("inputs/7.txt")
	if err != nil {
		return nil, err
	}
	puzzle := make(map[int][]int)
	for line := range strings.SplitSeq(string(data), "\n") {
		parts := strings.Split(line, ":")
		numbers := []int{}
		for num := range strings.SplitSeq(strings.TrimSpace(parts[1]), " ") {
			numbers = append(numbers, stringToInt(num))
		}
		puzzle[stringToInt(parts[0])] = numbers
	}
	return puzzle, err
}

func stringToInt(s string) int {
	val, _ := strconv.Atoi(s)
	return int(val)
}
