package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const priority = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	priorityMap := makePriorityMap(strings.Split(priority, ""))
	//fmt.Println(priorityMap)

	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	duplicateItems := make(map[string]bool)

	for scanner.Scan() {
		rucksack := strings.Split(scanner.Text(), "")
		//fmt.Println(rucksack)

		fisrtCompartment := rucksack[:len(rucksack)/2]
		secondCompartment := rucksack[len(rucksack)/2:]
		//fmt.Println(fisrtCompartment)
		//fmt.Println(secondCompartment)

		for _, firstItem := range fisrtCompartment { //find items duplicated in each compartment
			for _, secondItem := range secondCompartment {
				if firstItem == secondItem {
					duplicateItems[firstItem] = true
				}
			}
		}
		//break
	}
	//fmt.Println(duplicateItems)
	fmt.Println("Sum of priorities = ", getSumofPriorities(duplicateItems, priorityMap))

}
func makePriorityMap(priority []string) map[string]int {
	priorityMap := make(map[string]int)

	for index, item := range priority {
		priorityMap[item] = index + 1
	}
	return priorityMap
}
func getSumofPriorities(duplicates map[string]bool, priorityMap map[string]int) int {
	total := 0
	for key := range duplicates {
		total += priorityMap[key]
	}
	return total
}
