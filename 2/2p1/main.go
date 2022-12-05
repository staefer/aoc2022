package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	readFile, err := os.Open("input.txt")
	checkError(err)

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var scanResult []string
	result := 0

	for fileScanner.Scan() {
		scanResult = strings.Split(fileScanner.Text(), " ")

		if scanResult[1] == "X" { // ROCK
			result += 1
			if scanResult[0] == "C" { // SCISSOR, win
				result += 6
			} else if scanResult[0] == "A" { // ROCK, even
				result += 3
			}
		} else if scanResult[1] == "Y" { // PAPER
			result += 2
			if scanResult[0] == "A" { // ROCK, win
				result += 6
			} else if scanResult[0] == "B" { // PAPER, even
				result += 3
			}
		} else if scanResult[1] == "Z" { // SCISSOR
			result += 3
			if scanResult[0] == "B" { // PAPER, win
				result += 6
			} else if scanResult[0] == "C" { // SCISSOR, even
				result += 3
			}
		}
	}

	fmt.Println("Result = ", result)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(-1)
	}
}
