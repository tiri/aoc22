package aoc22

import (
	"testing"
)

func TestDay06Part1(t *testing.T) {
	RunTest(t, "06_known.txt", TuningProblemPart1, 5)
	RunTest(t, "06_challenge.txt", TuningProblemPart1, 1134)
}

func TestDay06Part2(t *testing.T) {
	RunTest(t, "06_known.txt", TuningProblemPart2, 23)
	RunTest(t, "06_challenge.txt", TuningProblemPart2, 2263)
}
