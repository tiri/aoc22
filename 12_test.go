package aoc22

import (
	"testing"
)

func TestDay12Part1(t *testing.T) {
	RunTest(t, "12_known.txt", HillClimbingAlgorithmPart1, 31)
	RunTest(t, "12_challenge.txt", HillClimbingAlgorithmPart1, 449)
}

func TestDay12Part2(t *testing.T) {
	RunTest(t, "12_known.txt", HillClimbingAlgorithmPart2, 29)
	RunTest(t, "12_challenge.txt", HillClimbingAlgorithmPart2, 443)
}
