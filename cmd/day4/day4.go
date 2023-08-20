package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current wd")
		os.Exit(1)
	}

	data, err := os.ReadFile(filepath.Join(cwd, "inputs", "day4.txt"))
	if err != nil {
		fmt.Println("Error reading file")
		os.Exit(2)
	}
	sum1 := 0
	sum2 := 0
	lines := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")
	for _, line := range lines {
		if len(line) == 0 {
			break
		}
		tasks := strings.Split(line, ",")
		first := createSubInterval(tasks[0])
		second := createSubInterval(tasks[1])
		if isSubInterval(first, second) {
			sum1++
		}
		if doesOverlap(first, second) {
			sum2++
		}
	}
	fmt.Println(sum1)
	fmt.Println(sum2)
}

func createSubInterval(intervalAsString string) []int {
	asStrings := strings.Split(intervalAsString, "-")
	var asInts []int
	for _, val := range asStrings {
		intValue, err := strconv.Atoi(val)
		if err != nil {
			fmt.Println("Invalid conversion to integer")
			os.Exit(1)
		}
		asInts = append(asInts, intValue)
	}
	return asInts
}

// determines if one interval is a fully contained within another
func isSubInterval(first []int, second []int) bool {
	if first[0] <= second[0] && first[1] >= second[1] {
		return true
	}
	if first[0] >= second[0] && first[1] <= second[1] {
		return true
	}
	return false
}

// determines if intervals overlap at all
func doesOverlap(first []int, second []int) bool {
	if second[1] < first[0] {
		return false
	}
	if first[1] < second[0] {
		return false
	}
	return true
}
