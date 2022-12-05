package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := readInput()

	gamma, epsilon := BinaryDiagnostic(input)
	result := gamma * epsilon
	fmt.Println("result part one: ", result)

	orating, corating := lifeSupportRating(input)
	fmt.Println("result part two: ", orating*corating)
}

func lifeSupportRating(input []string) (int, int) {
	common := []string{}
	leaser := []string{}
	commonList, leaserList := input, input
	for index := 0; index < len(input[0]); index++ {
		if len(commonList) > 1 {
			common = searchElement(commonList, index, true)
		}
		if len(leaserList) > 1 {
			leaser = searchElement(leaserList, index, false)
		}
		commonList = common
		leaserList = leaser
	}
	commonDecimal, _ := strconv.ParseInt(common[0], 2, 64)
	leaserDecimal, _ := strconv.ParseInt(leaser[0], 2, 64)
	return int(commonDecimal), int(leaserDecimal)
}

func searchElement(input []string, index int, common bool) []string {
	zeros := make([]string, 0)
	ones := make([]string, 0)
	for _, value := range input {
		if "0" == string(value[index]) {
			zeros = append(zeros, value)
		} else {
			ones = append(ones, value)
		}
	}
	if (len(ones) >= len(zeros)) == common {
		return ones
	}
	return zeros
}

func BinaryDiagnostic(input []string) (int, int) {
	resultLen := len(input)
	diagnostic := make([]int, len(input[0]))
	var epsilonstr, gammastr string
	mediumValue := resultLen / 2
	sumDiagnostic(input, diagnostic)
	gammastr, epsilonstr = validateDiagnostic(diagnostic, mediumValue, gammastr, epsilonstr)
	epsilon, _ := strconv.ParseInt(epsilonstr, 2, 64)
	gamma, _ := strconv.ParseInt(gammastr, 2, 64)
	return int(gamma), int(epsilon)
}

func validateDiagnostic(diagnostic []int, mediumValue int, gammastr string, epsilonstr string) (string, string) {
	for _, value := range diagnostic {
		if value > mediumValue {
			gammastr += "1"
			epsilonstr += "0"
		} else {
			gammastr += "0"
			epsilonstr += "1"
		}
	}
	return gammastr, epsilonstr
}

func sumDiagnostic(input []string, diagnostic []int) {
	for _, value := range input {
		for index, r := range value {
			binary, _ := strconv.Atoi(string(r))
			diagnostic[index] += binary
		}
	}
}

func readInput() []string {
	filename := "input.txt"
	readFile, err := os.Open(filename)
	var input []string
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		//intvalue, _ := strconv.Atoi(fileScanner.Text())
		input = append(input, fileScanner.Text())
	}
	return input
}
