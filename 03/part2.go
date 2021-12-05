package main

import (
	"fmt"
	"strconv"
	"strings"
)

func part2(data string) int64 {
	binaries := strings.Split(data, "\n")
	oxygen_binaries := search_binaries_for("most", binaries)
	co2_binaries := search_binaries_for("least", binaries)
	fmt.Println("oxygen_binaries:", oxygen_binaries)
	fmt.Println("co2_binaries:", co2_binaries)

	oxygen, _ := strconv.ParseInt(oxygen_binaries[0], 2, 64)
	co2, _ := strconv.ParseInt(co2_binaries[0], 2, 64)
	fmt.Println("oxygen:", oxygen)
	fmt.Println("co2:", co2)
	return co2 * oxygen
}

func findRemainingBinaries(binaries []string, cursor int, compare string) []string {
	var zero_binaries_at_cursor []string
	var one_binaries_at_cursor []string

	for _, binary := range binaries {
		if string(binary[cursor]) == "1" {
			one_binaries_at_cursor = append(one_binaries_at_cursor, binary)
		} else {
			zero_binaries_at_cursor = append(zero_binaries_at_cursor, binary)
		}
	}

	if compare == "most" {
		if len(zero_binaries_at_cursor) > len(one_binaries_at_cursor) {
			return zero_binaries_at_cursor
		} else {
			return one_binaries_at_cursor
		}
	} else {
		if len(one_binaries_at_cursor) < len(zero_binaries_at_cursor) {
			return one_binaries_at_cursor
		} else {
			return zero_binaries_at_cursor
		}
	}
}

func search_binaries_for(compare string, binaries []string) []string {
	current_binaries := binaries

	for i := 0; i < len(binaries[0]); i++ {
		if len(current_binaries) == 1 {
			break
		}
		current_binaries = findRemainingBinaries(current_binaries, i, compare)
	}

	return current_binaries
}
