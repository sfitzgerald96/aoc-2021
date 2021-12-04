package main

import (
	"io/ioutil"
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

func main() {
	data := readFile("data.txt")
	movements := strings.Split(data, "\n")
	part1(movements)
	part2(movements)
}
