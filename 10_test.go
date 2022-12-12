package aoc22

import (
	"testing"
)

func TestDay10Part1(t *testing.T) {
	RunTest(t, "10_known.txt", CathodeRayTubePart1, 13140)
	RunTest(t, "10_challenge.txt", CathodeRayTubePart1, 14320)
}

func TestDay10Part2(t *testing.T) {
	RunTest(t, "10_known.txt", CathodeRayTubePart2,
		`##..##..##..##..##..##..##..##..##..##..
###...###...###...###...###...###...###.
####....####....####....####....####....
#####.....#####.....#####.....#####.....
######......######......######......####
#######.......#######.......#######.....`,
	)
	RunTest(t, "10_challenge.txt", CathodeRayTubePart2,
		`###...##..###..###..#..#..##..###....##.
#..#.#..#.#..#.#..#.#.#..#..#.#..#....#.
#..#.#....#..#.###..##...#..#.#..#....#.
###..#....###..#..#.#.#..####.###.....#.
#....#..#.#....#..#.#.#..#..#.#....#..#.
#.....##..#....###..#..#.#..#.#.....##..`,
	)
}
