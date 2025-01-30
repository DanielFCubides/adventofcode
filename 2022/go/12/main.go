package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Position struct {
	x, y int
}

type PositionVisited struct {
	position          Position
	stepsFromTheStart int
	altitude          string
}

func main() {
	fmt.Println("Welcome to the day 12 first part 2022")
	puzzle := readInput("input.txt")
	answer := HillClimbingAlgorithm(puzzle)
	fmt.Println("Steps required:", answer)
}

func HillClimbingAlgorithm(puzzle [][]string) int {
	currentPosition := GetPosition(puzzle, "S")
	finalPosition := GetPosition(puzzle, "E")
	fmt.Printf("Start Position: %v\nEnd Position: %v\n", currentPosition, finalPosition)
	result := HillClimbing(puzzle, currentPosition)
	return result
}

func HillClimbing(puzzle [][]string, currentPosition Position) int {
	queue := make([]PositionVisited, 0)
	visited := make(map[string]bool)

	// Initialize with start position
	queue = append(queue, PositionVisited{
		position:          currentPosition,
		stepsFromTheStart: 0,
		altitude:          "a", // S has elevation a
	})
	visited[createKey(currentPosition)] = true

	directions := []Position{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	maxY := len(puzzle)
	maxX := len(puzzle[0])

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		// Check if reached end
		if puzzle[current.position.y][current.position.x] == "E" {
			return current.stepsFromTheStart
		}

		// Try each direction
		for _, direction := range directions {
			nextPosition := Position{
				x: current.position.x + direction.x,
				y: current.position.y + direction.y,
			}

			// Check bounds
			if nextPosition.x < 0 || nextPosition.x >= maxX ||
				nextPosition.y < 0 || nextPosition.y >= maxY {
				continue
			}

			// Skip if visited
			key := createKey(nextPosition)
			if visited[key] {
				continue
			}

			// Get current and next elevation
			currentElevation := getElevation(puzzle[current.position.y][current.position.x])

			nextElevation := getElevation(puzzle[nextPosition.y][nextPosition.x])

			// Check if can move
			if nextElevation <= currentElevation+1 {
				stepsFromTheStart := current.stepsFromTheStart
				if currentElevation != int('a') {
					stepsFromTheStart = current.stepsFromTheStart + 1
				}
				queue = append(queue, PositionVisited{
					position:          nextPosition,
					stepsFromTheStart: stepsFromTheStart,
					altitude:          puzzle[nextPosition.y][nextPosition.x],
				})
				visited[key] = true
			}
		}
	}

	return -1 // No path found
}

func getElevation(char string) int {
	switch char {
	case "S":
		return int('a')
	case "E":
		return int('z')
	default:
		return int(char[0])
	}
}

func createKey(current Position) string {
	return fmt.Sprintf("%d,%d", current.x, current.y)
}

func GetPosition(puzzle [][]string, character string) Position {
	for y, row := range puzzle {
		for x, cell := range row {
			if cell == character {
				return Position{x, y}
			}
		}
	}
	return Position{0, 0}
}

func readInput(fileName string) [][]string {
	puzzle := make([][]string, 0)
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return puzzle
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		columns := strings.Split(line, "")
		puzzle = append(puzzle, columns)
	}

	return puzzle
}
