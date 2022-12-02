package aoc22

import (
	"io"
	"strconv"
	"testing"
)

func TestDay02Part1(t *testing.T) {
	testFunc := func(input io.Reader) string {
		got := TotalScoreWithElfStrategy(input)
		return strconv.Itoa(got)
	}
	GoldenTest(t, "02_1_known.golden", testFunc)
	GoldenTest(t, "02_1_challenge.golden", testFunc)
}

func TestDay02Part2(t *testing.T) {
	testFunc := func(input io.Reader) string {
		got := TotalScoreWithElfStrategyWinLose(input)
		return strconv.Itoa(int(got))
	}
	GoldenTest(t, "02_2_known.golden", testFunc)
	GoldenTest(t, "02_2_challenge.golden", testFunc)
}
