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

	reader, err := os.Open(path)
	if err != nil {
		t.Fatal("error reading file:", err)
	}

	t.Run(testName, func(t *testing.T) {
		got := f(reader)

		if got != want {
			t.Errorf("\n==== got:\n%v\n==== want:\n%v\n", got, want)
		}
	})
}

func RunBenchmark[T comparable](b *testing.B, testFile string, f func(io.Reader) T) {
	path := filepath.Join(testFile)
	_, filename := filepath.Split(path)
	testName := filename[:len(filename)-len(filepath.Ext(path))]

	reader, err := os.Open(path)
	if err != nil {
		b.Fatal("error reading file:", err)
	}

	b.Run(testName, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			f(reader)
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
