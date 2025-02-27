package main

import (
	"fmt"
	"os"

	"github.com/jammutkarsh/AdventOfCode2024/day1"
	"github.com/jammutkarsh/AdventOfCode2024/day2"
	"github.com/jammutkarsh/AdventOfCode2024/day3"
	"github.com/jammutkarsh/AdventOfCode2024/day4"
	"github.com/jammutkarsh/AdventOfCode2024/day5"
	"github.com/jammutkarsh/AdventOfCode2024/day6"
	"github.com/jammutkarsh/AdventOfCode2024/day7"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Day Argument Missing\nUsage: \n\t%s 1\n", os.Args[0])
		os.Exit(1)
	}
	switch os.Args[1] {
	case "1":
		day1.Solution()
	case "2":
		day2.Solution()
	case "3":
		fmt.Println("Day 3 Part 1")
		day3.Solution()
		fmt.Println("Day 3 Part 2")
		day3.SolutionPart2()
	case "4":
		fmt.Println("Day 4 Part 1")
		day4.Solution()
		fmt.Println("Day 4 Part 2")
		day4.SolutionPart2()
	case "5":
		fmt.Println("Day 5 Part 1")
		day5.Solution()
		fmt.Println("Day 5 Part 2")
		day5.SolutionPart2()
	case "6":
		day6.Solution()
	case "7":
		day7.Solution()
	case "8":
		panic("Implement Me!!!")
	case "9":
		panic("Implement Me!!!")
	case "10":
		panic("Implement Me!!!")
	case "11":
		panic("Implement Me!!!")
	case "12":
		panic("Implement Me!!!")
	case "13":
		panic("Implement Me!!!")
	case "14":
		panic("Implement Me!!!")
	case "15":
		panic("Implement Me!!!")
	case "16":
		panic("Implement Me!!!")
	case "17":
		panic("Implement Me!!!")
	case "18":
		panic("Implement Me!!!")
	case "19":
		panic("Implement Me!!!")
	case "20":
		panic("Implement Me!!!")
	case "21":
		panic("Implement Me!!!")
	case "22":
		panic("Implement Me!!!")
	case "23":
		panic("Implement Me!!!")
	case "24":
		panic("Implement Me!!!")
	case "25":
		panic("Implement Me!!!")
	}
}
