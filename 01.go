package aoc22

import (
	"bufio"
	"io"
	"sort"
	"strconv"
)

func CalorieCountingPart1(reader io.Reader) uint {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	var maxCalories, currentCalories uint
	maxCalories = 0
	currentCalories = 0
	for scanner.Scan() {
		if len(scanner.Text()) < 1 {
			if maxCalories < currentCalories {
				maxCalories = currentCalories
			}

			currentCalories = 0
			continue
		}
		calories, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return 0
		}
		currentCalories += uint(calories)
	}
	return maxCalories
}

func CalorieCountingPart2(reader io.Reader) uint {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	var caloriesByElf []int
	currentElf := 0
	for scanner.Scan() {
		if len(scanner.Text()) < 1 {
			caloriesByElf = append(caloriesByElf, currentElf)
			currentElf = 0
			continue
		}
		calories, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return 0
		}
		currentElf += calories
	}
	caloriesByElf = append(caloriesByElf, currentElf)
	sort.Sort(sort.Reverse(sort.IntSlice(caloriesByElf)))
	totalCaloriesByTopThree := caloriesByElf[0] + caloriesByElf[1] + caloriesByElf[2]

	return uint(totalCaloriesByTopThree)
}
