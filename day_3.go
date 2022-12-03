package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func getAlphaNum(i int32) int32 {
	//check if lowercase
	end := 90
	top := 52
	if i >= 97 && i <= 122 {
		end = 122
		top = 26
	}
	r := int32(end) - i
	return int32(top) - r
}

// only part two, i deleted part one :/
func dayThree() {
	var (
		i     int
		total int
	)
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	r := bufio.NewReader(f)
	shared := make([]rune, 100)
	for {
		var (
			err   error
			lines [3]string
		)

		for i := 0; i < 3; i++ {
			var l []byte
			l, _, err = r.ReadLine()
			lines[i] = string(l)
		}
		if err == io.EOF {
			break
		}
		counts := make(map[rune]int)
		for _, l := range lines {
			var items = make(map[rune]any)
			for _, c := range l {
				items[(c)] = nil
			}
			for c, _ := range items {
				counts[c] += 1
			}
		}
		for c, count := range counts {
			if count == 3 {
				shared[i] = c
			}
		}
		i++
	}
	for _, r := range shared {
		total += int(getAlphaNum(r))
	}
	fmt.Println(total)
}
