package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("could not get cwd")
		os.Exit(2)
	}

	data, err := os.ReadFile(filepath.Join(cwd, "inputs", "day6.txt"))
	if err != nil {
		fmt.Println("could not read file")
		os.Exit(2)
	}

	stringData := string(data)
	var bfr []rune
	// number of distinct elements that determine sequence
	distinctElemsForSeq := 14
	for i := 0; i < distinctElemsForSeq; i++ {
		bfr = append(bfr, rune(stringData[i]))
	}

	for i := distinctElemsForSeq; i < len(stringData); i++ {
		bfr = bfr[1:]
		bfr = append(bfr, rune(stringData[i]))
		printBfr(bfr)
		if isDistinctSequence(bfr, distinctElemsForSeq) {
			fmt.Println(i + 1)
			break
		}
	}
}

func printBfr(buffer []rune) {
	fmt.Printf("[")
	for _, r := range buffer {
		fmt.Printf(" %c", r)
	}
	fmt.Printf(" ]\n")
}

func isDistinctSequence(buffer []rune, distinctElems int) bool {
	if len(buffer) != distinctElems {
		return false
	}
	a := make(map[rune]bool)
	for _, val := range buffer {
		a[val] = true
	}
	return len(a) == distinctElems
}
