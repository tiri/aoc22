package D06

import (
	"aoc22"
	"testing"
)

func TestPart1(t *testing.T) {
	aoc22.RunTest(t, "known.txt", TuningProblemPart1, 5)
	aoc22.RunTest(t, "challenge.txt", TuningProblemPart1, 1134)
}

func TestPart2(t *testing.T) {
	aoc22.RunTest(t, "known.txt", TuningProblemPart2, 23)
	aoc22.RunTest(t, "challenge.txt", TuningProblemPart2, 2263)
}
