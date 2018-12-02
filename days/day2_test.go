package days_test

import (
	"adventofcode2018/days"
	"github.com/stretchr/testify/assert"
	"testing"
)

const Day2Input = "../testdata/day2.txt"

func Test_Day2Part1(t *testing.T) {
	input := InputLines(t, Day2Input)
	sum := days.ListChecksum(input)
	assert.Equal(t, 6888, sum)
}

func Test_Day2Part2(t *testing.T) {
	input := InputLines(t, Day2Input)
	common := days.CommonInList(input)
	assert.Equal(t, "icxjvbrobtunlelzpdmfkahgs", common)
}