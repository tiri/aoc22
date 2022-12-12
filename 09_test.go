package aoc22

import (
	"testing"
)

func TestDay09Part1(t *testing.T) {
	RunTest(t, "09_1_known.txt", RopeBridgePart1, 13)
	RunTest(t, "09_challenge.txt", RopeBridgePart1, 6018)
}

func TestDay09Part2(t *testing.T) {
	RunTest(t, "09_2_known.txt", RopeBridgePart2, 36)
	RunTest(t, "09_challenge.txt", RopeBridgePart2, 2619)
}
