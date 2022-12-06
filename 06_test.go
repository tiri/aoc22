package aoc22

import (
	"io"
	"strconv"
	"testing"
)

func TestDay06Part1(t *testing.T) {
	testFunc := func(input io.Reader) string {
		got := TuningProblemPart1(input)
		return strconv.Itoa(got)
	}
	GoldenTest(t, "06_1_known.golden", testFunc)
	GoldenTest(t, "06_1_challenge.golden", testFunc)
}

func TestDay06Part2(t *testing.T) {
	testFunc := func(input io.Reader) string {
		got := TuningProblemPart2(input)
		return strconv.Itoa(got)
	}
	GoldenTest(t, "06_2_known.golden", testFunc)
	GoldenTest(t, "06_2_challenge.golden", testFunc)
}
