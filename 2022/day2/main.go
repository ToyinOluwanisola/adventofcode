package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	rockScore     = 1
	paperScore    = 2
	scissorsScore = 3
	win           = 6
	draw          = 3
	lost          = 0
)

func main() {
	partTwo()
}

func partOne() {
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalScore := 0

	for scanner.Scan() {
		value := scanner.Text()
		results := strings.Split(value, " ")
		score := 0

		switch results[0] {
		case "A": // Rock
			if results[1] == "X" { //rock=draw
				score += rockScore + draw
			} else if results[1] == "Y" { //paper=win
				score += paperScore + win
			} else if results[1] == "Z" { //scissors=lost
				score += scissorsScore + lost
			}
		case "B": // Paper
			if results[1] == "X" { //rock=lost
				score += rockScore + lost
			} else if results[1] == "Y" { //paper=draw
				score += paperScore + draw
			} else if results[1] == "Z" { //scissors=win
				score += scissorsScore + win
			}
		case "C": // Scissors
			if results[1] == "X" { //rock=win
				score += rockScore + win
			} else if results[1] == "Y" { //paper=lost
				score += paperScore + lost
			} else if results[1] == "Z" { //scissors=draw
				score += scissorsScore + draw
			}
		}
		totalScore += score
		//fmt.Println("score is ", score)
	}

	fmt.Println("total score is ", totalScore)
}

func partTwo() {
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalScore := 0

	for scanner.Scan() {
		value := scanner.Text()
		results := strings.Split(value, " ")
		score := 0

		switch results[0] {
		case "A": // Rock
			if results[1] == "X" { //lose =scissors
				score += scissorsScore + lost
			} else if results[1] == "Y" { //rock=draw
				score += rockScore + draw
			} else if results[1] == "Z" { //paper=win
				score += paperScore + win
			}
		case "B": // Paper
			if results[1] == "X" { //rock=lost
				score += rockScore + lost
			} else if results[1] == "Y" { //paper=draw
				score += paperScore + draw
			} else if results[1] == "Z" { //scissors=win
				score += scissorsScore + win
			}
		case "C": // Scissors
			if results[1] == "X" { //paper=lose
				score += paperScore + lost
			} else if results[1] == "Y" { //scissors=draw
				score += scissorsScore + draw
			} else if results[1] == "Z" { //rock=win
				score += rockScore + win
			}
		}
		totalScore += score
		//fmt.Println("score is ", score)
	}

	fmt.Println("total score is ", totalScore)
}
