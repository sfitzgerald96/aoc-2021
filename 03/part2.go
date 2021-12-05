package main

import (
	"fmt"
	"strings"
)

func part2(data string) int {
	binaries := strings.Split(data, "\n")
	current_binaries := binaries

	for i := 0; i < len(binaries[0]); i++ {
		if len(current_binaries) == 1 {
			break
		}
		current_binaries = findRemainingBinaries(current_binaries, i)
	}

	fmt.Println("current_binaries:", current_binaries)
	return 0
}

func findRemainingBinaries(binaries []string, cursor int) []string {
	var zero_binaries_at_cursor []string
	var one_binaries_at_cursor []string

	for _, binary := range binaries {
		if string(binary[cursor]) == "1" {
			one_binaries_at_cursor = append(one_binaries_at_cursor, binary)
		} else {
			zero_binaries_at_cursor = append(zero_binaries_at_cursor, binary)
		}
	}

	if len(one_binaries_at_cursor) < len(zero_binaries_at_cursor) {
		return zero_binaries_at_cursor
	} else {
		return one_binaries_at_cursor
	}
}
