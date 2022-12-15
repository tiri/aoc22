package D14

import (
	"aoc22"
	"bufio"
	"fmt"
	"io"
)

type Type int

const (
	Sensor Type = iota
	Beacon
	CannotContainBeacon
)

type Position struct {
	x int
	y int
}

type Cave map[Position]Type

func (cave Cave) isPositionOccupied(pos Position) bool {
	_, ok := cave[pos]
	return ok
}

type MDSensor map[Position]int

type VectorX struct {
	from int
	to   int
}

func BeaconExclusionZonePart1(reader io.Reader, rowY int) int {

	cave, mdToBeacon := parse(reader)

	for pos, kind := range cave {
		if kind != Sensor {
			continue
		}

		md := mdToBeacon[pos]

		diff := aoc22.Abs(pos.y - rowY)
		if diff > md {
			continue
		}

		runXMaxLen := md - diff
		for runX := pos.x - runXMaxLen; runX <= pos.x+runXMaxLen; runX++ {
			possiblePos := Position{x: runX, y: rowY}
			if !cave.isPositionOccupied(possiblePos) {
				cave[possiblePos] = CannotContainBeacon
			}
		}
	}

	sum := 0
	for _, kind := range cave {
		if kind != CannotContainBeacon {
			continue
		}
		sum++
	}

	return sum
}

func BeaconExclusionZonePart2(reader io.Reader, maxXY int) int {

	cave, mdToBeacon := parse(reader)

	for y := 0; y < maxXY; y++ {

		var xVectors []VectorX
		for pos, kind := range cave {
			if kind != Sensor {
				continue
			}

			md := mdToBeacon[pos]

			diff := aoc22.Abs(pos.y - y)
			if diff > md {
				continue
			}

			runXMaxLen := md - diff

			xVectors = append(xVectors, VectorX{from: pos.x - runXMaxLen, to: pos.x + runXMaxLen})
		}

		runX := 0
		for {
			found := false
			for _, vector := range xVectors {
				if vector.from <= runX && vector.to >= runX {
					runX = vector.to + 1
					found = true
					break
				}
			}
			if runX > maxXY {
				break
			}
			if found {
				continue
			}
			return runX*4000000 + y
		}

	}

	return -1
}

func parse(reader io.Reader) (Cave, MDSensor) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	cave := Cave{}
	md := MDSensor{}

	for scanner.Scan() {
		line := scanner.Text()
		sensorXY := Position{}
		beaconXY := Position{}
		_, err := fmt.Sscanf(line,
			"Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d",
			&sensorXY.x, &sensorXY.y, &beaconXY.x, &beaconXY.y)
		if err != nil {
			panic(err)
		}

		cave[sensorXY] = Sensor
		cave[beaconXY] = Beacon
		md[sensorXY] = aoc22.Abs(sensorXY.x-beaconXY.x) + aoc22.Abs(sensorXY.y-beaconXY.y)
	}

	return cave, md
}
