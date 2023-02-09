package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

var ChoiceValue = map[string]int{"X": 1, "Y": 2, "Z": 3}
var ResultValue = map[string]int{"Y": 3, "X": 0, "Z": 6}
var ProcessValue = map[string]int{"X": 1, "Y": 2, "Z": 3, "A": 1, "B": 2, "C": 3}
var decisionMap = map[string]string{
	"AY": "X",
	"AX": "Z",
	"AZ": "Y",
	"BY": "Y",
	"BX": "X",
	"BZ": "Z",
	"CY": "Z",
	"CX": "Y",
	"CZ": "X",
}

const WinValue = 6
const LoseValue = 0
const DrawValue = 3

func main() {
	input := readFile()
	result := RPStournament(input, scoreRound)
	fmt.Printf("part 1: %d\n", result)
	result = RPStournament(input, scoreRoundPart2)
	fmt.Printf("part 2: %d\n", result)
}

func scoreRoundPart2(opponentChoice string, result string) int {
	yourChoice := decisionMap[fmt.Sprintf("%s", opponentChoice+result)]
	fmt.Printf("%s - %s : %s == (%d + %d)total result %d  ", opponentChoice, yourChoice, result, ChoiceValue[yourChoice], ResultValue[result], ChoiceValue[yourChoice]+ResultValue[result])
	return ChoiceValue[yourChoice] + ResultValue[result]
}

func RPStournament(input []string, scoreFunction func(string, string) int) int {
	result := 0
	for _, round := range input {
		choices := strings.Split(round, " ")
		result += scoreFunction(choices[0], choices[1])
		fmt.Println(result)
	}
	return result
}

func readFile() []string {
	puzzle := make([]string, 0)
	readFile, err := os.Open("input.txt")
	//readFile, err := os.Open("example.txt")
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
