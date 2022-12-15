package D14

import (
	"aoc22"
	"testing"
)

func TestPart1(t *testing.T) {
	aoc22.RunTestWithParam(t, "known.txt", BeaconExclusionZonePart1, 10, 26)
	aoc22.RunTestWithParam(t, "challenge.txt", BeaconExclusionZonePart1, 2000000, 5832528)
}

func TestPart2(t *testing.T) {
	aoc22.RunTestWithParam(t, "known.txt", BeaconExclusionZonePart2, 20, 56000011)
	aoc22.RunTestWithParam(t, "challenge.txt", BeaconExclusionZonePart2, 4000000, 13360899249595)
}
