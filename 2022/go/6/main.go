package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filename := "input.txt"
	puzzle := readFile(filename)
	startOfPacketMarket := detectMarker(puzzle, 4)
	startOfMessageMarket := detectMarker(puzzle, 14)
	fmt.Println(startOfPacketMarket)
	fmt.Println(startOfMessageMarket)

}

func readFile(filename string) string {
	puzzle := ""
	readFile, err := os.Open(filename)
	if err != nil {
		panic("no file")
	}
	fileScanner := bufio.NewScanner(readFile)
	for fileScanner.Scan() {
		puzzle = fileScanner.Text()
	}
	return puzzle
}

func detectMarker(datastream string, sizeOfMarker int) int {
	size := len(datastream)
	sizeOfTheMarker := sizeOfMarker
	for i := 0; i < size-sizeOfTheMarker; i++ {

		duplicates := detectDuplicates(datastream[i : i+sizeOfTheMarker])
		//fmt.Printf("%s -> %t : firstMarket at : %d\n", datastream[i:i+sizeOfTheMarker], duplicates, i+sizeOfTheMarker)
		if !duplicates {
			return i + sizeOfTheMarker
		}
	}
	return size - 1
}

func detectDuplicates(slice string) bool {
	frequencyMap := make(map[string]int, 4)
	for _, element := range slice {
		if _, ok := frequencyMap[string(element)]; ok {
			fmt.Printf("%s has a repeated element %s\n", slice, string(element))
			return true
		} else {
			frequencyMap[string(element)] = 1
		}
	}
	return false
}
