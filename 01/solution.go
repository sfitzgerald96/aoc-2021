package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func readFile(path string) string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic("Unable to read file: " + err.Error())
	} else {
		return string(data)
	}
}

func part1(data string) int {
	depths := strings.Split(data, "\n")
	count := 0
	for i := 1; i < len(depths); i++ {
		current, _ := strconv.Atoi(depths[i])
		previous, _ := strconv.Atoi(depths[i-1])

		if current > previous {
			count++
		}
	}
	return count
}

func part2(data string) int {
	depth_strings := strings.Split(data, "\n")
	var depths []int
	for i, depth := range depth_strings {
		depths[i], err = strconv.Atoi(depth)
		if err != nil {
			panic("Failed to cast string to int: " + err.Error())
		}
	}

	var prev_window []int
	var num_times_increased int

	for i := 1; i < len(depths); i++ {
		current_window := []int{depths[i]}

		fmt.Println(current_window)
		fmt.Println(depths)
	}

	return 1
}

func main() {
	data := readFile("data.txt")
	fmt.Println(part1(data))
}
