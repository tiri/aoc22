package aoc22

import (
	"io"
	"strconv"
	"testing"
)

func TestDay03Part1(t *testing.T) {
	testFunc := func(input io.Reader) string {
		got := RucksackReorganizationCompartment(input)
		return strconv.Itoa(got)
	}
	GoldenTest(t, "03_1_known.golden", testFunc)
	GoldenTest(t, "03_1_challenge.golden", testFunc)
}

func TestDay03Part2(t *testing.T) {
	testFunc := func(input io.Reader) string {
		got := RucksackReorganizationBadges(input)
		return strconv.Itoa(got)
	}
	GoldenTest(t, "03_2_known.golden", testFunc)
	GoldenTest(t, "03_2_challenge.golden", testFunc)
}
