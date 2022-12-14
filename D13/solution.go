package D13

import (
	"bufio"
	"encoding/json"
	"io"
	"math"
)

type CompareResult int

const (
	RightOrder CompareResult = iota
	WrongOrder
	UndeterminedOrder
)

func isSlice(input interface{}) bool {
	switch input.(type) {
	case []interface{}:
		return true
	}
	return false
}

func DistressSignalPart1(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	sum := 0
	index := 1
	for scanner.Scan() {
		firstRowStr := scanner.Text()

		if len(firstRowStr) == 0 {
			continue
		}

		if !scanner.Scan() {
			panic("didn't find second row")
		}
		secondRowStr := scanner.Text()

		leftTransmission := buildDistressSignalStruct(firstRowStr)
		rightTransmission := buildDistressSignalStruct(secondRowStr)
		if compareDistressSignal(leftTransmission, rightTransmission) == RightOrder {
			sum += index
		}

		index++
	}

	return sum
}

func DistressSignalPart2(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	two := buildDistressSignalStruct("[[2]]")
	six := buildDistressSignalStruct("[[6]]")

	twoRightOrder, sixRightOrder := 1, 1

	for scanner.Scan() {
		transmissionStr := scanner.Text()

		if len(transmissionStr) == 0 {
			continue
		}

		transmission := buildDistressSignalStruct(transmissionStr)

		if compareDistressSignal(transmission, two) == RightOrder {
			twoRightOrder++
		}
		if compareDistressSignal(transmission, six) == RightOrder {
			sixRightOrder++
		}
	}

	return twoRightOrder * (sixRightOrder + 1)
}

func compareDistressSignal(left interface{}, right interface{}) CompareResult {
	isLeftSlice, isRightSlice := isSlice(left), isSlice(right)

	if !isLeftSlice && isRightSlice {
		return compareDistressSignal([]interface{}{left}, right)
	} else if isLeftSlice && !isRightSlice {
		return compareDistressSignal(left, []interface{}{right})
	}

	leftSlice, rightSlice := left.([]interface{}), right.([]interface{})
	minLen := int(math.Min(float64(len(leftSlice)), float64(len(rightSlice))))

	for i := 0; i < minLen; i++ {
		leftItem, rightItem := leftSlice[i], rightSlice[i]

		if isSlice(leftItem) || isSlice(rightItem) {
			if result := compareDistressSignal(leftItem, rightItem); result != UndeterminedOrder {
				return result
			}
			continue
		}

		if leftItem.(float64) > rightItem.(float64) {
			return WrongOrder
		}

		if leftItem.(float64) < rightItem.(float64) {
			return RightOrder
		}
	}

	if len(rightSlice) == len(leftSlice) {
		return UndeterminedOrder
	}

	if len(rightSlice) == minLen && len(leftSlice) != minLen {
		return WrongOrder
	}

	return RightOrder
}

func buildDistressSignalStruct(str string) []interface{} {
	var output []interface{}
	err := json.Unmarshal([]byte(str), &output)
	if err != nil {
		panic("can't not parse input: " + str)
	}
	return output
}
