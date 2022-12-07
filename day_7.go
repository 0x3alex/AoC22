package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type fs struct {
	name   string
	size   int
	child  []*fs
	mother *fs
}

func isCmd(s string) bool {
	return s[0] == '$'
}

func (f *fs) getTotalSize(res *int) {
	if f.child != nil {
		for _, v := range f.child {
			v.getTotalSize(res)
		}
	}
	*res += f.size
}

var res []int

func (f *fs) getSizes() {
	var r int
	for _, v := range f.child {
		v.getTotalSize(&r)
		res = append(res, r)
		r = 0
		v.getSizes()
	}
}

func daySeven() {
	file, e := os.Open("input.txt")
	f := fs{}
	currDir := &f
	if e != nil {
		panic(e.Error())
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		args := strings.Split(scanner.Text(), " ")
		if isCmd(scanner.Text()) {
			switch args[1] {
			case "cd":
				switch args[2] {
				case "..":
					currDir = currDir.mother
					fmt.Println("going back a dir")
					fmt.Printf("Current dir is now %s\n", currDir.name)
				default:
					fmt.Printf("Creating new struct with name %s\n", args[2])
					tmp := new(fs)
					tmp.name = args[2]
					tmp.mother = currDir
					currDir.child = append(currDir.child, tmp)
					currDir = tmp
					fmt.Println("moving up a dir")
				}
				//move
			}
			continue
		}
		//args[0] is filesize
		fmt.Println(scanner.Text())
		num, _ := strconv.Atoi(args[0])
		currDir.size += num
	}
	/* part one
	start := &f
	start.getSizes()
	fmt.Println(res)
	var r int
	for _, v := range res {
		if v <= 100000 {
			r += v
		}
	}*/
	/*part two */
	updateSize := 30000000
	var totalSize int
	f.child[0].getTotalSize(&totalSize)
	start := &f
	start.getSizes()
	var r []int
	for _, v := range res {
		if (70000000 - (totalSize - v)) >= updateSize {
			r = append(r, v)
		}
	}
	sort.Ints(r)
	fmt.Println(r[0])

}
