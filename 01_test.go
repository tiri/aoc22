package aoc22

import (
	"io"
	"strconv"
	"testing"
)

func TestDay01Part1(t *testing.T) {
	testFunc := func(input io.Reader) string {
		got := HowMuchCaloriesHasTheTopElf(input)
		return strconv.Itoa(int(got))
	}
	GoldenTest(t, "01_1_known.golden", testFunc)
	GoldenTest(t, "01_1_challenge.golden", testFunc)
}

func TestDay01Part2(t *testing.T) {
	testFunc := func(input io.Reader) string {
		got := HowMuchCaloriesHaveTheTopThreeElves(input)
		return strconv.Itoa(int(got))
	}
	GoldenTest(t, "01_2_known.golden", testFunc)
	GoldenTest(t, "01_2_challenge.golden", testFunc)
}
