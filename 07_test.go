package aoc22

import (
	"io"
	"strconv"
	"testing"
)

func TestDay07Part1(t *testing.T) {
	testFunc := func(input io.Reader) string {
		got := NoSpaceLeftOnDevicePart1(input)
		return strconv.Itoa(got)
	}
	GoldenTest(t, "07_1_known.golden", testFunc)
	GoldenTest(t, "07_1_challenge.golden", testFunc)
}

func TestDay07Part2(t *testing.T) {
	testFunc := func(input io.Reader) string {
		got := NoSpaceLeftOnDevicePart2(input)
		return strconv.Itoa(got)
	}
	GoldenTest(t, "07_2_known.golden", testFunc)
	GoldenTest(t, "07_2_challenge.golden", testFunc)
}
