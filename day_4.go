package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func dayFour() {
	var count int
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err.Error())
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		//split into two
		args := strings.Split(scanner.Text(), ",")
		//first half
		num1, _ := strconv.Atoi(strings.Split(args[0], "-")[0])
		num2, _ := strconv.Atoi(strings.Split(args[0], "-")[1])

		//second half
		num3, _ := strconv.Atoi(strings.Split(args[1], "-")[0])
		num4, _ := strconv.Atoi(strings.Split(args[1], "-")[1])

		//check if fully contained - part one
		/*if (num1 <= num3 && num2 >= num4) || (num1 >= num3 && num2 <= num4) {
			fmt.Println(scanner.Text())
			count++
		}*/
		//check if overlap - part two
		if (num1 >= num3 && num1 <= num4) || (num2 >= num3 && num2 <= num4) ||
			(num3 >= num1 && num3 <= num2) || (num4 >= num1 && num4 <= num2) {
			count++
		}
	}
	fmt.Println(count)
}
