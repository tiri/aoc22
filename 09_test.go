package aoc22

import (
	"io"
	"strconv"
	"testing"
)

func TestDay09Part1(t *testing.T) {
	testFunc := func(input io.Reader) string {
		got := RopeBridgePart1(input)
		return strconv.Itoa(got)
	}
	GoldenTest(t, "09_1_known.golden", testFunc)
	GoldenTest(t, "09_1_challenge.golden", testFunc)
}

func TestDay09Part2(t *testing.T) {
	testFunc := func(input io.Reader) string {
		got := RopeBridgePart2(input)
		return strconv.Itoa(got)
	}
	GoldenTest(t, "09_2_known.golden", testFunc)
	GoldenTest(t, "09_2_challenge.golden", testFunc)
}
