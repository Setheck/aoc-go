package days_test

import (
	"aoc-go/2018/days"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	Day3Input       = "testdata/day3.txt"
	Day3Part1Result = 107043
	Day3Part2Result = 346
)

func Test_Day3Part1(t *testing.T) {
	input := InputLines(t, Day3Input)
	assert.Equal(t, Day3Part1Result, days.OverlappingClaims(input), "inches of overlap did not match accepted answer")
}

func Test_Day3Part2(t *testing.T) {
	input := InputLines(t, Day3Input)
	assert.Equal(t, []int{Day3Part2Result}, days.NoOverlap(input), "Id did not match the accepted answer")
}

func TestSquare_NoOverlap(t *testing.T) {
	s1 := days.NewSquare(0, 1, 1, 1)
	s2 := days.NewSquare(0, 3, 1, 1)
	overlap := s1.Overlap(s2)
	assert.True(t, overlap.Empty())
}

func TestSquare_Overlap(t *testing.T) {
	s1 := days.NewSquare(5, 5, 10, 10)
	s2 := days.NewSquare(6, 6, 2, 2)
	overlap := s1.Overlap(s2)
	assert.False(t, overlap.Empty())
	assert.Equal(t, 4, overlap.Area())
}
