package aoc22

import (
	"io"
	"strconv"
	"testing"
)

func TestDay08Part1(t *testing.T) {
	testFunc := func(input io.Reader) string {
		got := TreetopTreeHousePart1(input)
		return strconv.Itoa(got)
	}
	GoldenTest(t, "08_1_known.golden", testFunc)
	GoldenTest(t, "08_1_challenge.golden", testFunc)
}

func TestDay08Part2(t *testing.T) {
	testFunc := func(input io.Reader) string {
		got := TreetopTreeHousePart2(input)
		return strconv.Itoa(got)
	}
	GoldenTest(t, "08_2_known.golden", testFunc)
	GoldenTest(t, "08_2_challenge.golden", testFunc)
}
