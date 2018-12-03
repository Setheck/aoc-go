package days_test

import (
	"aoc-go/2018/days"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	Day2Input       = "testdata/day2.txt"
	Day2Part1Result = 6888
	Day2Part2Result = "icxjvbrobtunlelzpdmfkahgs"
)

// Get the simple list checksum of the input
func Test_Day2Part1(t *testing.T) {
	input := InputLines(t, Day2Input)
	sum := days.ListChecksum(input)
	assert.Equal(t, Day2Part1Result, sum, "Checksum did not match the accepted answer")
}

// Find common letters between the two ids that differ by one letter
func Test_Day2Part2(t *testing.T) {
	input := InputLines(t, Day2Input)
	common := days.CommonInList(input)
	assert.Equal(t, Day2Part2Result, common, "String of characters did not match the accepted answer")
}
