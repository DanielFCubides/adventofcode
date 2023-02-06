package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

var ChoiceValue = map[string]int{"X": 1, "Y": 2, "Z": 3}
var ProcessValue = map[string]int{"X": 1, "Y": 2, "Z": 3, "A": 1, "B": 2, "C": 3}

const WinValue = 6
const LoseValue = 0
const DrawValue = 3

func main() {
	input := readFile()
	result := RPStournament(input)
	fmt.Println(result)
}

func RPStournament(input []string) int {
	result := 0
	for _, round := range input {
		choices := strings.Split(round, " ")
		result += scoreRound(choices[0], choices[1])
	}
	return result
}

func readFile() []string {
	puzzle := make([]string, 0)
	readFile, err := os.Open("input.txt")
	if err != nil {
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		puzzle = append(puzzle, fileScanner.Text())
	}
	return puzzle
}

func scoreRound(opponentChoice string, yourChoice string) int {
	matchValue := matchCalculator(opponentChoice, yourChoice)
	choiceValue := ChoiceValue[yourChoice]
	result := choiceValue + matchValue
	return result
}

func matchCalculator(choiceOpponent string, yourChoice string) int {
	opponentValue := ProcessValue[choiceOpponent]
	yourValue := ProcessValue[yourChoice]
	difference := int(math.Abs(float64(opponentValue - yourValue)))
	switch difference {
	case 0:
		return DrawValue
	case 1:
		if yourValue > opponentValue {
			return WinValue
		}
		return LoseValue
	default:
		if yourValue < opponentValue {
			return WinValue
		}
		return LoseValue
	}

}
