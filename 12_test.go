package aoc22

import (
	"io"
	"strconv"
	"testing"
)

func TestDay12Part1(t *testing.T) {
	testFunc := func(input io.Reader) string {
		got := HillClimbingAlgorithmPart1(input)
		return strconv.Itoa(got)
	}
	GoldenTest(t, "12_1_known.golden", testFunc)
	GoldenTest(t, "12_1_challenge.golden", testFunc)
}

func TestDay12Part2(t *testing.T) {
	testFunc := func(input io.Reader) string {
		got := HillClimbingAlgorithmPart2(input)
		return strconv.Itoa(got)
	}
	GoldenTest(t, "12_2_known.golden", testFunc)
	GoldenTest(t, "12_2_challenge.golden", testFunc)
}
