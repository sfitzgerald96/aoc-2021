package main

import (
	"fmt"
	"io/ioutil"
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
	data := readFile("sample.txt")
	fmt.Println("part 1 solution:", part1(data))
	fmt.Println("-------------")
	fmt.Println(part2(data))
}
