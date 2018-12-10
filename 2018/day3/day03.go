package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const inputFile = "./2018/day3/input.txt"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getInput() (input string) {
	bytes, err := ioutil.ReadFile(inputFile)
	check(err)
	input = strings.TrimSpace(string(bytes))
	return
}

func assignment01(input string) (sqFeetOverlap int) {
	fabric := make(map[int]map[int]int, 1000)
	for x := 0; x <= 1000; x++ {
		fabric[x] = make(map[int]int, 1000)
	}

	for _, line := range strings.Split(string(input), "\n") {
		var id, left, top, width, height int
		fmt.Sscanf(line, "#%d @ %d,%d: %dx%d", &id, &left, &top, &width, &height)

		for x := left; x < left+width; x++ {
			for y := top; y < top+height; y++ {
				fabric[x][y]++
				if fabric[x][y] == 2 {
					sqFeetOverlap++
				}
			}
		}
	}
	return
}

func assignment02(input string) (id string) {
	return
}

func main() {
	input := getInput()

	fmt.Printf("The solution to Assignment 1 is: %d\n", assignment01(input))
	fmt.Printf("The solution to Assignment 2 is: %s\n", assignment02(input))
}
