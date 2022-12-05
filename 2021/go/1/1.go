package main
import (
		"fmt"
		"os"
		"bufio"
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

func main(){
	filename := "input.txt"
	readFile, err := os.Open(filename)

    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)
    var fileLines []int

    for fileScanner.Scan() {
		intvalue, _ := strconv.Atoi(fileScanner.Text())
        fileLines = append(fileLines, intvalue)
    }

    readFile.Close()

	result, _ := PartOne(fileLines)

    fmt.Println(result)
}
