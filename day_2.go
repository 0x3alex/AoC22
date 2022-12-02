package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func points(p string) int {
	switch p {
	case "X", "A":
		return 1
	case "Y", "B":
		return 2
	case "Z", "C":
		return 3
	}
	return 0
}

func hasWon(p1, p2 string) bool {
	return (p1 == "A" && p2 == "Y") ||
		(p1 == "B" && p2 == "Z") ||
		(p1 == "C" && p2 == "X")
}

func isDraw(p1, p2 string) bool {
	return (p1 == "A" && p2 == "X") ||
		(p1 == "B" && p2 == "Y") ||
		(p1 == "C" && p2 == "Z")
}

func chooseLooseCounter(p1 string) int {
	switch p1 {
	case "A":
		return points("C") //Scissors
	case "B":
		return points("A")
	case "C":
		return points("B")
	}
	return 0
}

func chooseWonCounter(p1 string) int {
	switch p1 {
	case "A":
		return points("B")
	case "B":
		return points("C")
	case "C":
		return points("A")
	}
	return 0
}

func dayTwo() {
	var score int
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err.Error())
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		txt := strings.Split(scanner.Text(), " ")

		//part one
		/*if hasWon(txt[0], txt[1]) {
			score += points(txt[1]) + 6
		} else if isDraw(txt[0], txt[1]) {
			score += points(txt[1]) + 3
		} else {
			score += points(txt[1])
		}*/

		//part two
		switch txt[1] {
		case "X":
			score += chooseLooseCounter(txt[0])
		case "Y":
			score += points(txt[0]) + 3
		case "Z":
			score += chooseWonCounter(txt[0]) + 6
		}

	}
	fmt.Println(score)
}
