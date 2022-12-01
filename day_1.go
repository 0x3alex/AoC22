package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func dayOne() {
	var current int
	calories := make([]int, 0)
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err.Error())
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		//blank line
		if scanner.Text() == "" {
			calories = append(calories, current)
			current = 0
			continue
		}
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err.Error())
		}
		current += num
	}
	sort.Ints(calories)

	//result part one
	fmt.Println(calories[len(calories)-1])

	//result part two
	var res int
	for i := 1; i <= 3; i++ {
		res += calories[len(calories)-i]
	}
	fmt.Println(res)
}
