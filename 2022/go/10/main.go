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
	//printStringList(input)
	startClockCircuit(input)
}

func startClockCircuit(instructions []string) {
	RegisterX := 1
	lastCycle := 240
	executionInProgress := 0
	valueToAdd := 0
	result := 0
	for i := 1; i <= lastCycle; i++ {
		if (i-20)%40 == 0 {
			result += i * RegisterX
			fmt.Printf("cycle %d - X value is %d - %d - result value is %d\n", i, RegisterX, i*RegisterX, result)
		}
		if executionInProgress > 0 {
			executionInProgress--
			if valueToAdd != 0 && executionInProgress == 0 {
				RegisterX += valueToAdd
				//fmt.Printf("setting x to %d\n", RegisterX)
				valueToAdd = 0
			}
			continue
		}
		instruction := strings.Split(instructions[0], " ")
		instructions = instructions[1:]
		if len(instruction) == 1 {
			//fmt.Printf("end of cycle %d\n", i)
			continue
		}
		if len(instruction) == 2 {
			executionInProgress++
			valueToAdd, _ = strconv.Atoi(instruction[1])
			//fmt.Printf("end of cycle %d\n", i)
		}

	}
}

func printStringList(input []string) {
	for i, item := range input {
		fmt.Println(i, item)
	}
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
