package main

import (
	"strconv"
	"strings"
)

func part2(data string) int64 {
	binaries := strings.Split(data, "\n")
	oxygenBinaries := searchBinariesFor("most", binaries)
	co2Binaries := searchBinariesFor("least", binaries)

	oxygen, _ := strconv.ParseInt(oxygenBinaries[0], 2, 64)
	co2, _ := strconv.ParseInt(co2Binaries[0], 2, 64)
	return co2 * oxygen
}

func findRemainingBinaries(binaries []string, cursor int, compare string) []string {
	var zeroBinariesAtCursor []string
	var oneBinariesAtCursor []string

	for _, binary := range binaries {
		if string(binary[cursor]) == "1" {
			oneBinariesAtCursor = append(oneBinariesAtCursor, binary)
		} else {
			zeroBinariesAtCursor = append(zeroBinariesAtCursor, binary)
		}
	}

	if compare == "most" {
		if len(zeroBinariesAtCursor) > len(oneBinariesAtCursor) {
			return zeroBinariesAtCursor
		} else {
			return oneBinariesAtCursor
		}
	} else {
		if len(oneBinariesAtCursor) < len(zeroBinariesAtCursor) {
			return oneBinariesAtCursor
		} else {
			return zeroBinariesAtCursor
		}
	}
}

func searchBinariesFor(compare string, binaries []string) []string {
	currentBinaries := binaries

	for i := 0; i < len(binaries[0]); i++ {
		if len(currentBinaries) == 1 {
			break
		}
		currentBinaries = findRemainingBinaries(currentBinaries, i, compare)
	}

	return currentBinaries
}
