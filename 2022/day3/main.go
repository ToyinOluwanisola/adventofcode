package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const priority = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var priorityDuplicateItems = []string{}

func main() {
	priorityMap := makePriorityMap(strings.Split(priority, ""))
	//fmt.Println(priorityMap)

	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		rucksack := strings.Split(scanner.Text(), "")
		//fmt.Println(rucksack)

		fisrtCompartment := rucksack[:len(rucksack)/2]
		secondCompartment := rucksack[len(rucksack)/2:]
		//fmt.Println(fisrtCompartment)
		//fmt.Println(secondCompartment)

		duplicateItems := make(map[string]bool)
		for _, firstItem := range fisrtCompartment { //find items duplicated in each compartment
			for _, secondItem := range secondCompartment {
				if firstItem == secondItem {
					duplicateItems[firstItem] = true
				}
			}
		}
		appendDuplicates(duplicateItems)
		//break
	}
	//fmt.Println(priorityDuplicateItems)
	fmt.Println("Sum of priorities = ", getSumofPriorities(priorityDuplicateItems, priorityMap))

}
func makePriorityMap(priority []string) map[string]int {
	priorityMap := make(map[string]int)

	for index, item := range priority {
		priorityMap[item] = index + 1
	}
	return priorityMap
}

func appendDuplicates(duplicates map[string]bool) {
	for key := range duplicates {
		priorityDuplicateItems = append(priorityDuplicateItems, key)
	}

}
func getSumofPriorities(duplicates []string, priorityMap map[string]int) int {
	total := 0
	for _, key := range duplicates {
		total += priorityMap[key]
	}
	return total
}
