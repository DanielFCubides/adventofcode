package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := readFile()
	result := RuckSackReorganization(input)
	fmt.Print(result)
}

func RuckSackReorganization(input []string) int {
	for _, s := range input {
		left, right := DivideCompartment(s)
		element := intersection(left, right)
		fmt.Printf("izq: %s | der: %s = element : %s\n", left, right, element)
	}
	return 1
}

func intersection(slice1 string, slice2 string) string {
	slice1Map := make(map[string]int, 0)
	for _, element := range slice1 {
		if _, ok := slice1Map[string(element)]; !ok {
			slice1Map[string(element)] = 1
		}

	}
	for _, charRune := range slice2 {
		if _, ok := slice1Map[string(charRune)]; ok {
			return string(charRune)
		}
	}

	return ""
}

func DivideCompartment(rucksack string) (string, string) {
	return rucksack[:len(rucksack)/2], rucksack[len(rucksack)/2:]
}

func readFile() []string {
	puzzle := make([]string, 0)
	//readFile, err := os.Open("input.txt")
	readFile, err := os.Open("example.txt")
	if err != nil {
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		puzzle = append(puzzle, fileScanner.Text())
	}
	return puzzle
}
