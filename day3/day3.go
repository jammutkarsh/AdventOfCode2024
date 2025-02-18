package day3

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jammutkarsh/AdventOfCode2024/internal"
)

// The regex pattern that we are looking for is mul(*,*), but 'mul(' is common across all.
// So we will search for that.
const instruction = "mul("

func Solution() {
	corruptedData, err := readPuzzleInput()
	if err != nil {
		internal.ExitErr(err)
	}
	sum, instructionSize := 0, len(instruction)

	for i := 0; i < len(corruptedData); i++ {
		if i+instructionSize < len(corruptedData) && corruptedData[i:i+instructionSize] == instruction {
			validFormat := 0
		extractValidSequence:
			for j := i + instructionSize; /*'mul(' is identified */ j < len(corruptedData); j++ {
				switch {
				case 0 <= corruptedData[j]-'0' && corruptedData[j]-'0' <= 9:
					continue
				case corruptedData[j] == ',':
					continue
				case corruptedData[j] == ')':
					validFormat = j
					break extractValidSequence // Search Complete
				default:
					break extractValidSequence // Doesn't follow the regex format: "mul(*,*)"
				}
			}
			if validFormat != 0 {
				sum += getMultiplier(corruptedData[i : validFormat+1])
			}
		}
	}
	fmt.Println(sum)
}

func getMultiplier(s string) int {
	s = strings.TrimPrefix(s, instruction)
	s = strings.TrimSuffix(s, ")")

	num1, err := strconv.Atoi(strings.Split(s, ",")[0])
	if err != nil {
		return 0
	}
	num2, err := strconv.Atoi(strings.Split(s, ",")[1])
	if err != nil {
		return 0
	}

	return num1 * num2
}

func readPuzzleInput() (corruptedData string, err error) {
	data, err := os.ReadFile("inputs/3.txt")
	if err != nil {
		return "", err
	}
	return string(data), nil
}
