package aoc22

import (
	"bufio"
	"io"
	"math"
	"sort"
)

type PointOnMap struct {
	height rune
	edges  []*PointOnMap
}

func containsEdgeInRoute(route []*PointOnMap, edge *PointOnMap) bool {
	for _, testPoint := range route {
		if testPoint == edge {
			return true
		}
	}
	return false
}

type PrioQueue struct {
	cost int
	node *PointOnMap
}

func findInPrioQueue(queue []PrioQueue, node *PointOnMap) *PrioQueue {
	for _, item := range queue {
		if item.node == node {
			return &item
		}
	}
	return nil
}

func uniformCostSearch(startPoint *PointOnMap, endPoint *PointOnMap) int {
	frontier := []PrioQueue{{node: startPoint, cost: 0}}
	expanded := map[*PointOnMap]struct{}{}
	var exists = struct{}{}
	for len(frontier) > 0 {
		cur := frontier[0]
		frontier = frontier[1:]
		if cur.node == endPoint {
			return cur.cost
		}
		expanded[cur.node] = exists
		for _, edge := range cur.node.edges {
			if !(edge.height-cur.node.height <= 1) {
				continue
			}
			prioQueueItem := findInPrioQueue(frontier, edge)
			nextCost := cur.cost + 1
			if _, ok := expanded[edge]; !ok && prioQueueItem == nil {
				frontier = append(frontier, PrioQueue{node: edge, cost: nextCost})
			} else if prioQueueItem != nil && prioQueueItem.cost > nextCost {
				prioQueueItem.cost = nextCost
			}
		}
		sort.Slice(frontier, func(i, j int) bool {
			return frontier[i].cost < frontier[j].cost
		})
	}
	return -1
}

func HillClimbingAlgorithmPart1(input io.Reader) int {
	_, startPoint, endPoint := parseHeightMap(input)
	return uniformCostSearch(startPoint, endPoint)
}

func HillClimbingAlgorithmPart2(input io.Reader) int {
	heightMap, _, endPoint := parseHeightMap(input)

	fewestCost := math.MaxInt
	for _, row := range heightMap {
		for _, item := range row {
			if item.height == 'a' {
				cost := uniformCostSearch(item, endPoint)
				if cost != -1 && cost < fewestCost {
					fewestCost = cost
				}
			}
		}
	}

	return fewestCost
}

func parseHeightMap(input io.Reader) ([][]*PointOnMap, *PointOnMap, *PointOnMap) {
	reader := bufio.NewReader(input)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	var heightMap [][]*PointOnMap

	for scanner.Scan() {
		var heightMapRow []*PointOnMap
		for _, height := range scanner.Text() {
			pointOnMap := PointOnMap{height: height}
			heightMapRow = append(heightMapRow, &pointOnMap)
		}
		heightMap = append(heightMap, heightMapRow)
	}

	var startPoint *PointOnMap
	var endPoint *PointOnMap
	for rowIdx, row := range heightMap {
		for colIdx, cur := range row {
			if cur.height == 'S' {
				startPoint = cur
				startPoint.height = 'a'
			}
			if cur.height == 'E' {
				endPoint = cur
				endPoint.height = 'z'
				continue
			}
			//top
			if rowIdx-1 >= 0 {
				cur.edges = append(cur.edges, heightMap[rowIdx-1][colIdx])
			}
			//right
			if colIdx+1 < len(heightMap[0]) {
				cur.edges = append(cur.edges, heightMap[rowIdx][colIdx+1])
			}
			//bottom
			if rowIdx+1 < len(heightMap) {
				cur.edges = append(cur.edges, heightMap[rowIdx+1][colIdx])
			}
			//left
			if colIdx-1 >= 0 {
				cur.edges = append(cur.edges, heightMap[rowIdx][colIdx-1])
			}
		}
	}

	return heightMap, startPoint, endPoint
}
