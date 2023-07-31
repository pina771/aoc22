package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var handToValue = map[string]string{
	"A": "rock",
	"B": "paper",
	"C": "scissors",
	"X": "rock",
	"Y": "paper",
	"Z": "scissors",
}

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current wrk dir")
		return
	}
	data, err := os.ReadFile(filepath.Join(cwd, "inputs", "day2.txt"))
	if err != nil {
		fmt.Println("Error reading file")
		fmt.Printf("err: %v\n", err)
		return
	}

	lines := strings.Split(string(data), "\n")

	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	var sum int
	for _, val := range lines {
		opponentAndOwn := strings.Split(val, " ")
		if len(opponentAndOwn) < 2 {
			break
		}
		handPoints := outcomePoints(handToValue[opponentAndOwn[0]], handToValue[opponentAndOwn[1]])
		sum += handPoints
		ownHand, ok := handToValue[opponentAndOwn[1]]
		if !ok {
			fmt.Println("Invalid value for opponent hand")
			os.Exit(1)
		}
		sum += ownHandPoints(ownHand)
	}
	fmt.Println(sum)
}

func part2(lines []string) {
	var sum int
	for _, val := range lines {
		opponentAndOutcome := strings.Split(val, " ")
		if len(opponentAndOutcome) < 2 {
			break
		}
		shouldThrowValue, ok := determineHandForOutcome(handToValue[opponentAndOutcome[0]], opponentAndOutcome[1])
		if !ok {
			fmt.Println("Error while determining what hand should throw")
			os.Exit(1)
		}
		sum += outcomePoints(handToValue[opponentAndOutcome[0]], shouldThrowValue)
		sum += ownHandPoints(shouldThrowValue)
	}
	fmt.Println(sum)
}

func determineHandForOutcome(opponentHand string, outcome string) (string, bool) {
	// X = lose
	// Y = draw
	// Z = win
	if outcome == "Y" {
		return opponentHand, true
	}
	switch opponentHand {
	case "rock":
		if outcome == "X" {
			return "scissors", true
		} else {
			return "paper", true
		}
	case "paper":
		if outcome == "X" {
			return "rock", true
		} else {
			return "scissors", true
		}
	case "scissors":
		if outcome == "X" {
			return "paper", true
		} else {
			return "rock", true
		}
	}
	return "", false
}

func ownHandPoints(ownHand string) int {
	switch ownHand {
	case "rock":
		return 1
	case "paper":
		return 2
	case "scissors":
		return 3
	}
	return 0
}

// Returns number of points for a hand
func outcomePoints(opponentHand string, ownHand string) int {
	switch opponentHand {
	case "rock":
		if ownHand == "paper" {
			return 6
		} else if ownHand == "rock" {
			return 3
		} else {
			return 0
		}
	case "paper":
		if ownHand == "scissors" {
			return 6
		} else if ownHand == "paper" {
			return 3
		} else {
			return 0
		}
	case "scissors":
		if ownHand == "rock" {
			return 6
		} else if ownHand == "scissors" {
			return 3
		} else {
			return 0
		}
	default:
		fmt.Println("Something went wrong")
		os.Exit(1)
	}
	return 0
}
