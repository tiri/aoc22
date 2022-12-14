package D02

import (
	"aoc22"
	"testing"
)

func TestPart1(t *testing.T) {
	aoc22.RunTest(t, "known.txt", RockPaperScissorsPart1, 15)
	aoc22.RunTest(t, "challenge.txt", RockPaperScissorsPart1, 9177)
}

func TestPart2(t *testing.T) {
	aoc22.RunTest(t, "known.txt", RockPaperScissorsPart2, 12)
	aoc22.RunTest(t, "challenge.txt", RockPaperScissorsPart2, 12111)
}
