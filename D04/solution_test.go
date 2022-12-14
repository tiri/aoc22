package D04

import (
	"aoc22"
	"testing"
)

func TestPart1(t *testing.T) {
	aoc22.RunTest(t, "known.txt", CampCleanupPart1, 2)
	aoc22.RunTest(t, "challenge.txt", CampCleanupPart1, 494)
}

func TestPart2(t *testing.T) {
	aoc22.RunTest(t, "known.txt", CampCleanupPart2, 4)
	aoc22.RunTest(t, "challenge.txt", CampCleanupPart2, 833)
}
