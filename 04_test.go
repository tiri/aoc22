package aoc22

import (
	"testing"
)

func TestDay04Part1(t *testing.T) {
	RunTest(t, "04_known.txt", CampCleanupPart1, 2)
	RunTest(t, "04_challenge.txt", CampCleanupPart1, 494)
}

func TestDay04Part2(t *testing.T) {
	RunTest(t, "04_known.txt", CampCleanupPart2, 4)
	RunTest(t, "04_challenge.txt", CampCleanupPart2, 833)
}
