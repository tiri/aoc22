package aoc22

import (
	"bufio"
	"io"
	"strings"
)

func offsetRune(char rune) int {
	if char > 64 && char < 91 {
		return int(char) - 65 + 27
	}
	if char > 96 && char < 123 {
		return int(char) - 97 + 1
	}
	return 0
}

func RucksackReorganizationCompartment(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	sum := 0
	for scanner.Scan() {
		rucksack := scanner.Text()
		firstCompartment := rucksack[:len(rucksack)/2]
		secondCompartment := rucksack[len(rucksack)/2:]
		for _, char := range firstCompartment {
			if strings.IndexRune(secondCompartment, char) < 0 {
				continue
			}
			sum += offsetRune(char)
			break
		}
	}

	return sum
}

func RucksackReorganizationBadges(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	sum := 0
	var elfGroup []string
	for scanner.Scan() {
		rucksack := scanner.Text()
		elfGroup = append(elfGroup, rucksack)

		if len(elfGroup) != 3 {
			continue
		}

		var matchedFirst []rune

		for _, char := range elfGroup[0] {
			if strings.IndexRune(elfGroup[1], char) < 0 {
				continue
			}

			matchedFirst = append(matchedFirst, char)
		}

		for _, char := range elfGroup[2] {
			if strings.IndexRune(string(matchedFirst), char) < 0 {
				continue
			}
			sum += offsetRune(char)
			break
		}

		elfGroup = nil
	}

	return sum
}
