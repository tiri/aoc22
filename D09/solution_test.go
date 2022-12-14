package D09

import (
	"aoc22"
	"testing"
)

func TestPart1(t *testing.T) {
	aoc22.RunTest(t, "known_1.txt", RopeBridgePart1, 13)
	aoc22.RunTest(t, "challenge.txt", RopeBridgePart1, 6018)
}

func TestPart2(t *testing.T) {
	aoc22.RunTest(t, "known_2.txt", RopeBridgePart2, 36)
	aoc22.RunTest(t, "challenge.txt", RopeBridgePart2, 2619)
}
