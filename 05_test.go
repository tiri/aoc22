package aoc22

import (
	"testing"
)

func TestDay05Part1(t *testing.T) {
	RunTest(t, "05_known.txt", SupplyStacksPart1, "CMZ")
	RunTest(t, "05_challenge.txt", SupplyStacksPart1, "ZWHVFWQWW")
}

func TestDay05Part2(t *testing.T) {
	RunTest(t, "05_known.txt", SupplyStacksPart2, "MCD")
	RunTest(t, "05_challenge.txt", SupplyStacksPart2, "HZFZCCWWV")
}
