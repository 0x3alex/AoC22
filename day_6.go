package main

import (
	"fmt"
	"os"
)

func daySix() {
	//packetLen := 4 - for part one
	packetLen := 14
	var buff string
	var buffStart int
	m := make(map[int32]int, 0)
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err.Error())
	}
	str := string(bytes)
	for {
		if buffStart+packetLen > len(str) {
			break
		}
		for _, k := range buff {
			_, ok := m[k]
			if ok { //double
				//reset all
				buffStart++
				buff = ""
				for k := range m {
					delete(m, k)
				}
				break
			}
			m[k] = 1

		}
		if len(buff) == packetLen {
			break
		}
		buff = str[buffStart : buffStart+packetLen]
		fmt.Println(buff)
	}
	fmt.Println("")
	fmt.Println(buffStart + len(buff))
}
