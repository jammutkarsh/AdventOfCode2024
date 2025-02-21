package day5

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jammutkarsh/AdventOfCode2024/internal"
)

func Solution() {
	pageOrders, printOrders, err := readPuzzleInput()
	if err != nil {
		internal.ExitErr(err)
	}
	sum := 0
	for _, printOrder := range printOrders { // [75 47 61 53 29]
		validOrder := true
		for i, pageNumber := range printOrder { // 75
			for j := i + 1; j < len(printOrder); j++ { // Use this loop to make pairs
				order := fmt.Sprintf("%d|%d", pageNumber, printOrder[j])
				if !pageOrders[order] {
					validOrder = false
					break
				}
			}
		}
		if validOrder {
			mid := len(printOrder) / 2
			sum += printOrder[mid]
		}
	}
	fmt.Println(sum)
}

func readPuzzleInput() (map[string]bool, [][]int, error) {
	data, err := os.ReadFile("inputs/5.txt")
	if err != nil {
		return nil, [][]int{}, err
	}
	pageOrder := make(map[string]bool)
	printOrder := [][]int{}
	changeSection := false
	for _, line := range strings.Split(string(data), "\n") {
		if line == "" {
			changeSection = true
			continue
		}

		if !changeSection { // Parsing Page Order
			pageOrder[line] = true
		} else { // Parsing Print Order
			arr := []int{}
			for _, num := range strings.Split(line, ",") {
				arr = append(arr, stringToInt(num))
			}
			printOrder = append(printOrder, arr)
		}
	}
	return pageOrder, printOrder, nil
}

func stringToInt(s string) int {
	val, _ := strconv.Atoi(s)
	return val
}
