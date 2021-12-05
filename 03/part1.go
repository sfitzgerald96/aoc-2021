package main

import (
	"strconv"
	"strings"
)

func part1(data string) int64 {
	binaries := strings.Split(data, "\n")
	counts := make([]int, len(binaries[0]))

	for _, binary := range binaries {
		for i, bit := range binary {
			num := int(bit - '0')
			counts[i] += num
		}
	}

	var gammaBits string
	var epsilonBits string
	for _, count := range counts {
		if count > len(binaries)-count {
			gammaBits += "1"
			epsilonBits += "0"
		} else {
			gammaBits += "0"
			epsilonBits += "1"
		}
	}
	gammaInt, _ := strconv.ParseInt(gammaBits, 2, len(gammaBits)+1)
	epsilonInt, _ := strconv.ParseInt(epsilonBits, 2, len(epsilonBits)+1)
	return gammaInt * epsilonInt
}
