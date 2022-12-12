package aoc22

import (
	"testing"
)

func TestDay02Part1(t *testing.T) {
	RunTest(t, "02_known.txt", RockPaperScissorsPart1, 15)
	RunTest(t, "02_challenge.txt", RockPaperScissorsPart1, 9177)
}

func TestDay02Part2(t *testing.T) {
	RunTest(t, "02_known.txt", RockPaperScissorsPart2, 12)
	RunTest(t, "02_challenge.txt", RockPaperScissorsPart2, 12111)
}
