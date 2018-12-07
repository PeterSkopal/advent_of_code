package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"strings"
)

const inputFile = "./2018/day2/input.txt"

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

func stringDiff(a string, b string) (int, int) {
	missCount := 0
	index := -1

	for i := range a {
		if a[i] != b[i] {
			index = i
			missCount++
		}
	}
	return missCount, index
}

func assignment01(input string) int {
	twos := 0
	threes := 0
	scanner := bufio.NewScanner(strings.NewReader(input))

	for scanner.Scan() {
		foundTwo := false
		foundThree := false

		for _, char := range scanner.Text() {
			if foundTwo && foundThree {
				break
			}
			if strings.Count(scanner.Text(), string(char)) == 2 {
				foundTwo = true
			}
			if strings.Count(scanner.Text(), string(char)) == 3 {
				foundThree = true
			}

		}

		if foundTwo {
			twos++
		}
		if foundThree {
			threes++
		}

	}
	return twos * threes
}

func assignment02(input string) string {
	ids := strings.Split(input, "\n")

	for _, first := range ids {
		for _, second := range ids {
			if diff, index := stringDiff(first, second); diff == 1 && index != 0 {
				return first[:index] + first[index+1:]
			}
		}
	}

	panic("Couldn't find any id matches")
}

func main() {
	input := getInput()

	fmt.Printf("The solution to Assignment 1 is: %d\n", assignment01(input))
	fmt.Printf("The solution to Assignment 2 is: %s\n", assignment02(input))
}
