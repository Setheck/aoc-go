package days_test

import (
	"aoc-go/2018/days"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	Day4Input       = "testdata/day4.txt"
	Day4Part1Result = 0
	Day4Part2Result = 0
)

func Test_Day4Part1(t *testing.T) {
	t.SkipNow()
	input := InputLines(t, Day4Input)
	events := days.ParseEvents(input)
	id := days.MostAsleep(events)
	assert.Equal(t, id, Day4Part1Result)
}
