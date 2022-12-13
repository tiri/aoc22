package aoc22

import (
	"testing"
)

func TestDay13Part1(t *testing.T) {
	RunTest(t, "13_known.txt", DistressSignalPart1, 13)
	RunTest(t, "13_challenge.txt", DistressSignalPart1, 5905)
}

func TestDay13Part2(t *testing.T) {
	RunTest(t, "13_known.txt", DistressSignalPart2, 140)
	RunTest(t, "13_challenge.txt", DistressSignalPart2, 21691)
}
