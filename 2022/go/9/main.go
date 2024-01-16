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

func (p position) calculateDistance(position2 position) int {
	return int(math.Sqrt(math.Pow(float64(position2.x-p.x), 2) + math.Pow(float64(position2.y-p.y), 2)))
}

type rope struct {
	knotPosition  map[int]*position
	tailPositions map[string]int
}

func (r rope) move(direction string, units int) {
	fmt.Printf("\n%s - %d | | | tail movementes: %+v \n\n ", direction, units, r.tailPositions)
	for i := 0; i < units; i++ {
		r.moveKnots(direction)
	}
}

func (r rope) printPositions() {
	for i := 0; i < len(r.knotPosition); i++ {
		knot := r.knotPosition[i]
		fmt.Printf("knot %d [ %d, %d]\n", i, knot.x, knot.y)
	}
	fmt.Println()
}

func (r rope) moveKnots(direction string) {
	switch direction {
	case "R":
		r.knotPosition[0].x++
		break
	case "L":
		r.knotPosition[0].x--
		break
	case "U":
		r.knotPosition[0].y++
		break
	case "D":
		r.knotPosition[0].y--
		break
	default:
		break
	}
	for i := 0; i < len(r.knotPosition); i++ {
		knot := r.knotPosition[i]
		if i == 0 {
			continue
		}
		knotPosition := *knot
		fmt.Printf("%d - %+v:  ", i, knotPosition)
		PlankMove(r.knotPosition[i-1], knot)
		if i == len(r.knotPosition)-1 {
			tailPosition := fmt.Sprintf("%d,%d", knot.x, knot.y)
			if _, ok := r.tailPositions[tailPosition]; !ok {
				r.tailPositions[tailPosition] = 1
			}
		}
	}
	r.printPositions()
}

func PlankMove(head *position, tail *position) {
	distance := head.calculateDistance(*tail)
	//fmt.Printf("(%d, %d) head , (%d, %d) Tail, (%d, %d) hpp, (%d) distance\n",
	//	head.x,
	//	head.y,
	//	tail.x,
	//	tail.y,
	//	previousHeadPosition.x,
	//	previousHeadPosition.y,
	//	distance)
	if distance <= 1 {
		fmt.Println()
		return
	}
	fmt.Printf("  move knot from (%d,%d) ", tail.x, tail.y)
	xMovement := 0
	if tail.x > head.x {
		xMovement = -1
	}
	if tail.x < head.x {
		xMovement = 1
	}
	yMovement := 0
	if tail.y > head.y {
		yMovement = -1
	}
	if tail.y < head.y {
		yMovement = 1
	}
	*tail = position{x: tail.x + xMovement, y: tail.y + yMovement}
	fmt.Printf(" to (%d,%d)\n\n", tail.x, tail.y)
}

func tailsPositions(movements []string) int {
	rope := rope{
		knotPosition: map[int]*position{
			0: {0, 0},
			1: {0, 0},
			2: {0, 0},
			3: {0, 0},
			4: {0, 0},
			5: {0, 0},
			6: {0, 0},
			7: {0, 0},
			8: {0, 0},
			9: {0, 0},
		},
		tailPositions: make(map[string]int, 0),
	}
	for _, move := range movements {
		rope.move(parseMovement(move))
	}
	fmt.Print(rope.tailPositions)
	return len(rope.tailPositions)
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
	result := tailsPositions(puzzle)
	fmt.Println(result)
}
