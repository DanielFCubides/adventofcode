package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func one(instruction string, units int, x *int, y *int, aim *int) {
	if instruction == "forward" {
		*x += units
	}
	if instruction == "down" {
		*y += units
	}
	if instruction == "up" {
		*y -= units
	}
}

func two(instruction string, units int, x *int, y *int, aim *int) {
	if instruction == "forward" {
		*y += units * *aim
		*x += units
	}
	if instruction == "down" {
		*aim += units
	}
	if instruction == "up" {
		*aim -= units
	}
}

func process_instruction(process func(string, int, *int, *int, *int), instruction []string, units []int) int {
	x, y, aim := 0, 0, 0

	for i := 0; i < len(instruction); i++ {
		process(instruction[i], units[i], &x, &y, &aim)
	}
	return x * y
}

func mapInput(input string) (string, int) {
	inputs := strings.Split(input, " ")
	instruction := inputs[0]
	unit, _ := strconv.Atoi(inputs[1])
	return instruction, unit
}

func main() {
	filename := "input.txt"
	readFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var instructions []string
	var units []int

	for fileScanner.Scan() {
		//intvalue, _ := strconv.Atoi(fileScanner.Text())
		instruction, unit := mapInput(fileScanner.Text())
		instructions = append(instructions, instruction)
		units = append(units, unit)

	}

	result := process_instruction(one, instructions, units)
	fmt.Printf("part one: %d\n\n", result)

	result = process_instruction(two, instructions, units)
	fmt.Printf("part two: %d", result)
}
