package days_test

import (
	"adventofcode2018/days"
	"github.com/stretchr/testify/assert"
	"testing"
)

const Day1Input = "../testdata/day1.txt"

func Test_Day1Part1(t *testing.T) {
	input := InputLines(t, Day1Input)
	device := days.NewDevice()
	device.CalibrateDevice(input, func(d *days.Device, i int) bool {
		return i == 1
	})
	assert.Equal(t, int64(595), device.Frequency(), "Frequency did not match accepted answer")
}

func Test_Day1Part2(t *testing.T) {
	input := InputLines(t, Day1Input)
	device := days.NewDevice()
	device.CalibrateDevice(input, func(d *days.Device, i int) bool {
		return d.SeenBefore(d.Frequency()) > 1
	})
	assert.Equal(t, int64(80598), device.Frequency(), "Frequency did not match accepted answer")
}
