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

		if scanResult[1] == "X" { // I SHOULD LOSE
			result += shouldBeLose(scanResult[0])
		} else if scanResult[1] == "Y" { // I SHOULD DRAW
			result += shouldBeDraw(scanResult[0])
		} else if scanResult[1] == "Z" { // I SHOULD WIN
			result += shouldBeWin(scanResult[0])
		}

	}

	fmt.Println("Result = ", result)
}

func shouldBeDraw(oponentMove string) (resultPoints int) {
	if oponentMove == "A" {
		return 3 + 1
	} else if oponentMove == "B" {
		return 3 + 2
	} else if oponentMove == "C" {
		return 3 + 3
	} else {
		panic("Draw: Oponent input wrong")
	}
}

func shouldBeWin(oponentMove string) (resultPoints int) {
	if oponentMove == "A" {
		return (6 + 2)
	} else if oponentMove == "B" {
		return (6 + 3)
	} else if oponentMove == "C" {
		return (6 + 1)
	} else {
		panic("Win: Oponent input wrong")
	}
}

func shouldBeLose(oponentMove string) (resultPoints int) {
	if oponentMove == "A" {
		return 3
	} else if oponentMove == "B" {
		return 1
	} else if oponentMove == "C" {
		return 2
	} else {
		panic("Lose: Oponent input wrong")
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(-1)
	}
}
