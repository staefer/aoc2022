package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	readFile, err := os.Open("input.txt")
	checkError(err)

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	mostLoadedElf := scanAndSumNextElf(fileScanner)
	currentElf := 0

	for {
		currentElf = scanAndSumNextElf(fileScanner)

		if currentElf > mostLoadedElf {
			mostLoadedElf = currentElf
		}

		if currentElf == -1 {
			break
		}
	}

	fmt.Println("Most loaded elf KcalCount: ", mostLoadedElf)
}

func scanAndSumNextElf(fileScanner *bufio.Scanner) (elfTotalKcal int) {
	kcalSum := 0
	for fileScanner.Scan() {
		temp := fileScanner.Text()

		if temp == "" { // EoE -> End of Elf
			return kcalSum
		} else {
			asInt, err := strconv.Atoi(temp)
			checkError(err)
			kcalSum += asInt
		}
	}
	return -1
}
