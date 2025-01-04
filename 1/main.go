package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
)

// advent of code day one
func main() {

	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	fmt.Println("Contents of file:", string(data))

	leftList := []int{}
	rightList := []int{}

	for _, line := range strings.Split(string(data), "\n") {
		if line == "" {
			continue
		}
		parts := strings.Split(line, " ")

		left, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error converting left number", err)
			return
		}
		right, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println("Error converting right number", err)
			return
		}
		leftList = append(leftList, left)
		rightList = append(rightList, right)
	}

	sort.Ints(leftList)
	sort.Ints(rightList)

	distance := 0
	for i := 0; i < len(leftList); i++ {
		distance += abs(leftList[i] - rightList[i])
	}

	fmt.Println("Distance:", distance)
}
