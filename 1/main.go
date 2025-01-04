package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

// advent of code day one
func main() {
	data, err := os.ReadFile("./1/input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	leftList := []int{}
	rightList := []int{}

	for _, line := range strings.Split(string(data), "\n") {
		// Skip empty lines
		if line == "" {
			continue
		}

		// Split the line and check if we have exactly 2 parts
		parts := strings.Fields(line) // Using Fields instead of Split to handle multiple spaces
		if len(parts) != 2 {
			fmt.Printf("Invalid line format: %s\n", line)
			continue
		}

		left, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Printf("Error converting left number '%s': %v\n", parts[0], err)
			continue
		}

		right, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Printf("Error converting right number '%s': %v\n", parts[1], err)
			continue
		}

		leftList = append(leftList, left)
		rightList = append(rightList, right)
	}

	slices.Sort(leftList)
	slices.Sort(rightList)

	distance := 0
	for i := 0; i < len(leftList); i++ {
		diff := float64(leftList[i] - rightList[i])
		distance += int(math.Abs(diff))
	}

	fmt.Println("Distance:", distance)
}
