package D14

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"strings"
)

type Valve struct {
	flowRate int
	leadsTo  map[string]int
}

func (v Valve) hasLeadsTo(check string) bool {
	_, found := v.leadsTo[check]
	return found
}

func (v Valve) filterLeadsTo(filter map[string]Valve) Valve {
	leadsTo := map[string]int{}
	for name, weight := range v.leadsTo {
		for filterName := range filter {
			if filterName == name {
				leadsTo[name] = weight
			}
		}
	}
	v.leadsTo = leadsTo
	return v
}

type HasBeenOpened []string

func (o HasBeenOpened) contains(check string) bool {
	for _, a := range o {
		if a == check {
			return true
		}
	}
	return false
}

func (o HasBeenOpened) add(add string) HasBeenOpened {
	return append(o, add)
}

type Valves map[string]Valve

func (valves Valves) addAllLeadsTo() Valves {
	distances := map[string]map[string]int{}
	threshold := math.MaxInt / len(valves)

	for iName, iValve := range valves {
		distances[iName] = map[string]int{}
		for jName := range valves {
			if iName == jName {
				distances[iName][jName] = 0
			} else if iValve.hasLeadsTo(jName) {
				distances[iName][jName] = 1
			} else {
				distances[iName][jName] = threshold
			}

		}
	}

	for kName := range valves {
		for iName := range valves {
			for jName := range valves {
				if distances[iName][kName]+distances[kName][jName] < distances[iName][jName] {
					distances[iName][jName] = distances[iName][kName] + distances[kName][jName]
				}
			}
		}
	}

	newValves := Valves{}
	for iName, iValve := range valves {
		leadsTo := map[string]int{}
		for jName := range valves {
			leadsTo[jName] = distances[iName][jName]
		}
		newValves[iName] = Valve{
			flowRate: iValve.flowRate,
			leadsTo:  leadsTo,
		}
	}

	return newValves

}

func (valves Valves) filterAllValvesThatCanBeOpened() Valves {
	valvesThatCanBeOpened := Valves{}

	for name, valve := range valves {
		if !(valve.flowRate > 0) {
			continue
		}

		valvesThatCanBeOpened[name] = valve
	}
	return valvesThatCanBeOpened
}

func (valves Valves) filterLeadsTo(filter map[string]Valve) Valves {
	for name, valve := range valves {
		valves[name] = valve.filterLeadsTo(filter)
	}
	return valves
}

func ProboscideaVolcaniumPart1(reader io.Reader) int {
	valves := parse(reader)

	valves = valves.addAllLeadsTo()

	valvesThatCanBeOpened := valves.filterAllValvesThatCanBeOpened()
	valvesThatCanBeOpened = valvesThatCanBeOpened.filterLeadsTo(valvesThatCanBeOpened)
	valvesThatCanBeOpened["AA"] = valves["AA"].filterLeadsTo(valvesThatCanBeOpened)

	sum := runX(valvesThatCanBeOpened, "AA", 0, 30+1, nil)
	return sum
}

func ProboscideaVolcaniumPart2(reader io.Reader) int {
	valves := parse(reader)

	valves = valves.addAllLeadsTo()

	valvesThatCanBeOpened := valves.filterAllValvesThatCanBeOpened()
	valvesThatCanBeOpened = valvesThatCanBeOpened.filterLeadsTo(valvesThatCanBeOpened)

	best := 0
	for i := 0; i < (1<<len(valvesThatCanBeOpened))/2; i++ {
		humanValves := Valves{}
		elephantValves := Valves{}
		j := 0
		for name, elem := range valvesThatCanBeOpened {
			if i&(1<<j) > 0 {
				humanValves[name] = elem
			} else {
				elephantValves[name] = elem
			}
			j++
		}

		humanValves.filterLeadsTo(humanValves)
		elephantValves.filterLeadsTo(elephantValves)

		humanValves["AA"] = valves["AA"].filterLeadsTo(humanValves)
		elephantValves["AA"] = valves["AA"].filterLeadsTo(elephantValves)

		sum := runX(humanValves, "AA", 0, 26+1, HasBeenOpened{})
		sum += runX(elephantValves, "AA", 0, 26+1, HasBeenOpened{})
		if sum > best {
			best = sum
		}
	}

	return best
}

func runX(valves Valves, currentValveName string, weight int, timeLeft int, hasBeenOpened HasBeenOpened) int {
	timeLeft -= weight + 1

	if timeLeft < 1 {
		return 0
	}

	if len(hasBeenOpened) == len(valves) {
		return 0
	}

	currentValve := valves[currentValveName]

	hasBeenOpened = hasBeenOpened.add(currentValveName)
	calc := timeLeft * currentValve.flowRate

	bestResult := 0
	for leadsToValve, leadsToWeight := range currentValve.leadsTo {
		if hasBeenOpened.contains(leadsToValve) {
			continue
		}
		result := runX(valves, leadsToValve, leadsToWeight, timeLeft, hasBeenOpened)

		if result > bestResult {
			bestResult = result
		}
	}
	return bestResult + calc
}

func parse(reader io.Reader) Valves {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	valves := Valves{}

	for scanner.Scan() {
		var name, leadsToStr string
		var flowRate int
		var nilStr string

		line := scanner.Text()
		_, err := fmt.Sscanf(line,
			"Valve %s has flow rate=%d; %s %s to %s %s",
			&name, &flowRate, &nilStr, &nilStr, &nilStr, &leadsToStr)
		if err != nil {
			panic(err)
		}

		leadsTo := map[string]int{}
		for _, leadsToName := range strings.Split(leadsToStr, ",") {
			leadsTo[leadsToName] = 1
		}

		valves[name] = Valve{
			flowRate: flowRate,
			leadsTo:  leadsTo,
		}
	}

	return valves
}
