package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
--- Day 2: Red-Nosed Reports ---
Fortunately, the first location The Historians want to search isn't a long walk from the Chief Historian's office.

While the Red-Nosed Reindeer nuclear fusion/fission plant appears to contain no sign of the Chief Historian, the engineers there run up to you as soon as they see you.
Apparently, they still talk about the time Rudolph was saved through molecular synthesis from a single electron.

They're quick to add that - since you're already here - they'd really appreciate your help analyzing some unusual data from the Red-Nosed reactor.
You turn to check if The Historians are waiting for you, but they seem to have already divided into groups that are currently searching every corner of the facility. You offer to help with the unusual data.

The unusual data (your puzzle input) consists of many reports, one report per line. Each report is a list of numbers called levels that are separated by spaces. For example:

7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9
This example data contains six reports each containing five levels.

The engineers are trying to figure out which reports are safe. The Red-Nosed reactor safety systems can only tolerate levels that are either gradually increasing or gradually decreasing.
So, a report only counts as safe if both of the following are true:

The levels are either all increasing or all decreasing.
Any two adjacent levels differ by at least one and at most three.
In the example above, the reports can be found safe or unsafe by checking those rules:

7 6 4 2 1: Safe because the levels are all decreasing by 1 or 2.
1 2 7 8 9: Unsafe because 2 7 is an increase of 5.
9 7 6 2 1: Unsafe because 6 2 is a decrease of 4.
1 3 2 4 5: Unsafe because 1 3 is increasing but 3 2 is decreasing.
8 6 4 4 1: Unsafe because 4 4 is neither an increase or a decrease.
1 3 6 7 9: Safe because the levels are all increasing by 1, 2, or 3.
So, in this example, 2 reports are safe.

Analyze the unusual data from the engineers. How many reports are safe?

---

The first step is to read the input data. The input data is a list of reports, where each report is a list of numbers (levels).
The reports are separated by new lines, and the level numbers are separated by spaces.
*/
func getReportsFromFile() []string {
	data, err := os.ReadFile("./2/input.txt")
	if err != nil {
		fmt.Println("Error reading file", err)
		os.Exit(1)
	}

	reports := strings.Split(string(data), "\n")
	return reports
}

func isReportSafe(report string) bool {
	levels := strings.Fields(report)
	if len(levels) == 0 {
		return false
	}

	// Check if the levels are increasing or decreasing
	// If they are not increasing or decreasing, the report is not safe
	// The rate of increase or decrease must be 1, 2, or 3
	// If the difference between any two adjacent levels is not 1, 2, or 3, the report is not safe
	increasing := false
	decreasing := false
	for i := 1; i < len(levels); i++ {
		current, err := strconv.Atoi(levels[i])
		if err != nil {
			fmt.Printf("Error converting level '%s': %v\n", levels[i], err)
		}

		previous, err := strconv.Atoi(levels[i-1])
		if err != nil {
			fmt.Printf("Error converting level '%s': %v\n", levels[i-1], err)
		}

		diff := int(math.Abs(float64(current - previous)))
		if diff < 1 || diff > 3 {
			return false
		}

		if current > previous {
			increasing = true
		}

		if current < previous {
			decreasing = true
		}

		if increasing && decreasing {
			return false
		}

	}

	return true
}

func isReportSafeWithSkip(levels []string, skipIndex int) bool {
	if len(levels) <= 1 {
		return false
	}

	increasing := false
	decreasing := false
	prevNum := -1
	for i := 0; i < len(levels); i++ {
		if i == skipIndex {
			continue
		}

		current, err := strconv.Atoi(levels[i])
		if err != nil {
			return false
		}

		if prevNum != -1 {
			diff := current - prevNum
			abdDiff := int(math.Abs(float64(diff)))
			if abdDiff < 1 || abdDiff > 3 {
				return false
			}

			if diff > 0 {
				increasing = true
			}

			if diff < 0 {
				decreasing = true
			}

			if increasing && decreasing {
				return false
			}
		}

		prevNum = current
	}
	return true
}

func isReportSafeWithDampener(report string) bool {
	levels := strings.Fields(report)
	if len(levels) == 0 {
		return false
	}

	if isReportSafeWithSkip(levels, -1) {
		return true
	}

	for i := 0; i < len(levels); i++ {
		if isReportSafeWithSkip(levels, i) {
			return true
		}
	}

	return false
}

func main() {
	reports := getReportsFromFile()
	fmt.Println("Number of Reports:", len(reports))

	safeReports := 0
	for _, report := range reports {
		if isReportSafe(report) {
			safeReports++
		}
	}

	fmt.Println("Number of Safe Reports:", safeReports)

	safeReports = 0
	for _, report := range reports {
		if isReportSafeWithDampener(report) {
			safeReports++
		}
	}

	fmt.Println("Number of Safe Reports with Dampener:", safeReports)

}

/*
--- Part Two ---
The engineers are surprised by the low number of safe reports until they realize they forgot to tell you about the Problem Dampener.

The Problem Dampener is a reactor-mounted module that lets the reactor safety systems tolerate a single bad level in what would otherwise be a safe report. It's like the bad level never happened!

Now, the same rules apply as before, except if removing a single level from an unsafe report would make it safe, the report instead counts as safe.

More of the above example's reports are now safe:

7 6 4 2 1: Safe without removing any level.
1 2 7 8 9: Unsafe regardless of which level is removed.
9 7 6 2 1: Unsafe regardless of which level is removed.
1 3 2 4 5: Safe by removing the second level, 3.
8 6 4 4 1: Safe by removing the third level, 4.
1 3 6 7 9: Safe without removing any level.
Thanks to the Problem Dampener, 4 reports are actually safe!

Update your analysis by handling situations where the Problem Dampener can remove a single level from unsafe reports. How many reports are now safe?
*/
