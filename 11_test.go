package aoc22

import (
	"testing"
)

func TestDay11Part1(t *testing.T) {
	RunTest(t, "11_known.txt", MonkeyInTheMiddlePart1, 10605)
	RunTest(t, "11_challenge.txt", MonkeyInTheMiddlePart1, 50616)
}

func TestDay11Part2(t *testing.T) {
	RunTest(t, "11_known.txt", MonkeyInTheMiddlePart2, 2713310158)
	RunTest(t, "11_challenge.txt", MonkeyInTheMiddlePart2, 11309046332)
}
