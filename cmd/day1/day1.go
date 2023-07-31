package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory!")
		return
	}
	data, err := os.ReadFile(filepath.Join(cwd, "inputs", "day1.txt"))
	if err != nil {
		fmt.Println("Error reading file")
		return
	}

	lines := strings.Split(string(data), "\n")

	sums := []int{}
	current := 0

	for _, val := range lines {
		if len(val) == 0 {
			sums = append(sums, current)
			current = 0
		} else {
			i, err := strconv.Atoi(val)
			if err != nil {
				fmt.Println("Error while converting string to integer!")
				return
			}
			current = current + i
		}
	}

	fmt.Println(max(sums...))

	sort.Ints(sums)
	current = 0
	for _, val := range sums[len(sums)-3:] {
		current += val
	}
	fmt.Println(current)
}

// Finds the maximum value of a variable number of integers
func max(nums ...int) int {
	max := 0
	for _, i := range nums {
		if i > max {
			max = i
		}
	}
	return max
}
