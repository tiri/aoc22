package aoc22

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const testDataFolder = "testdata"

func goldenTestSplit(data []byte, atEOF bool) (advance int, token []byte, err error) {

	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	const substr = "\n-----\n"
	if i := strings.Index(string(data), substr); i >= 0 {
		return i + len(substr), data[0:i], nil
	}

	if atEOF {
		return len(data), data, nil
	}

	return
}

func GoldenTest(t *testing.T, goldenFile string, f func(io.Reader) string) {
	path := filepath.Join(testDataFolder, goldenFile)
	_, filename := filepath.Split(path)
	testName := filename[:len(filename)-len(filepath.Ext(path))]

	t.Run(testName, func(t *testing.T) {
		reader, err := os.Open(path)
		if err != nil {
			t.Fatal("error reading file:", err)
		}

		scanner := bufio.NewScanner(reader)
		scanner.Split(goldenTestSplit)

		var input, want string
		if scanner.Scan() {
			input = scanner.Text()
		}
		if scanner.Scan() {
			want = scanner.Text()
		}

		got := f(strings.NewReader(input))

		if !(got == want) {
			t.Errorf("\n==== got:\n%s\n==== want:\n%s\n", got, want)
		}
	})
}
