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
	spritePosition := []int{1, 2, 3}
	CurrentCRTRow := "# "
	for i := 1; i <= lastCycle; i++ {
		//fmt.Printf("Start cycle %d: \n", i)
		if (i)%40 == 0 {
			result += i * RegisterX
			fmt.Println(CurrentCRTRow)
			CurrentCRTRow = ""
		}
		cycle := i % 40
		if executionInProgress > 0 {
			executionInProgress--
			if valueToAdd != 0 && executionInProgress == 0 {
				RegisterX += valueToAdd
				spritePosition = []int{RegisterX - 1, RegisterX, RegisterX + 1}
				//fmt.Printf("Sprite position [%d,%d,%d]\n", spritePosition[0], spritePosition[1], spritePosition[2])
				valueToAdd = 0
			}
			printCurrentCRTRwo(&CurrentCRTRow, spritePosition, cycle)
			continue
		}
		instruction := strings.Split(instructions[0], " ")
		instructions = instructions[1:]
		if len(instruction) == 1 {
			//fmt.Printf("end of cycle %d\n", i)
			printCurrentCRTRwo(&CurrentCRTRow, spritePosition, cycle)
			continue
		}
		if len(instruction) == 2 {
			executionInProgress++
			valueToAdd, _ = strconv.Atoi(instruction[1])
			//fmt.Printf("end of cycle %d\n", i)
		}
		printCurrentCRTRwo(&CurrentCRTRow, spritePosition, cycle)
	}
}
func printCurrentCRTRwo(CurrentCRTRow *string, spritePosition []int, col int) {
	if contains(spritePosition, col) {
		*CurrentCRTRow += "# "
	} else {
		*CurrentCRTRow += ". "
	}
	//fmt.Println(*CurrentCRTRow)
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
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
