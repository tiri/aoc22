package aoc22

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
)

type monkeyCalc func(int) int

type Monkey struct {
	items           []int
	calcFunc        monkeyCalc
	inspectedItems  int
	divisibleBy     int
	ifTrueToMonkey  int
	ifFalseToMonkey int
}

func MonkeyInTheMiddlePart1(input io.Reader) int {
	reader := bufio.NewReader(input)

	monkeys := parseMonkey(reader)

	for i := 0; i < 20; i++ {
		for monkeyIdx := range monkeys {
			monkey := &monkeys[monkeyIdx]
			for _, item := range monkey.items {
				result := monkey.calcFunc(item)
				result /= 3
				var targetMonkey *Monkey
				if result%monkey.divisibleBy == 0 {
					targetMonkey = &monkeys[monkey.ifTrueToMonkey]
				} else {
					targetMonkey = &monkeys[monkey.ifFalseToMonkey]
				}
				targetMonkey.items = append(targetMonkey.items, result)
				monkey.inspectedItems++
			}
			monkey.items = []int{}
		}
	}

	var inspectedItemSlice []int
	for _, monkey := range monkeys {
		inspectedItemSlice = append(inspectedItemSlice, monkey.inspectedItems)
	}
	sort.Ints(inspectedItemSlice)

	return inspectedItemSlice[len(inspectedItemSlice)-1] * inspectedItemSlice[len(inspectedItemSlice)-2]
}

func MonkeyInTheMiddlePart2(input io.Reader) int {
	reader := bufio.NewReader(input)

	monkeys := parseMonkey(reader)

	gcm := 1
	for _, monkey := range monkeys {
		gcm *= monkey.divisibleBy
	}

	for i := 0; i < 10000; i++ {
		for monkeyIdx := range monkeys {
			monkey := &monkeys[monkeyIdx]
			for _, item := range monkey.items {
				result := monkey.calcFunc(item)
				var targetMonkey *Monkey
				if result%monkey.divisibleBy == 0 {
					targetMonkey = &monkeys[monkey.ifTrueToMonkey]
				} else {
					targetMonkey = &monkeys[monkey.ifFalseToMonkey]
				}
				result %= gcm
				targetMonkey.items = append(targetMonkey.items, result)
				monkey.inspectedItems++
			}
			monkey.items = []int{}
		}
	}

	var inspectedItemSlice []int
	for _, monkey := range monkeys {
		inspectedItemSlice = append(inspectedItemSlice, monkey.inspectedItems)
	}
	sort.Ints(inspectedItemSlice)

	return inspectedItemSlice[len(inspectedItemSlice)-1] * inspectedItemSlice[len(inspectedItemSlice)-2]
}

func parseMonkey(reader *bufio.Reader) []Monkey {
	var monkeys []Monkey
	for {
		var monkeyNumber, divisibleBy, ifTrueToMonkey, ifFalseToMonkey int
		var startingItemsStr, operationFirstElement, operationSecondElement string
		var operatorRune rune
		_, err := fmt.Fscanf(reader,
			"Monkey %d:\n"+
				"  Starting items: %s\n"+
				"  Operation: new = %s %c %s\n"+
				"  Test: divisible by %d\n"+
				"    If true: throw to monkey %d\n"+
				"    If false: throw to monkey %d",
			&monkeyNumber,
			&startingItemsStr,
			&operationFirstElement, &operatorRune, &operationSecondElement,
			&divisibleBy,
			&ifTrueToMonkey,
			&ifFalseToMonkey,
		)

		if err != nil {
			panic(err)
		}

		var items []int
		for _, itemStr := range strings.Split(startingItemsStr, ",") {
			item, err := strconv.Atoi(itemStr)
			if err != nil {
				panic(err)
			}
			items = append(items, item)
		}

		var operation monkeyCalc

		ops := map[rune]func(int, int) int{
			'+': func(a, b int) int { return a + b },
			'*': func(a, b int) int { return a * b },
		}

		if operationFirstElement == "old" && operationSecondElement == "old" {
			operation = func(old int) int {
				return ops[operatorRune](old, old)
			}
		} else if operationFirstElement == "old" {
			secondElement, err := strconv.Atoi(operationSecondElement)
			if err != nil {
				panic(err)
			}
			operation = func(old int) int {
				return ops[operatorRune](old, secondElement)
			}
		} else {
			firstElement, err := strconv.Atoi(operationFirstElement)
			if err != nil {
				panic(err)
			}
			operation = func(old int) int {
				return ops[operatorRune](firstElement, old)
			}
		}

		monkey := Monkey{
			items:           items,
			inspectedItems:  0,
			divisibleBy:     divisibleBy,
			ifTrueToMonkey:  ifTrueToMonkey,
			ifFalseToMonkey: ifFalseToMonkey,
			calcFunc:        operation,
		}

		monkeys = append(monkeys, monkey)

		_, _, err = reader.ReadRune()
		_, _, err = reader.ReadRune()
		if err != nil {
			break
		}
	}
	return monkeys
}
