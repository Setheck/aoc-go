package days_test

import (
	"aoc-go/2018/days"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	Day3Input       = "testdata/test.txt" //"testdata/day2.txt"
	Day3Part1Result = 0
	Day3Part2Result = 0
)

func Test_Day3Part1(t *testing.T) {
	input := InputLines(t, Day3Input)
	// claim := days.NewClaim("#9 @ 109,286: 11x16")
	fmt.Println("Result:", days.OverlappingClaims(input))
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
