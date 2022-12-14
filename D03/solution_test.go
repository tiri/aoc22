package D03

import (
	"aoc22"
	"testing"
)

func TestPart1(t *testing.T) {
	aoc22.RunTest(t, "known.txt", RucksackReorganizationPart1, 157)
	aoc22.RunTest(t, "challenge.txt", RucksackReorganizationPart1, 7878)
}

func TestPart2(t *testing.T) {
	aoc22.RunTest(t, "known.txt", RucksackReorganizationPart2, 70)
	aoc22.RunTest(t, "challenge.txt", RucksackReorganizationPart2, 2760)
}
