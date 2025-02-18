package day3

import (
	"fmt"

	"github.com/jammutkarsh/AdventOfCode2024/internal"
)

/*
	mul() function is enabled by default, but
	there are additional functions like don't() and do() which enables and disables mul()

*/

const disableMul = "don't()"
const enableMul = "do()"

func SolutionPart2() {
	corruptedData, err := readPuzzleInput()
	if err != nil {
		internal.ExitErr(err)
	}
	// sum, instructionSize := 0, len(instruction)
	var (
		sum             = 0
		instructionSize = len(instruction)
		disableSize     = len(disableMul)
		enableSize      = len(enableMul)
	)

	mulEnabled := true // Default Condition

	for i := 0; i < len(corruptedData); i++ {

		if i+disableSize < len(corruptedData) && corruptedData[i:i+disableSize] == disableMul {
			mulEnabled = false
		} else if i+enableSize < len(corruptedData) && corruptedData[i:i+enableSize] == enableMul {
			mulEnabled = true
		}
		
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
			if validFormat != 0 && mulEnabled {
				sum += getMultiplier(corruptedData[i : validFormat+1])
			}
		}
	}
	fmt.Println(sum)
}
