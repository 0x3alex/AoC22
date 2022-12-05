package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func dayFive() {
	boxes := [][]string{
		{"D", "L", "J", "R", "V", "G", "F"},
		{"T", "P", "M", "B", "V", "H", "J", "S"},
		{"V", "H", "M", "F", "D", "G", "P", "C"},
		{"M", "D", "P", "N", "G", "Q"},
		{"J", "L", "H", "N", "F"},
		{"N", "F", "V", "Q", "D", "G", "T", "Z"},
		{"F", "D", "B", "L"},
		{"M", "J", "B", "S", "V", "D", "N"},
		{"G", "L", "D"},
	}
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err.Error())
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		args := strings.Split(scanner.Text(), " ")
		amount, _ := strconv.Atoi(args[1])
		from, _ := strconv.Atoi(args[3])
		from--
		to, _ := strconv.Atoi(args[5])
		to--
		fmt.Println(boxes)
		fmt.Printf("move %d from %d to %d\n", amount, from, to)

		//part one
		/*for i := 0; i < amount; i++ {
			fromVal := boxes[from][len(boxes[from])-1]
			boxes[to] = append(boxes[to], fromVal)
			boxes[from] = boxes[from][:len(boxes[from])-1]

		}*/
		//part two
		s := boxes[from][len(boxes[from])-amount:]
		for _, i := range s {
			boxes[to] = append(boxes[to], i)
		}
		boxes[from] = boxes[from][:len(boxes[from])-amount]
	}
	//print top of each
	for _, i := range boxes {
		fmt.Printf(i[len(i)-1])
	}

}
