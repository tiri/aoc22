package D12

import (
	"aoc22"
	"testing"
)

func TestPart1(t *testing.T) {
	aoc22.RunTest(t, "known.txt", HillClimbingAlgorithmPart1, 31)
	aoc22.RunTest(t, "challenge.txt", HillClimbingAlgorithmPart1, 449)
}

func TestPart2(t *testing.T) {
	aoc22.RunTest(t, "known.txt", HillClimbingAlgorithmPart2, 29)
	aoc22.RunTest(t, "challenge.txt", HillClimbingAlgorithmPart2, 443)
}
