package days

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

type Device struct {
	frequency int64
	seenFrequencies map[int64]int64
}

// NewDevice create a new device to calibrate
func NewDevice() *Device {
	device := &Device{
		frequency:0,
		seenFrequencies:make(map[int64]int64)}
	device.Seen(0)
	return device
}

// Seen mark an int as seen by adding it to the internal histogram seenFrequencies
// if it has already been seen, then we increment the count in the histogram
func (d *Device) Seen(in int64) {
	if val,ok := d.seenFrequencies[in]; ok {
		d.seenFrequencies[in] = val + 1
	} else {
		d.seenFrequencies[in] = 1
	}
}

// SeenCount return the number of different frequencies the device has seen thus far
func (d *Device) SeenCount() int {
	return len(d.seenFrequencies)
}

// SeenBefore retrieve the count of how many times the device has seen the frequency `in` before
func (d *Device) SeenBefore(in int64) int64 {
	if val,ok := d.seenFrequencies[in]; ok {
		return val
	}
	return 0
}

// Calculate update the frequency of the device with the input frequency instruction
// this function all marks frequencies seen as the internal frequency is updated.
// ex: +1 will add 1 to the frequency, -1 will subtract 1 from the frequency
func (d *Device) Calculate(s string) error {
	i,err := strconv.ParseInt(s,10,0)
	if err != nil {
		return fmt.Errorf("failed to parse %s", s)
	}
	d.frequency += i
	d.Seen(d.frequency)
	return nil
}

// Frequency retrieve the current device frequency
func (d *Device) Frequency() int64 {
	return d.frequency
}

// CalibrateDevice read each line of input and Calculate the frequency of the device.
// the second parameter is a function that describes when a device condition is met and the calibration is complete
// inputs to the condition function are the device in the current state and the number iterations over the input
// that have occurred.
func (d *Device) CalibrateDevice(input []string, condition func(d *Device, iterations int) bool) *Device {
	iterations := 0
	for run := true; run; {
		for _,line := range input {
			if condition(d, iterations) {
				return d
			}
			if err := d.Calculate(line); err != nil {
				log.Fatal(err)
			}
		}
		iterations++
		if iterations >= math.MaxInt64 {
			run = false
		}
	}
	return d
}