package D11

import (
	"aoc22"
	"testing"
)

func TestPart1(t *testing.T) {
	aoc22.RunTest(t, "known.txt", MonkeyInTheMiddlePart1, 10605)
	aoc22.RunTest(t, "challenge.txt", MonkeyInTheMiddlePart1, 50616)
}

func TestPart2(t *testing.T) {
	aoc22.RunTest(t, "known.txt", MonkeyInTheMiddlePart2, 2713310158)
	aoc22.RunTest(t, "challenge.txt", MonkeyInTheMiddlePart2, 11309046332)
}
