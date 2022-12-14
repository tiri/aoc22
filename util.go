package aoc22

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func RunTest[T comparable](t *testing.T, testFile string, f func(io.Reader) T, want T) {
	path := filepath.Join(testFile)
	_, filename := filepath.Split(path)
	testName := filename[:len(filename)-len(filepath.Ext(path))]

	t.Run(testName, func(t *testing.T) {
		reader, err := os.Open(path)
		if err != nil {
			t.Fatal("error reading file:", err)
		}

		got := f(reader)

		if got != want {
			t.Errorf("\n==== got:\n%v\n==== want:\n%v\n", got, want)
		}
	})
}

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
