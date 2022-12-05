package aoc22

import "strings"

func SplitAtEmptyRow(data []byte, atEOF bool) (advance int, token []byte, err error) {

	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	const substr = "\n\n"
	if i := strings.Index(string(data), substr); i >= 0 {
		return i + len(substr), data[0:i], nil
	}

	if atEOF {
		return len(data), data, nil
	}

	return
}
