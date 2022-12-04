package aoc22

import (
	"io"
	"strconv"
	"testing"
)

func TestDay04Part1(t *testing.T) {
	testFunc := func(input io.Reader) string {
		got := CampCleanupPart1(input)
		return strconv.Itoa(got)
	}
	GoldenTest(t, "04_1_known.golden", testFunc)
	GoldenTest(t, "04_1_challenge.golden", testFunc)
}

func TestDay04Part2(t *testing.T) {
	testFunc := func(input io.Reader) string {
		got := CampCleanupPart2(input)
		return strconv.Itoa(got)
	}
	GoldenTest(t, "04_2_known.golden", testFunc)
	GoldenTest(t, "04_2_challenge.golden", testFunc)
}
