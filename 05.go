package aoc22

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func SupplyStacksPart1(input io.Reader) string {
	scanner := bufio.NewScanner(input)
	scanner.Split(SplitAtEmptyRow)

	scanner.Scan()
	stack := parseStack(scanner.Text())

	scanner.Scan()
	reader := strings.NewReader(scanner.Text())

	var amount, from, to int
	var popped rune
	for {
		_, err := fmt.Fscanf(reader, "move %d from %d to %d\n", &amount, &from, &to)
		if err != nil {
			break
		}

		for i := 0; i < amount; i++ {
			popped, stack[from-1] = stack[from-1][len(stack[from-1])-1], stack[from-1][:len(stack[from-1])-1]
			stack[to-1] = append(stack[to-1], popped)
		}
	}

	return getFirstFromEachColumn(stack)
}

func SupplyStacksPart2(input io.Reader) string {
	scanner := bufio.NewScanner(input)
	scanner.Split(SplitAtEmptyRow)

	scanner.Scan()
	stack := parseStack(scanner.Text())

	scanner.Scan()
	reader := strings.NewReader(scanner.Text())

	var amount, from, to int
	var popped []rune
	for {
		_, err := fmt.Fscanf(reader, "move %d from %d to %d\n", &amount, &from, &to)
		if err != nil {
			break
		}

		popped, stack[from-1] = stack[from-1][len(stack[from-1])-amount:], stack[from-1][:len(stack[from-1])-amount]
		stack[to-1] = append(stack[to-1], popped...)
	}

	return getFirstFromEachColumn(stack)
}

func getFirstFromEachColumn(stack [][]rune) string {
	var output string
	for _, runes := range stack {
		output += string(runes[len(runes)-1])
	}
	return output
}

func parseStack(stringStack string) [][]rune {
	scanner := bufio.NewScanner(strings.NewReader(stringStack))
	scanner.Split(bufio.ScanLines)

	var stackRows []string
	for scanner.Scan() {
		stackRows = append(stackRows, scanner.Text())
	}

	var indexRow string
	indexRow, stackRows = stackRows[len(stackRows)-1], stackRows[:len(stackRows)-1]
	numberOfColumns := (len(indexRow) + 2) / 4
	stack := make([][]rune, numberOfColumns)

	for i := 0; i < numberOfColumns; i++ {
		var stackColumn []rune
		runePos := i*4 + 1
		for rowIndex := len(stackRows) - 1; rowIndex >= 0; rowIndex-- {
			if runePos > len(stackRows[rowIndex]) {
				continue
			}
			elems := rune(stackRows[rowIndex][runePos])
			if elems == 32 {
				break
			}
			stackColumn = append(stackColumn, elems)
		}
		stack[i] = stackColumn
	}
	return stack
}
