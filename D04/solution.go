package D04

import (
	"bufio"
	"fmt"
	"io"
)

type SectionRange struct {
	start int
	end   int
}

func (p SectionRange) len() int {
	return p.end - p.start
}

func CampCleanupPart1(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	sum := 0
	for scanner.Scan() {
		var firstElf, secondElf SectionRange
		_, err := fmt.Sscanf(scanner.Text(), "%d-%d,%d-%d", &firstElf.start, &firstElf.end, &secondElf.start, &secondElf.end)
		if err != nil {
			return 0
		}

		if firstElf.len() < secondElf.len() {
			tmp := firstElf
			firstElf = secondElf
			secondElf = tmp
		}

		if firstElf.start <= secondElf.start && firstElf.end >= secondElf.end {
			sum++
		}
	}

	return sum
}

func CampCleanupPart2(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	sum := 0
	for scanner.Scan() {
		var firstElf, secondElf SectionRange
		_, err := fmt.Sscanf(scanner.Text(), "%d-%d,%d-%d", &firstElf.start, &firstElf.end, &secondElf.start, &secondElf.end)
		if err != nil {
			return 0
		}

		if firstElf.len() < secondElf.len() {
			tmp := firstElf
			firstElf = secondElf
			secondElf = tmp
		}

		if firstElf.start <= secondElf.start && firstElf.end >= secondElf.start ||
			firstElf.start <= secondElf.end && firstElf.end >= secondElf.end {
			sum++
		}
	}
	return sum
}
