// Package rolling implements rolling sum, mean (avg), count and nunique calculations over a
// given number of float64 values.
package rolling

import (
	"strings"
)

// RollingObject - the struct that holds the 'settings' and current values to be used in any
// calculations.
type RollingStringObject struct {
	window           int
	values           []string
	ignoreZeroValues bool
}

// SetIgnoreInfValues - controls if we want to ignore zero values when producing the outputs of
// any calculations
func (ro *RollingStringObject) SetIgnoreZeroValues(ignoreZeroValues bool) {
	ro.ignoreZeroValues = ignoreZeroValues
}

// Add - if given value meets the given conditions, append to the values used in the calculation,
// adjusting this so it it relevant for the supplied window
func (ro *RollingStringObject) Add(value string) {
	if ro.ignoreZeroValues && (value == "") {
		return
	}

	if len(ro.values) >= ro.window {
		ro.values = ro.values[1:len(ro.values)]
	}
	ro.values = append(ro.values, value)
}

func (ro *RollingStringObject) Values() []string {
	return ro.values
}

func (ro *RollingStringObject) Join(sep string, latestFirst bool) string {
	tnRtn := ro.values
	if latestFirst {
		last := len(tnRtn) - 1
		for i := 0; i < len(tnRtn)/2; i++ {
			tnRtn[i], tnRtn[last-i] = tnRtn[last-i], tnRtn[i]
		}
	}
	return strings.Join(tnRtn, sep)
}

// NewRollingObject - set up a new rolling object with a supplied window with the default settings
func NewRollingStringObject(window int) *RollingStringObject {
	return &RollingStringObject{
		window:           window,
		values:           []string{},
		ignoreZeroValues: ignoreZeroValuesDefault,
	}
}
