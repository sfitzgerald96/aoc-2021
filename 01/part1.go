package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("01/data.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	depths := strings.Split(string(data), "\n")
	count := 0
	for i := 1; i < len(depths); i++ {
		current, _ := strconv.Atoi(depths[i])
		previous, _ := strconv.Atoi(depths[i-1])

		if current > previous {
			count++
		}
	}

	fmt.Println(count)
}
