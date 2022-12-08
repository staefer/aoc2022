package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileScanner := openScanner("input.txt")

	var sum int
	for {
		input := scanNextInput(fileScanner)

		if input == "" { // if EOF
			break
		}

		range1, range2 := splitToRanges(input)
		range1_1, range1_2 := splitRange(range1)
		range2_1, range2_2 := splitRange(range2)

		if range1_1 <= range2_1 && range1_2 >= range2_2 {
			sum++
		} else if range2_1 <= range1_1 && range2_2 >= range1_2 {
			sum++
		}
	}
	fmt.Println(sum)
}

func openScanner(fileName string) *bufio.Scanner {
	filePath, err := os.Open(fileName)
	if err != nil {
		panic("Can't open file")
	}
	fileScanner := bufio.NewScanner(filePath)
	fileScanner.Split(bufio.ScanLines)
	return fileScanner
}

func scanNextInput(fileScanner *bufio.Scanner) string {
	scanStatus := fileScanner.Scan()
	if scanStatus {
		return fileScanner.Text()
	} else {
		return "" // EOF case
	}
}

func splitToRanges(line string) (string, string) {
	lines := strings.Split(line, ",")
	return lines[0], lines[1]
}

func splitRange(line string) (int, int) {
	result := strings.Split(line, "-")
	from, err := strconv.Atoi(result[0])
	if err != nil {
		panic("splitRange error")
	}
	to, err := strconv.Atoi(result[1])
	if err != nil {
		panic("splitRange error")
	}
	return from, to
}
