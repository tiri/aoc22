package D05

import (
	"aoc22"
	"testing"
)

func TestPart1(t *testing.T) {
	aoc22.RunTest(t, "known.txt", SupplyStacksPart1, "CMZ")
	aoc22.RunTest(t, "challenge.txt", SupplyStacksPart1, "ZWHVFWQWW")
}

func TestPart2(t *testing.T) {
	aoc22.RunTest(t, "known.txt", SupplyStacksPart2, "MCD")
	aoc22.RunTest(t, "challenge.txt", SupplyStacksPart2, "HZFZCCWWV")
}
