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
	}

	//largestFromRight := make([]int, size)

	return visibility
}

func main() {
	treeMap := readFile("input.txt")
	for _, s := range treeMap {
		fmt.Println(s)
	}
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
