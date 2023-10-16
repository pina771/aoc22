package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Stack []string

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) pop() (string, bool) {
	if s.isEmpty() {
		return "", false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

func (s *Stack) popN(n int) ([]string, bool) {
	if s.isEmpty() || n > len(*s) {
		return make([]string, 0), false
	} else {
		elements := (*s)[len(*s)-n:]
		(*s) = (*s)[0 : len(*s)-n]
		return elements, true
	}
}

func (s *Stack) push(elem string) {
	*s = append(*s, elem)
}

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current wd")
		os.Exit(1)
	}

	data, err := os.ReadFile(filepath.Join(cwd, "inputs", "day5.txt"))
	if err != nil {
		os.Exit(2)
	}
	lines := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")

	var idxOfBlank int
	for idx, val := range lines {
		if len(val) == 0 {
			idxOfBlank = idx
			break
		}
	}

	numOfStacks := getNumberOfStacks(lines[idxOfBlank-1])
	stacks := make([]Stack, numOfStacks)
	for i := idxOfBlank - 2; i >= 0; i-- {
		lineValues := strings.Split(lines[i], "")
		for idx := 0; idx < len(lineValues); idx++ {
			if lineValues[idx] != " " {
				stackIndex := determineStackIndex(idx)
				stacks[stackIndex].push(lineValues[idx+1])
				idx += 2
			}
		}
	}
	fmt.Println("Stacks: ")
	for idx, val := range stacks {
		fmt.Println(idx+1, " : ", val)
	}

	for _, val := range lines[idxOfBlank+1:] {
		if len(val) == 0 {
			break
		}
		lineValues := strings.Fields(val)
		numOfMove, _ := strconv.Atoi(lineValues[1])
		from, _ := strconv.Atoi(lineValues[3])
		to, _ := strconv.Atoi(lineValues[5])

		// craneMove9000(&stacks[from-1], &stacks[to-1], numOfMove)
		craneMove9001(&stacks[from-1], &stacks[to-1], numOfMove)
	}

	fmt.Println("After: ")
	for idx, val := range stacks {
		fmt.Println(idx+1, " : ", val)
	}

	var solution []string
	for idx, val := range stacks {
		topvalue, ok := val.pop()
		if !ok {
			fmt.Println("crate empty!", idx+1)
		}
		solution = append(solution, topvalue)
	}
	fmt.Println(strings.Join(solution, ""))
}

func craneMove9000(from *Stack, to *Stack, numOfMoves int) {
	for i := 0; i < numOfMoves; i++ {
		popped, ok := from.pop()
		if !ok {
			fmt.Println("Error during pop")
			fmt.Println("Attempted to move ", numOfMoves, " from ", from, "to ", to)
			os.Exit(5)
		}
		to.push(popped)
	}
}

func craneMove9001(from *Stack, to *Stack, numOfMoves int) {
	elemsToMove, ok := from.popN(numOfMoves)
	if !ok {
		fmt.Println("Error during popN")
		fmt.Println("Attempted to move ", numOfMoves, " from ", from, "to ", to)
		os.Exit(5)
	}
	for _, val := range elemsToMove {
		to.push(val)
	}
}

func determineStackIndex(idx int) int {
	n := idx
	retval := 0
	for n > 2 {
		n = n - 4
		retval++
	}

	return retval
}

func getNumberOfStacks(lineWithStackNums string) int {
	stackNumbers := strings.Fields(lineWithStackNums)
	numOfStacks, err := strconv.Atoi(stackNumbers[len(stackNumbers)-1])
	if err != nil {
		fmt.Println("Invalid last value in stack numbers")
		os.Exit(3)
	}
	fmt.Println("Number of stacks : ", numOfStacks)
	return numOfStacks
}
