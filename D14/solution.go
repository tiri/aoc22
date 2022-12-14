package D14

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type Element int

const (
	Sand Element = iota
	Rock
)

type PositionXY struct {
	x int
	y int
}

type Cave map[PositionXY]Element

func (cave Cave) isPositionOccupied(grainOfSand PositionXY) bool {
	_, ok := cave[grainOfSand]
	return ok
}

func (cave Cave) findMaxDepth() int {
	maxDepth := 0
	for xy, _ := range cave {
		if maxDepth < xy.y {
			maxDepth = xy.y
		}
	}
	return maxDepth
}

func RegolithReservoirPart1(reader io.Reader) int {
	sandOrigin := PositionXY{x: 500, y: 0}

	cave := parse(reader)

	maxY := cave.findMaxDepth()

	grainOfSand := sandOrigin

	for {
		if grainOfSand.y == maxY {
			break
		}

		// below
		below := PositionXY{x: grainOfSand.x, y: grainOfSand.y + 1}
		if !cave.isPositionOccupied(below) {
			grainOfSand = below
			continue
		}
		// left below
		leftBelow := PositionXY{x: grainOfSand.x - 1, y: grainOfSand.y + 1}
		if !cave.isPositionOccupied(leftBelow) {
			grainOfSand = leftBelow
			continue
		}
		// right below
		rightBelow := PositionXY{x: grainOfSand.x + 1, y: grainOfSand.y + 1}
		if !cave.isPositionOccupied(rightBelow) {
			grainOfSand = rightBelow
			continue
		}

		// create new grandofsand
		cave[grainOfSand] = Sand
		grainOfSand = sandOrigin
	}

	sum := 0
	for _, element := range cave {
		if element == Sand {
			sum++
		}
	}

	return sum
}

func RegolithReservoirPart2(reader io.Reader) int {
	sandOrigin := PositionXY{x: 500, y: 0}

	cave := parse(reader)

	maxY := cave.findMaxDepth() + 2

	grainOfSand := sandOrigin

	for {
		// below
		below := PositionXY{x: grainOfSand.x, y: grainOfSand.y + 1}
		if !cave.isPositionOccupied(below) && below.y != maxY {
			grainOfSand = below
			continue
		}
		// left below
		leftBelow := PositionXY{x: grainOfSand.x - 1, y: grainOfSand.y + 1}
		if !cave.isPositionOccupied(leftBelow) && leftBelow.y != maxY {
			grainOfSand = leftBelow
			continue
		}
		// right below
		rightBelow := PositionXY{x: grainOfSand.x + 1, y: grainOfSand.y + 1}
		if !cave.isPositionOccupied(rightBelow) && rightBelow.y != maxY {
			grainOfSand = rightBelow
			continue
		}

		// create new grandofsand
		cave[grainOfSand] = Sand

		if grainOfSand.y == sandOrigin.y {
			break
		}

		grainOfSand = sandOrigin
	}

	sum := 0
	for _, element := range cave {
		if element == Sand {
			sum++
		}
	}

	return sum
}

func parse(reader io.Reader) Cave {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	cave := Cave{}

	for scanner.Scan() {
		line := scanner.Text()
		lineScanner := bufio.NewScanner(strings.NewReader(line))
		lineScanner.Split(bufio.ScanWords)
		var from *PositionXY
		for lineScanner.Scan() {
			word := lineScanner.Text()
			if word == "->" {
				continue
			}

			to := &PositionXY{}
			_, err := fmt.Sscanf(word, "%d,%d", &to.x, &to.y)
			if err != nil {
				panic(err)
			}

			if from == nil {
				from = to
				continue
			}

			start := from
			end := to

			// run vertical
			if to.y == from.y {
				if start.x > end.x {
					end = from
					start = to
				}

				for runX := start.x; runX <= end.x; runX++ {
					cave[PositionXY{x: runX, y: start.y}] = Rock
				}
			}

			// run horizontal
			if to.x == from.x {
				if start.y > end.y {
					end = from
					start = to
				}

				for runY := start.y; runY <= end.y; runY++ {
					cave[PositionXY{x: start.x, y: runY}] = Rock
				}
			}

			from = to
		}
	}

	return cave
}
