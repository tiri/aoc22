package aoc22

import (
	"testing"
)

func TestDay07Part1(t *testing.T) {
	RunTest(t, "07_known.txt", NoSpaceLeftOnDevicePart1, 95437)
	RunTest(t, "07_challenge.txt", NoSpaceLeftOnDevicePart1, 1667443)
}

func TestDay07Part2(t *testing.T) {
	RunTest(t, "07_known.txt", NoSpaceLeftOnDevicePart2, 24933642)
	RunTest(t, "07_challenge.txt", NoSpaceLeftOnDevicePart2, 8998590)
}
