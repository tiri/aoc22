package D06

import (
	"bufio"
	"io"
	"strings"
)

func TuningProblemPart1(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	stream := scanner.Text()

	for i := 3; i < len(stream); i++ {
		if areDuplicateRunesInString(stream[i-3 : i+1]) {
			return i + 1
		}
	}
	return 0
}

func TuningProblemPart2(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	stream := scanner.Text()

	for i := 13; i < len(stream); i++ {
		if areDuplicateRunesInString(stream[i-13 : i+1]) {
			return i + 1
		}
	}
	return 0
}

func areDuplicateRunesInString(sub string) bool {
	for index, char := range sub {
		if strings.ContainsRune(sub[index+1:], char) {
			return false
		}
	}
	return true
}
