package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type position struct {
	x int
	y int
}

func tailsPositions(movements []string) int {
	headPositionsMap := map[position]int{position{0, 0}: 1}
	tailPositionsMap := map[position]int{position{0, 0}: 1}
	currentHeadPosition := position{0, 0}
	currentTailPosition := position{0, 0}
	for _, move := range movements {
		simulateMovement(&headPositionsMap, &tailPositionsMap, &currentHeadPosition, &currentTailPosition, move)
	}
	fmt.Printf("H: %d - T: %d\n\n", len(headPositionsMap), len(tailPositionsMap))
	//for key, _ := range tailPositionsMap {
	//	fmt.Printf("x: %d y: %d\n ", key.x, key.y)
	//}
	return len(tailPositionsMap)
}

func moveHead(direction string, positions map[position]int, currentPosition *position) {
	switch direction {
	case "R":
		currentPosition.x++
		break
	case "L":
		currentPosition.x--
		break
	case "U":
		currentPosition.y++
		break
	case "D":
		currentPosition.y--
		break
	default:
		break
	}
	AddPosition(positions, currentPosition)

}

func AddPosition(Positions map[position]int, CurrentPosition *position) {
	if _, ok := Positions[*CurrentPosition]; !ok {
		Positions[*CurrentPosition] = 1
	}
	Positions[*CurrentPosition]++
}

func simulateMovement(headPositions *map[position]int, tailPositionsMap *map[position]int, currentHeadPosition *position, currentTailPosition *position, move string) {
	direction, units := parseMovement(move)
	//fmt.Printf(" %s %d\n", direction, units)
	for i := 0; i < units; i++ {
		previousHeadPosition := *currentHeadPosition
		moveHead(direction, *headPositions, currentHeadPosition)
		shouldMove := TailShouldMove(*currentTailPosition, *currentHeadPosition)
		if shouldMove {
			moveTail(currentTailPosition, previousHeadPosition, tailPositionsMap)
		}

		//fmt.Printf("H x:%d, y:%d\n", currentHeadPosition.x, currentHeadPosition.y)
		//fmt.Printf("T x:%d, y:%d\n", currentTailPosition.x, currentTailPosition.y)
	}
	//fmt.Println()
}

func moveTail(tail *position, head position, positions *map[position]int) {
	tail.x = head.x
	tail.y = head.y
	AddPosition(*positions, tail)
}

func TailShouldMove(tail position, head position) bool {
	return math.Abs(float64(tail.x-head.x)) > 1 || math.Abs(float64(tail.y-head.y)) > 1
}

func parseMovement(move string) (string, int) {
	instructions := strings.Split(move, " ")
	directions := instructions[0]
	units, _ := strconv.Atoi(instructions[1])
	return directions, units
}

func readFile() []string {
	filename := "input.txt"
	input := make([]string, 0)
	readFile, err := os.Open(filename)
	if err != nil {
		panic("no file")
	}
	fileScanner := bufio.NewScanner(readFile)
	for fileScanner.Scan() {
		input = append(input, fileScanner.Text())
	}

	return input
}

func main() {
	puzzle := readFile()
	//fmt.Println(puzzle)
	result := tailsPositions(puzzle)
	fmt.Println(result)
}
