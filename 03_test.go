package aoc22

import (
	"testing"
)

func TestDay03Part1(t *testing.T) {
	RunTest(t, "03_known.txt", RucksackReorganizationPart1, 157)
	RunTest(t, "03_challenge.txt", RucksackReorganizationPart1, 7878)
}

func TestDay03Part2(t *testing.T) {
	RunTest(t, "03_known.txt", RucksackReorganizationPart2, 70)
	RunTest(t, "03_challenge.txt", RucksackReorganizationPart2, 2760)
}
