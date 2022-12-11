package aoc22

import (
	"io"
	"strconv"
	"testing"
)

func TestDay10Part1(t *testing.T) {
	testFunc := func(input io.Reader) string {
		got := CathodeRayTubePart1(input)
		return strconv.Itoa(got)
	}
	GoldenTest(t, "10_1_known.golden", testFunc)
	GoldenTest(t, "10_1_challenge.golden", testFunc)
}

func TestDay10Part2(t *testing.T) {
	testFunc := func(input io.Reader) string {
		got := CathodeRayTubePart2(input)
		return got
	}
	GoldenTest(t, "10_2_known.golden", testFunc)
	GoldenTest(t, "10_2_challenge.golden", testFunc)
}
