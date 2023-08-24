package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const SHORTER = 0
const LARGEST = 9

func readFile(filename string) []string {
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

func visibleFromTheOutside(treeMap []string) [][]int {
	size := len(treeMap[0])
	visibility := make([][]int, size)
	for i := range visibility {
		visibility[i] = make([]int, size)
	}
	largestFromTop := make([]int, size)
	largestFromLeft := make([]int, size)
	largestFromBottom := make([]int, size)
	largestFromRight := make([]int, size)
	for row := 0; row < size; row++ {
		for col := 0; col < size; col++ {
			VisibilityCheck(treeMap, row, col, size, largestFromTop, visibility, largestFromLeft, largestFromBottom, largestFromRight)

		}
	}

	//largestFromRight := make([]int, size)

	return visibility
}

func VisibilityCheck(treeMap []string, row int, col int, size int, largestFromTop []int, visibility [][]int, largestFromLeft []int, largestFromBottom []int, largestFromRight []int) {
	tree, _ := strconv.Atoi(string(treeMap[row][col]))
	oppositeTree, _ := strconv.Atoi(string(treeMap[size-row-1][size-col-1]))
	//fmt.Printf("tree: %d, opposite: %d", tree, oppositeTree)
	if tree > largestFromTop[col] {
		largestFromTop[col] = tree
		visibility[row][col] = 1
		//fmt.Printf(" tree visible from top ")
	}
	if tree > largestFromLeft[row] {
		largestFromLeft[row] = tree
		visibility[row][col] = 1
		//fmt.Printf(" tree visible from left ")
	}
	if oppositeTree > largestFromBottom[size-1-col] {
		largestFromBottom[size-1-col] = oppositeTree
		visibility[size-row-1][size-col-1] = 1
		//fmt.Printf(" opposite visible from bottom ")
	}
	if oppositeTree > largestFromRight[size-1-row] {
		largestFromRight[size-1-row] = oppositeTree
		visibility[size-row-1][size-col-1] = 1
		//fmt.Printf(" opposite visible from right ")
	}
	if col == 0 || row == 0 || col == size-1 || row == size-1 {
		visibility[row][col] = 1
		//fmt.Printf(" tree is a corner ")
	}
	//fmt.Println(" . ")
}

func highestScenicScore(treeMap []string) int {
	size := len(treeMap)
	//fmt.Println(size)
	highScenicScore := 0
	for row := 0; row < size; row++ {
		for col := 0; col < size; col++ {
			isTreeOnEdge := row == 0 || col == 0 || row == size-1 || col == size-1
			if !isTreeOnEdge {
				scenicScore := calculateScenicScore(treeMap, row, col)
				if scenicScore > highScenicScore {
					highScenicScore = scenicScore
				}
			}
		}
	}
	return highScenicScore
}

func calculateScenicScore(treeMap []string, row int, col int) int {
	size := len(treeMap[0])
	currentTree, _ := strconv.Atoi(string(treeMap[row][col]))
	//check up
	upScore := 0
	for i := row - 1; i >= 0; i-- {
		tree, _ := strconv.Atoi(string(treeMap[i][col]))
		if tree >= currentTree {
			upScore++
			break
		}
		upScore++
	}
	//check down
	downScore := 0
	for i := row + 1; i < size; i++ {
		tree, _ := strconv.Atoi(string(treeMap[i][col]))
		if tree >= currentTree {
			downScore++
			break
		}
		downScore++
	}
	//check left
	leftScore := 0
	for i := col - 1; i >= 0; i-- {
		tree, _ := strconv.Atoi(string(treeMap[row][i]))
		if tree >= currentTree {
			leftScore++
			break
		}
		leftScore++
	}
	//check right
	rightScore := 0
	for i := col + 1; i < size; i++ {
		tree, _ := strconv.Atoi(string(treeMap[row][i]))
		if tree >= currentTree {
			rightScore++
			break
		}
		rightScore++
	}
	result := upScore * downScore * leftScore * rightScore
	//fmt.Printf("currentTree: %d, up: %d, down: %d, left: %d, right: %d, result :%d\n", currentTree, upScore, downScore, leftScore, rightScore, result)
	return result
}

func main() {
	treeMap := readFile("input.txt")
	//for _, s := range treeMap {
	//	fmt.Println(s)
	//}
	//visibilityFromTheOutside(treeMap)
	fmt.Println(highestScenicScore(treeMap))
}

func visibilityFromTheOutside(treeMap []string) {
	fromTheOutside := visibleFromTheOutside(treeMap)
	result := 0
	for _, values := range fromTheOutside {
		for _, value := range values {
			fmt.Printf("%d ", value)
			result += value
		}
		fmt.Printf("\n")
	}
	fmt.Print(result)
}
