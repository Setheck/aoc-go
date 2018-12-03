package days_test

import (
	"aoc-go/2018/days"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	Day1Input       = "testdata/day1.txt"
	Day1Part1Result = 595
	Day1Part2Result = 80598
)

// Get the frequency for completing a single iteration of the input
func Test_Day1Part1(t *testing.T) {
	input := InputLines(t, Day1Input)
	device := days.NewDevice()
	device.CalibrateDevice(input, func(d *days.Device, i int) bool {
		return i == 1
	})
	assert.Equal(t, int64(Day1Part1Result), device.Frequency(), "Frequency did not match accepted answer")
}

// Get the first frequency that appears twice, after continually iterating over the input
func Test_Day1Part2(t *testing.T) {
	input := InputLines(t, Day1Input)
	device := days.NewDevice()
	device.CalibrateDevice(input, func(d *days.Device, i int) bool {
		return d.SeenBefore(d.Frequency()) > 1
	})
	assert.Equal(t, int64(Day1Part2Result), device.Frequency(), "Frequency did not match accepted answer")
}
