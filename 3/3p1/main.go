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
			fmt.Println(line)
			fmt.Println(lhs, rhs)
			duplicate := findDoubleItem(lhs, rhs)
			sum += calcPrio(duplicate)
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
	if len(str)%2 != 0 {
		fmt.Println("NOT EVEN LMAo")
	}
	lhs := str[0 : len(str)/2]
	rhs := str[len(str)/2 : len(str)]
	if len(lhs) == len(rhs) {
		return lhs, rhs
	} else {
		panic("splitInHalf: result is not the same size")
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
