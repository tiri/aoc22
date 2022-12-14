package D09

import (
	"bufio"
	"fmt"
	"io"
)

type PointXY struct {
	x int
	y int
}

func (point *PointXY) add(pointToAdd PointXY) {
	point.x += pointToAdd.x
	point.y += pointToAdd.y
}

func (point *PointXY) sub(pointToSub PointXY) {
	point.x -= pointToSub.x
	point.y -= pointToSub.y
}

func (point *PointXY) clamp(min int, max int) {
	if point.x < min {
		point.x = min
	}
	if point.y < min {
		point.y = min
	}
	if point.x > max {
		point.x = max
	}
	if point.y > max {
		point.y = max
	}
}

func (point *PointXY) moveTowards(target PointXY) {
	diff := PointXY{target.x - point.x, target.y - point.y}
	if diff.x > 1 || diff.x < -1 || diff.y > 1 || diff.y < -1 {
		diff.clamp(-1, 1)
		point.add(diff)
	}
}

var DIRECTIONS = map[rune]PointXY{
	'L': {-1, 0},
	'R': {+1, 0},
	'U': {0, -1},
	'D': {0, +1},
}

func RopeBridgePart1(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	head, tail := PointXY{}, PointXY{}

	tailTrace := map[PointXY]bool{}

	for scanner.Scan() {
		direction, steps := ' ', 0
		_, err := fmt.Sscanf(scanner.Text(), "%c %d", &direction, &steps)
		if err != nil {
			panic(err)
		}

		move := DIRECTIONS[direction]

		for i := 0; i < steps; i++ {
			head.add(move)
			tail.moveTowards(head)
			tailTrace[tail] = true
		}
	}

	return len(tailTrace)
}

const TailLength = 9

func RopeBridgePart2(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	var head PointXY

	knots := [TailLength]PointXY{}
	knot9Trace := map[PointXY]bool{}

	for scanner.Scan() {
		direction, steps := ' ', 0
		_, err := fmt.Sscanf(scanner.Text(), "%c %d", &direction, &steps)
		if err != nil {
			panic(err)
		}

		move := DIRECTIONS[direction]

		for i := 0; i < steps; i++ {
			head.add(move)

			prev := &head

			for i := 0; i < len(knots); i++ {
				knots[i].moveTowards(*prev)
				prev = &knots[i]
			}

			knot9Trace[knots[8]] = true
		}
	}

	return len(knot9Trace)
}
