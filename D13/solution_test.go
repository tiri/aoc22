package D13

import (
	"aoc22"
	"testing"
)

func TestPart1(t *testing.T) {
	aoc22.RunTest(t, "known.txt", DistressSignalPart1, 13)
	aoc22.RunTest(t, "challenge.txt", DistressSignalPart1, 5905)
}

func TestPart2(t *testing.T) {
	aoc22.RunTest(t, "known.txt", DistressSignalPart2, 140)
	aoc22.RunTest(t, "challenge.txt", DistressSignalPart2, 21691)
}
