package day1

import (
	"fmt"
	"main/internal"
	"os"
	"sort"
	"strconv"
	"strings"
)


func Solution() {
	left, right, err := ReadPuzzelInput()
	if err != nil {
		internal.ExitErr(err)
	}
	if len(left) != len(right) {
		internal.ExitErr(fmt.Errorf("Right and Left sides are not equal"))
	}
	/*
		The task is to get the diff between nth smallest from left and right and add those diff.
		We can sort both sides and then add those.
		Add the absolute distance and return
	*/
	sort.Ints(left)
	sort.Ints(right)
	distance := 0
	for i := 0; i < len(left); i++ {
		diff := left[i] - right[i]
		if diff < 0 {
			distance = distance - diff
		} else {
			distance = distance + diff
		}
	}
	fmt.Println(distance)
}

func ReadPuzzelInput() (left, right []int, err error) {
	data, err := os.ReadFile("inputs/1.txt")
	if err != nil {
		return left, right, err
	}
	for _, v := range strings.Split(string(data), "\n") {
		leftInt, err := strconv.Atoi(strings.Split(v, "   ")[0]) // Copy pasted the space from 1.txt file.
		if err != nil {
			return []int{}, []int{}, err
		}
		rightInt, err := strconv.Atoi(strings.Split(v, "   ")[1])
		if err != nil {
			return []int{}, []int{}, err
		}

		left = append(left, leftInt)
		right = append(right, rightInt)
	}
	return
}
