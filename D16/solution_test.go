package D14

import (
	"aoc22"
	"testing"
)

func TestPart1(t *testing.T) {
	aoc22.RunTest(t, "known.txt", ProboscideaVolcaniumPart1, 1651)
	aoc22.RunTest(t, "challenge.txt", ProboscideaVolcaniumPart1, 1828)
}

func TestPart2(t *testing.T) {
	aoc22.RunTest(t, "known.txt", ProboscideaVolcaniumPart2, 1707)
	aoc22.RunTest(t, "challenge.txt", ProboscideaVolcaniumPart2, 2292)
}
