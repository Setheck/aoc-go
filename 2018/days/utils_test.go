package days_test

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"testing"
)

// InputLines parse a file into an array of lines from the file
func InputLines(t *testing.T, filename string) []string {
	t.Helper()
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Error(err)
	}
	scanner := bufio.NewScanner(bytes.NewBuffer(b))
	output := make([]string, 0)
	for scanner.Scan() {
		output = append(output, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		t.Error(err)
	}
	return output
}
