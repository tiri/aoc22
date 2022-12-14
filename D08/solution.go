package D08

import (
	"bufio"
	"io"
	"strconv"
)

func TreetopTreeHousePart1(input io.Reader) int {
	scanner := bufio.NewScanner(input)

	treeMap := buildTreeMap(scanner)

	length := len(treeMap)
	//init visibility map
	var visibilityMap [][]bool
	visibilityMap = make([][]bool, length)
	for i := 0; i < length; i++ {
		visibilityMap[i] = make([]bool, length)
	}

	// top -> bottom
	for x := 0; x < length; x++ {
		tallest := -1
		for y := 0; y < length; y++ {
			if treeMap[y][x] > tallest {
				visibilityMap[y][x] = true
				tallest = treeMap[y][x]
			}
		}
	}

	// right -> left
	for y := 0; y < length; y++ {
		tallest := -1
		for x := length - 1; x > 0; x-- {
			if treeMap[y][x] > tallest {
				visibilityMap[y][x] = true
				tallest = treeMap[y][x]
			}
		}
	}

	// bottom -> top
	for x := 0; x < length; x++ {
		tallest := -1
		for y := length - 1; y > 0; y-- {
			if treeMap[y][x] > tallest {
				visibilityMap[y][x] = true
				tallest = treeMap[y][x]
			}
		}
	}

	// left -> right
	for y := 0; y < length; y++ {
		tallest := -1
		for x := 0; x < length; x++ {
			if treeMap[y][x] > tallest {
				visibilityMap[y][x] = true
				tallest = treeMap[y][x]
			}
		}
	}

	sum := 0
	for _, column := range visibilityMap {
		for _, tree := range column {
			if tree {
				sum++
			}
		}
	}

	return sum
}

func TreetopTreeHousePart2(input io.Reader) int {
	scanner := bufio.NewScanner(input)

	treeMap := buildTreeMap(scanner)

	topScore := 0
	for x := 1; x < len(treeMap)-1; x++ {
		for y := 1; y < len(treeMap)-1; y++ {
			height := treeMap[y][x]
			// left
			left := 0
			for i := x - 1; i >= 0; i-- {
				left++
				if height <= treeMap[y][i] {
					break
				}
			}

			// right
			right := 0
			for i := x + 1; i < len(treeMap); i++ {
				right++
				if height <= treeMap[y][i] {
					break
				}
			}

			// top
			top := 0
			for i := y - 1; i >= 0; i-- {
				top++
				if height <= treeMap[i][x] {
					break
				}
			}

			// bottom
			bottom := 0
			for i := y + 1; i < len(treeMap); i++ {
				bottom++
				if height <= treeMap[i][x] {
					break
				}
			}

			score := left * right * top * bottom
			if score > topScore {
				topScore = score
			}
		}
	}

	return topScore
}

func buildTreeMap(scanner *bufio.Scanner) [][]int {
	var treeMap [][]int
	scanner.Split(bufio.ScanRunes)
	var treeMapRow []int
	for scanner.Scan() {
		if scanner.Text() == "\n" {
			treeMap = append(treeMap, treeMapRow)
			treeMapRow = nil
			continue
		}
		height, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil
		}
		treeMapRow = append(treeMapRow, height)
	}

	if len(treeMap) > 0 {
		treeMap = append(treeMap, treeMapRow)
	}

	return treeMap
}
