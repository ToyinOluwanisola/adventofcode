package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	readFromFile()
}

func readFromFile() {
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	elfSlice := []int{}
	calories := 0

	for scanner.Scan() {
		value := scanner.Text()
		if len(value) != 0 {
			n, err := strconv.Atoi(value)
			if err != nil {
				fmt.Println("error = ", err)
				break
			} else {
				//fmt.Println("text = ", value)
				calories += n
			}
		} else {
			elfSlice = append(elfSlice, calories)
			//fmt.Println("calories = ", calories)
			calories = 0
		}
	}
	elfSlice = append(elfSlice, calories) //add last elf's calories

	//fmt.Println("calories  ", elfSlice)
	// sort.Slice(elfSlice, func(i, j int) bool {
	// 	return elfSlice[i] > elfSlice[j] // > = asending / < = descending
	// })
	// fmt.Println("Largest calories  ", elfSlice[0])
	// totalSum := elfSlice[0] + elfSlice[1] + elfSlice[2]

	// for _, v := range elfSlice {
	// 	fmt.Println(v)
	// }

	sort.Ints(elfSlice)
	countCalories := len(elfSlice)
	fmt.Println("Largest calories  ", elfSlice[countCalories-1])
	totalSum := elfSlice[countCalories-1] + elfSlice[countCalories-2] + elfSlice[countCalories-3]

	fmt.Println("Total of top 3 calories  ", totalSum)

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

}
