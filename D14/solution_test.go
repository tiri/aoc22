package D14

import (
	"aoc22"
	"testing"
)

func TestPart1(t *testing.T) {
	aoc22.RunTest(t, "known.txt", RegolithReservoirPart1, 24)
	aoc22.RunTest(t, "challenge.txt", RegolithReservoirPart1, 901)
}

func TestPart2(t *testing.T) {
	aoc22.RunTest(t, "known.txt", RegolithReservoirPart2, 93)
	aoc22.RunTest(t, "challenge.txt", RegolithReservoirPart2, 24589)
}
