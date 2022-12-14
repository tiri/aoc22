package D01

import (
	"aoc22"
	"testing"
)

func TestPart1(t *testing.T) {
	aoc22.RunTest(t, "known.txt", CalorieCountingPart1, 24000)
	aoc22.RunTest(t, "challenge.txt", CalorieCountingPart1, 69281)
}

func TestPart2(t *testing.T) {
	aoc22.RunTest(t, "known.txt", CalorieCountingPart2, 45000)
	aoc22.RunTest(t, "challenge.txt", CalorieCountingPart2, 201524)
}
