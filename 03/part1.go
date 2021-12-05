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

	var gamma_bits string
	var epsilon_bits string
	for _, count := range counts {
		if count > len(binaries)-count {
			gamma_bits += "1"
			epsilon_bits += "0"
		} else {
			gamma_bits += "0"
			epsilon_bits += "1"
		}
	}
	gamma_int, _ := strconv.ParseInt(gamma_bits, 2, len(gamma_bits)+1)
	epsilon_int, _ := strconv.ParseInt(epsilon_bits, 2, len(epsilon_bits)+1)
	return gamma_int * epsilon_int
}
