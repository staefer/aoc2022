package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	var threeMostLoadedElfs = make([]int, 3)
	threeMostLoadedElfs[0] = scanAndSumNextElf(fileScanner)
	threeMostLoadedElfs[1] = scanAndSumNextElf(fileScanner)
	threeMostLoadedElfs[2] = scanAndSumNextElf(fileScanner)
	sort.Ints(threeMostLoadedElfs)

	var nextElf int
	for {
		nextElf = scanAndSumNextElf(fileScanner)

		if nextElf == -1 {
			break
		}

		ifLargerThanLeastLoadedThenBubbleElf(threeMostLoadedElfs, nextElf)
	}

	topThreeSum := 0
	for _, elf := range threeMostLoadedElfs {
		topThreeSum += elf
	}

	fmt.Println("Sum of top three loaded elfs: ", topThreeSum)
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
	return -1 // All elfs are scanned
}

func ifLargerThanLeastLoadedThenBubbleElf(loadedElfs []int, possiblyLoadedElf int) {
	if len(loadedElfs) > 3 {
		os.Exit(-1)
	}

	if possiblyLoadedElf > loadedElfs[0] {
		addAndSortLoadedElf(loadedElfs, possiblyLoadedElf)
	}
}

func addAndSortLoadedElf(loadedElfs []int, loadedElf int) {
	loadedElfs[0] = loadedElf
	sort.Ints(loadedElfs)
}
