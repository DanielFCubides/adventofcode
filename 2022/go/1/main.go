package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	filename := "input2.txt"
	readFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	part2(fileScanner)

}

func part1(fileScanner *bufio.Scanner) {
	var maxCalories = 0
	var calories = 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "" {
			if calories >= maxCalories {
				maxCalories = calories
			}
			calories = 0
			continue
		}
		intvalue, _ := strconv.Atoi(line)
		calories += intvalue
	}
	fmt.Println(maxCalories)
}

func part2(fileScanner *bufio.Scanner) {
	var elfs = []int{0, 0, 0}
	var calories = 0
	var elf = 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line != "" {
			intvalue, _ := strconv.Atoi(line)
			calories += intvalue
			fmt.Printf("%d, ", intvalue)
			continue
		}
		fmt.Println("\n", calories, elf)
		if calories >= elfs[0] {
			elfs[0] = calories
		} else if calories >= elfs[1] {
			elfs[1] = calories
		} else if calories >= elfs[2] {
			elfs[2] = calories
		}
		calories = 0
		elf += 1
	}
	fmt.Printf("len=%d cap=%d %v\n", len(elfs), cap(elfs), elfs)
	fmt.Println(elfs[0] + elfs[1] + elfs[2])
}
