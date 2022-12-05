package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func PartOne(nums []int) (int, error) {
	last := nums[0]
	count := 0

	for _, val := range nums {
		if val > last {
			count += 1
		}
		last = val
	}

	return count, nil
}

func main() {
	filename := "input.txt"
	readFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var elf = 0
	var maxCalories = 0
	var calories = 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "" {
			if calories >= maxCalories {
				maxCalories = calories
			}
			elf += 1
			continue
		}
		intvalue, _ := strconv.Atoi(line)
		calories += intvalue
	}

	fmt.Println(maxCalories)
}
