package aoc22

import (
	"testing"
)

func TestDay08Part1(t *testing.T) {
	RunTest(t, "08_known.txt", TreetopTreeHousePart1, 21)
	RunTest(t, "08_challenge.txt", TreetopTreeHousePart1, 1690)
}

func TestDay08Part2(t *testing.T) {
	RunTest(t, "08_known.txt", TreetopTreeHousePart2, 8)
	RunTest(t, "08_challenge.txt", TreetopTreeHousePart2, 535680)
}
