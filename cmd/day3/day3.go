package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current wrk dir")
		return
	}
	fmt.Println("Current working dir: ", cwd)
	data, err := os.ReadFile(filepath.Join(cwd, "inputs", "day3.txt"))
	if err != nil {
		fmt.Println("Error reading file")
		fmt.Printf("err: %v\n", err)
		return
	}
	partOne(data)
	partTwo(data)
}

func partOne(data []byte) {
	lines := strings.Split(string(data), "\n")
	var sum int
	for _, val := range lines {
		if len(val) == 0 {
			break
		}
		compartments := splitHalves(val)
		alreadyChecked := map[rune]int{}

		for _, item := range compartments[0] {
			if _, ok := alreadyChecked[item]; strings.ContainsRune(compartments[1], rune(item)) && !ok {
				prio := determinePriority(item)
				if prio == -1 {
					fmt.Println("Wrong priority determined")
					os.Exit(1)
				}
				alreadyChecked[item] = 1
				sum += prio
			}
		}
	}

	fmt.Println(sum)
}

func partTwo(data []byte) {
	lines := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")
	sum := 0
	for i := 0; i < len(lines)-1; i += 3 {
		group := lines[i : i+3]
		groupItems := map[rune]int{}
		for idx, bag := range group {
			for _, item := range bag {
				if val, ok := groupItems[item]; val == idx && (ok || idx == 0) {
					groupItems[item] = idx + 1
				}
			}
		}
		badge, _ := determineBadge(groupItems)
		sum += determinePriority(badge)
	}
	fmt.Println(sum)
}

func determineBadge(groupItems map[rune]int) (rune, bool) {
	for key, value := range groupItems {
		if value == 3 {
			return key, true
		}
	}
	return 0, false
}

func splitHalves(rucksack string) []string {
	half := len(rucksack) / 2
	return []string{rucksack[:half], rucksack[half:]}
}

func determinePriority(item rune) int {
	// A = 65
	// Z = 96
	// a = 97
	// z = 122
	if item >= 'a' && item <= 'z' {
		return int(item) - 96
	}
	if item >= 'A' && item <= 'Z' {
		return int(item) - 38
	}
	return -1
}
