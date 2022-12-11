package aoc22

import (
	"io"
	"strconv"
	"testing"
)

func TestDay11Part1(t *testing.T) {
	testFunc := func(input io.Reader) string {
		got := MonkeyInTheMiddlePart1(input)
		return strconv.Itoa(got)
	}
	GoldenTest(t, "11_1_known.golden", testFunc)
	GoldenTest(t, "11_1_challenge.golden", testFunc)
}

func TestDay11Part2(t *testing.T) {
	testFunc := func(input io.Reader) string {
		got := MonkeyInTheMiddlePart2(input)
		return strconv.Itoa(got)
	}
	GoldenTest(t, "11_2_known.golden", testFunc)
	GoldenTest(t, "11_2_challenge.golden", testFunc)
}
