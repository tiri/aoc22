package D08

import (
	"aoc22"
	"testing"
)

func TestPart1(t *testing.T) {
	aoc22.RunTest(t, "known.txt", TreetopTreeHousePart1, 21)
	aoc22.RunTest(t, "challenge.txt", TreetopTreeHousePart1, 1690)
}

func TestPart2(t *testing.T) {
	aoc22.RunTest(t, "known.txt", TreetopTreeHousePart2, 8)
	aoc22.RunTest(t, "challenge.txt", TreetopTreeHousePart2, 535680)
}
