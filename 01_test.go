package aoc22

import (
	"testing"
)

func TestDay01Part1(t *testing.T) {
	RunTest(t, "01_known.txt", CalorieCountingPart1, 24000)
	RunTest(t, "01_challenge.txt", CalorieCountingPart1, 69281)
}

func TestDay01Part2(t *testing.T) {
	RunTest(t, "01_known.txt", CalorieCountingPart2, 45000)
	RunTest(t, "01_challenge.txt", CalorieCountingPart2, 201524)
}
