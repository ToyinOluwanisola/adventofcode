package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	partOne()
}

func partOne() {
	totalContainedPairs := 0
	totalOverlappedPairs := 0
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println("Error reading file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if isRangeFullyContained(strings.Split(scanner.Text(), ",")) {
			totalContainedPairs++
		}
		if isOverlapped(strings.Split(scanner.Text(), ",")) {
			totalOverlappedPairs++
		}
	}
	fmt.Println("fully contained assignment pairs: ", totalContainedPairs)
	fmt.Println("overlapped pairs: ", totalOverlappedPairs)
}

func isRangeFullyContained(pairs []string) bool {
	isRangeContained := false
	elfOne := getElfsSections(pairs[0])
	elfTwo := getElfsSections(pairs[1])

	duplicatedSections := make(map[int]bool)

	if len(elfOne) >= len(elfTwo) {
		for _, elfOneSection := range elfOne {
			for _, elfTwoSection := range elfTwo {
				if elfOneSection == elfTwoSection {
					duplicatedSections[elfOneSection] = true
				}
			}
		}
		if len(duplicatedSections) == len(elfTwo) {
			isRangeContained = true
			//fmt.Println("elft 1 = ", elfOne, "elf 2 =", elfTwo)
			//fmt.Println(duplicatedSections)
		}

	} else {
		for _, elfTwoSection := range elfTwo {
			for _, elfOneSection := range elfOne {
				if elfTwoSection == elfOneSection {
					duplicatedSections[elfTwoSection] = true
				}
			}
		}
		if len(duplicatedSections) == len(elfOne) {
			isRangeContained = true
			//fmt.Println("elft 1 = ", elfOne, "elf 2 =", elfTwo)
			//fmt.Println(duplicatedSections)
		}
	}

	return isRangeContained

}

func isOverlapped(pairs []string) bool {
	isOverlapped := false
	elfOne := getElfsSections(pairs[0])
	elfTwo := getElfsSections(pairs[1])

	if len(elfOne) >= len(elfTwo) {
		for _, elfOneSection := range elfOne {
			for _, elfTwoSection := range elfTwo {
				if elfOneSection == elfTwoSection {
					isOverlapped = true
					break
				}
			}
		}
	} else {
		for _, elfTwoSection := range elfTwo {
			for _, elfOneSection := range elfOne {
				if elfTwoSection == elfOneSection {
					isOverlapped = true
					break
				}
			}
		}
	}

	return isOverlapped

}

func getElfsSections(sections string) []int {
	assignment := strings.Split(sections, "-")
	result := make([]int, len(assignment))

	for i := range result {
		result[i], _ = strconv.Atoi(assignment[i])
	}

	elfSections := []int{}
	index := 0
	for section := result[0]; section < result[1]+1; section++ {
		elfSections = append(elfSections, section)
		index++
	}
	return elfSections

}
