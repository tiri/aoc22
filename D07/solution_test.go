package D07

import (
	"aoc22"
	"testing"
)

func TestPart1(t *testing.T) {
	aoc22.RunTest(t, "known.txt", NoSpaceLeftOnDevicePart1, 95437)
	aoc22.RunTest(t, "challenge.txt", NoSpaceLeftOnDevicePart1, 1667443)
}

func TestPart2(t *testing.T) {
	aoc22.RunTest(t, "known.txt", NoSpaceLeftOnDevicePart2, 24933642)
	aoc22.RunTest(t, "challenge.txt", NoSpaceLeftOnDevicePart2, 8998590)
}
