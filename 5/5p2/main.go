package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	stacks := make([][]string, 9)
	stacks[0] = []string{"R", "N", "F", "V", "L", "J", "S", "M"}
	stacks[1] = []string{"P", "N", "D", "Z", "F", "J", "W", "H"}
	stacks[2] = []string{"W", "R", "C", "D", "G"}
	stacks[3] = []string{"N", "B", "S"}
	stacks[4] = []string{"M", "Z", "W", "P", "C", "B", "F", "N"}
	stacks[5] = []string{"P", "R", "M", "W"}
	stacks[6] = []string{"R", "T", "N", "G", "L", "S", "W"}
	stacks[7] = []string{"Q", "T", "H", "F", "N", "B", "V"}
	stacks[8] = []string{"L", "M", "H", "Z", "N", "F"}

	fileScanner := openLineScanner("input.txt")

	for fileScanner.Scan() {
		boxAmount, fromStack, toStack := splitLine(fileScanner.Text())
		moveBoxes(stacks, (fromStack - 1), (toStack - 1), boxAmount)
	}

	for i := 0; i < 9; i++ {
		fmt.Print(stacks[i][len(stacks[i])-1])
	}
}

func openLineScanner(fileName string) *bufio.Scanner {
	filePath, err := os.Open(fileName)
	if err != nil {
		panic("Can't open file")
	}
	fileScanner := bufio.NewScanner(filePath)
	fileScanner.Split(bufio.ScanLines)
	return fileScanner
}

func moveBoxes(all_stacks [][]string, fromStackIndex int, toStackIndex int, boxAmount int) {
	fromStackLenght := len(all_stacks[fromStackIndex])
	fromStackLastIndex := fromStackLenght - 1
	movedBoxes := all_stacks[fromStackIndex][fromStackLastIndex-boxAmount+1:]
	all_stacks[fromStackIndex] = all_stacks[fromStackIndex][:fromStackLastIndex-boxAmount+1]
	all_stacks[toStackIndex] = append(all_stacks[toStackIndex], movedBoxes...)
}

func splitLine(line string) (int, int, int) {
	words := strings.Split(line, " ")
	boxes, _ := strconv.Atoi(words[1])
	from, _ := strconv.Atoi(words[3])
	to, _ := strconv.Atoi(words[5])
	return boxes, from, to
}
