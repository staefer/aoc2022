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

	var elf1 string
	var elf2 string
	var elf3 string

	var sum int = 0
	for fileScanner.Scan() {
		elf1 = fileScanner.Text()

		fileScanner.Scan()
		elf2 = fileScanner.Text()

		fileScanner.Scan()
		elf3 = fileScanner.Text()

		commonItem := findCommonItem(elf1, elf2, elf3)
		sum += calcPrio(commonItem)
	}

	fmt.Println(sum)
}

func errCheck(err error) {
	if err != nil {
		panic("Error found")
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

func findCommonItem(elf1 string, elf2 string, elf3 string) string {
	elf1_splice := strings.Split(elf1, "")
	elf2_splice := strings.Split(elf2, "")
	elf3_splice := strings.Split(elf3, "")

	for _, item1 := range elf1_splice {
		for _, item2 := range elf2_splice {
			if item1 == item2 {
				for _, item3 := range elf3_splice {
					if item1 == item3 {
						return item3
					}
				}
			}
		}
	}

	panic("No common element")
}
