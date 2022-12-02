package aoc22

import (
	"bufio"
	"fmt"
	"io"
)

type Option int64
type Result int64

const (
	Rock Option = iota
	Paper
	Scissors
)

const (
	Lost Result = iota
	Draw
	Win
)

func parsePlayerOpponent(input rune) Option {
	switch input {
	case 'A', 'X':
		return Rock
	case 'B', 'Y':
		return Paper
	case 'C', 'Z':
		return Scissors
	}
	return 0
}

func parseResult(input rune) Result {
	switch input {
	case 'X':
		return Lost
	case 'Y':
		return Draw
	case 'Z':
		return Win
	}
	return 0
}

func (option Option) play(opponent Option) Result {
	if option == opponent {
		return Draw
	}

	if option.prev() == opponent {
		return Win
	}

	return Lost
}

func (option Option) getScore() int {
	switch option {
	case Rock:
		return 1
	case Paper:
		return 2
	case Scissors:
		return 3
	}
	return 0
}

func (option Option) next() Option {
	if option > 1 {
		return 0
	}
	return option + 1
}

func (option Option) prev() Option {
	if option < 1 {
		return 2
	}
	return option - 1
}

func (result Result) getScore() int {
	switch result {
	case Lost:
		return 0
	case Draw:
		return 3
	case Win:
		return 6
	}
	return 0
}

func (result Result) choosePlay(opponent Option) Option {
	if result == Draw {
		return opponent
	}

	if result == Win {
		return opponent.next()
	}

	if result == Lost {
		return opponent.prev()
	}

	return 0
}

func TotalScoreWithElfStrategy(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	sum := 0
	for scanner.Scan() {
		var opponent, player rune
		_, err := fmt.Sscanf(scanner.Text(), "%c %c", &opponent, &player)
		if err != nil {
			return 0
		}

		optionOpponent := parsePlayerOpponent(opponent)
		optionPlayer := parsePlayerOpponent(player)

		result := optionPlayer.play(optionOpponent)

		sum += result.getScore() + optionPlayer.getScore()
	}

	return sum
}

func TotalScoreWithElfStrategyWinLose(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	sum := 0
	for scanner.Scan() {
		var opponent, winLoseOrDraw rune
		_, err := fmt.Sscanf(scanner.Text(), "%c %c", &opponent, &winLoseOrDraw)
		if err != nil {
			return 0
		}

		opponentPlay := parsePlayerOpponent(opponent)
		desiredResult := parseResult(winLoseOrDraw)

		playerPlay := desiredResult.choosePlay(opponentPlay)

		sum += desiredResult.getScore() + playerPlay.getScore()
	}

	return sum
}
