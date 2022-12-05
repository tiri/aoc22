package aoc22

import (
	"io"
	"testing"
)

func TestDay05Part1(t *testing.T) {
	testFunc := func(input io.Reader) string {
		got := SupplyStacksPart1(input)
		return got
	}
	GoldenTest(t, "05_1_known.golden", testFunc)
	GoldenTest(t, "05_1_challenge.golden", testFunc)
}

func TestDay05Part2(t *testing.T) {
	testFunc := func(input io.Reader) string {
		got := SupplyStacksPart2(input)
		return got
	}
	GoldenTest(t, "05_2_known.golden", testFunc)
	GoldenTest(t, "05_2_challenge.golden", testFunc)
}
