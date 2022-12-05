package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	readFile, err := os.Open("input.txt")
	errCheck(err)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var sum int
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line != "" {
			lhs, rhs := splitInHalf(line)
			duplicate := findDoubleItem(lhs, rhs)
			sum += calcPrio(duplicate)
			fmt.Println("Duplicate: ", duplicate, " Priority: ", calcPrio(duplicate))
		}
	}

	fmt.Println(sum)
}

func errCheck(err error) {
	if err != nil {
		panic("Error found")
	}
}

func splitInHalf(str string) (string, string) {
	lhs := str[0 : (len(str)/2)-1]
	rhs := str[(len(str) / 2) : len(str)-1]
	if len(lhs) == len(rhs) {
		return lhs, rhs
	} else {
		panic("splitInHalf no the same size")
	}
}

func calcPrio(item string) int {
	if item == "" {
		return 0
	}

	alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alphabetSlize := strings.Split(alphabet, "")

	for i := 0; i < len(alphabetSlize); i++ {
		if item == alphabetSlize[i] {
			return i + 1
		}
	}

	panic("Couldn't calculate priority")
}

func findDoubleItem(lhs string, rhs string) string {
	lhsSplice := strings.Split(lhs, "")
	rhsSplice := strings.Split(rhs, "")
	for _, lhsItem := range lhsSplice {
		for _, rhsItem := range rhsSplice {
			if lhsItem == rhsItem {
				return lhsItem
			}
		}
	}

	return ""
}
