package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := readFile()
	result := CleanUpSectionsOverLap(input)
	fmt.Print(result)
}

func CleanUpSectionsOverLap(input []string) int {
	result := 0
	for _, duple := range input {
		range1floor, range1roof, range2floor, range2roof := parseDuple(duple)
		if rangeOverlap(range1floor, range1roof, range2floor, range2roof) {
			//fmt.Printf("[%d : %d] - [%d:%d] \n", range1floor, range1roof, range2floor, range2roof)
			result += 1
		}
	}
	return result
}

func parseDuple(duple string) (int, int, int, int) {
	pairs := strings.Split(duple, ",")
	pair1 := strings.Split(pairs[0], "-")
	pair2 := strings.Split(pairs[1], "-")
	range1floor, _ := strconv.Atoi(pair1[0])
	range1roof, _ := strconv.Atoi(pair1[1])
	range2floor, _ := strconv.Atoi(pair2[0])
	range2roof, _ := strconv.Atoi(pair2[1])
	return range1floor, range1roof, range2floor, range2roof
}

func rangeOverlap(range1floor int, range1roof int, range2floor int, range2roof int) bool {
	if range1floor >= range2floor && range1floor <= range2roof {
		//fmt.Printf("%d >= %d and %d <= %d\n", range1floor, range2floor, range1floor, range2roof)
		return true
	}
	if range1roof >= range2floor && range1roof <= range2roof {
		//fmt.Printf("%d >= %d and %d <= %d\n", range1roof, range2floor, range1roof, range2roof)
		return true
	}
	return completeRangeOverLap(range1floor, range1roof, range2floor, range2roof)
}
func completeRangeOverLap(range1floor int, range1roof int, range2floor int, range2roof int) bool {
	if range1floor <= range2floor && range1roof >= range2roof {
		//fmt.Printf("%d <= %d and %d >= %d\n", range1floor, range2floor, range1roof, range2roof)
		return true
	}
	if range2floor <= range1floor && range2roof >= range1roof {
		//fmt.Printf("%d <= %d and %d >= %d\n", range2floor, range1floor, range2roof, range1roof)
		return true
	}
	return false
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
