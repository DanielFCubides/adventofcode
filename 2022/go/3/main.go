package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := readFile()
	result := RuckSackReorganization2(input)
	fmt.Print(result)
}

func RuckSackReorganization2(input []string) int {
	result := 0
	size := len(input) / 3
	for i := 0; i < size; i++ {
		badge := GetElfBadge(input, i)
		value := StringNumericPriority(badge)
		result += value

	}
	return result
}

func StringNumericPriority(char string) int {
	valueOfA := 65
	valueOfa := 97
	if int(char[0]) >= valueOfa {
		return int(char[0]) + 1 - valueOfa
	}
	return int(char[0]) + 27 - valueOfA
}

func GetElfBadge(input []string, i int) string {
	elf1 := i * 3
	elf2 := elf1 + 1
	elf3 := elf1 + 2
	intersect1 := stringIntersection(input[elf1], input[elf2])
	result := stringIntersection(intersect1, input[elf3])[0]
	return string(result)
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

func stringIntersection(slice1 string, slice2 string) string {
	result := ""
	slice1Map := make(map[string]int, 0)
	for _, element := range slice1 {
		if _, ok := slice1Map[string(element)]; !ok {
			slice1Map[string(element)] = 1
		}

	}
	for _, charRune := range slice2 {
		if _, ok := slice1Map[string(charRune)]; ok {
			result += string(charRune)
		}
	}

	return result
}
