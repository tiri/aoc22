package aoc22

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type Command string

const (
	noop Command = "noop"
	addx Command = "addx"
)

func (c Command) String() string {
	return string(c)
}

func parseInstruction(input string) Instruction {
	words := strings.Fields(input)
	switch words[0] {
	case noop.String():
		return Instruction{noop, 0, 1}
	case addx.String():
		val, err := strconv.Atoi(words[1])
		if err != nil {
			return Instruction{}
		}
		return Instruction{addx, val, 2}
	}
	return Instruction{}
}

type Instruction struct {
	command Command
	value   int
	cycles  int
}

func (i *Instruction) execute(registerX *int) {
	i.cycles--

	if i.cycles == 0 {
		switch i.command {
		case noop:
		case addx:
			*registerX += i.value
		}
	}
}

func (i *Instruction) hasBeenExecuted() bool {
	return i.cycles < 1
}

func CathodeRayTubePart1(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	registerX := 1
	totalCycles := 0
	signalStrength := 0
	var instruction Instruction
	for {
		if instruction.hasBeenExecuted() {
			if !scanner.Scan() {
				break
			}
			instruction = parseInstruction(scanner.Text())
		}

		totalCycles++

		if (totalCycles-20)%40 == 0 {
			signalStrength += totalCycles * registerX
		}

		instruction.execute(&registerX)
	}

	return signalStrength
}

func CathodeRayTubePart2(input io.Reader) string {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	registerX := 1
	totalCycles := 0
	crt := ""
	var instruction Instruction
	for {
		if instruction.hasBeenExecuted() {
			if !scanner.Scan() {
				break
			}
			instruction = parseInstruction(scanner.Text())
		}

		crtPos := totalCycles % 40
		if crtPos >= registerX-1 && crtPos <= registerX+1 {
			crt += "#"
		} else {
			crt += "."
		}

		if crtPos+1 >= 40 {
			crt += "\n"
		}

		instruction.execute(&registerX)
		totalCycles++
	}

	return crt
}
