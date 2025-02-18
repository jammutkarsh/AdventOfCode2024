package main

import (
	"fmt"
	"os"

	"github.com/jammutkarsh/AdventOfCode2024/day1"
	"github.com/jammutkarsh/AdventOfCode2024/day2"
	"github.com/jammutkarsh/AdventOfCode2024/day3"
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
		day3.Solution()
	case "4":
		panic("Implement Me!!!")
	case "5":
		panic("Implement Me!!!")
	case "6":
		panic("Implement Me!!!")
	case "7":
		panic("Implement Me!!!")
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
