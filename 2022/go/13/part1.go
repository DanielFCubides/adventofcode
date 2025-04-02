package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func solvePart1(input []string) {
	sum := 0
	for i := 0; i < len(input); i += 3 {
		left, right := input[i], input[i+1]

		if isOrderRight(left, right) {
			packetIndex := i/3 + 1
			sum += packetIndex
			fmt.Printf("index #%d ,total %d\n", i+1, sum)
		}
	}

	fmt.Println("Sum of indices:", sum)
}

func main() {
	input := ReadInputFile("input.txt")

	timer := time.Now()
	//solvePart1(input)
	solvePart2(input)
	fmt.Println("Time:", time.Since(timer))
}

func solvePart2(input []string) {
	DividerPackages := []string{"[[2]]", "[[6]]"}
	packets := CreatePacketsArray(input, DividerPackages)
	packets = SortPackets(packets)
	fmt.Println(packets)
	decoderKey := getDecoderKey(packets, DividerPackages)
	fmt.Println(decoderKey)

}

func getDecoderKey(packets []string, DividerPackages []string) int {
	decoderKey := 1
	for _, dividerPackage := range DividerPackages {
		for i, packet := range packets {
			if packet == dividerPackage {
				index := i + 1
				fmt.Printf("Packet #%d is divider package: %s\n", index, packet)
				decoderKey = decoderKey * index
			}
		}
	}
	return decoderKey
}

func SortPackets(packets []string) []string {
	for i := 0; i < len(packets); i++ {
		for j := i + 1; j < len(packets); j++ {
			if isOrderRight(packets[i], packets[j]) {
				packets[i], packets[j] = packets[j], packets[i]
			}
		}
	}
	return packets
}

func CreatePacketsArray(input []string, DividerPackages []string) []string {
	packets := make([]string, 0)
	for i := 0; i < len(input); i++ {
		if input[i] == "" || len(input) == 0 {
			continue
		}
		packets = append(packets, input[i])
	}
	packets = append(packets, DividerPackages[0])
	packets = append(packets, DividerPackages[1])
	return packets
}

func ReadInputFile(path string) []string {
	fmt.Println("Reading from input file...", path)

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	arr := make([]string, 0)
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		arr = append(arr, sc.Text())
	}

	return arr
}
