package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	stack, instructions := readFile()
	result := SupplyStack(stack, instructions)
	fmt.Printf("%s-%s\n%s", stack, instructions, result)
}

func SupplyStack(stack []string, instructions []string) (result string) {
	for _, instruction := range instructions {
		stack = executeInstruction(stack, instruction, false)
	}
	for _, pile := range stack {
		result += string(pile[0])
	}
	return
}

func executeInstruction(stack []string, instruction string, reversed bool) []string {
	//fmt.Println(instruction)
	instructionParsed := strings.Split(instruction, ",")
	items, _ := strconv.Atoi(instructionParsed[0])
	initialStack, _ := strconv.Atoi(instructionParsed[1])
	FinalStack, _ := strconv.Atoi(instructionParsed[2])
	initialStack -= 1
	FinalStack -= 1
	//fmt.Println(stack)
	//fmt.Printf("move %d from %d to %d\n", items, initialStack, FinalStack)
	elementsToMove := stack[initialStack][:items]
	stack[initialStack] = stack[initialStack][items:]
	if reversed {
		stack[FinalStack] = reverse(elementsToMove) + stack[FinalStack]
		return stack
	}
	stack[FinalStack] = elementsToMove + stack[FinalStack]
	return stack

}

func readFile() ([]string, []string) {

	//stackSize := 3    // 3
	//HighestStack := 4 // 3
	//lineLenght := 11  // 11
	//readFile, err := os.Open("example.txt")

	stackSize := 9    // 3
	HighestStack := 9 // 3
	lineLenght := 36  // 11
	readFile, err := os.Open("input.txt")

	stack := make([]string, stackSize)
	instructions := make([]string, 0)
	if err != nil {
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	line := 1
	for fileScanner.Scan() {
		if line < HighestStack {
			parseStacks(fileScanner, lineLenght, stack)
			line++
			continue
		}
		if line > HighestStack+1 {
			instructions = parseInstruction(fileScanner, instructions)
		}
		line++
	}
	return stack, instructions
}

func parseInstruction(fileScanner *bufio.Scanner, instructions []string) []string {
	text := fileScanner.Text()
	text = strings.ReplaceAll(text, "move ", "")
	text = strings.ReplaceAll(text, " from ", ",")
	text = strings.ReplaceAll(text, " to ", ",")
	instructions = append(instructions, text)
	return instructions
}

func parseStacks(fileScanner *bufio.Scanner, lineLenght int, stack []string) {
	text := fileScanner.Text()
	textSize := len(text)
	if textSize < lineLenght {
		text = fmt.Sprintf("%s%*s", text, lineLenght-textSize, "")
	}
	//fmt.Printf("%d-line %s -> \n", len(text), text)
	for i, CharRune := range text {
		stringValue := string(CharRune)
		if stringValue != " " && stringValue != "[" && stringValue != "]" {
			//1 5 9
			//fmt.Printf("store %s in the stack %d,because have an index of %d\n", stringValue, i/4, i)
			stack[i/4] += stringValue
		}
	}
	fmt.Println()
}

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}
