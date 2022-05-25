// Package rolling implements rolling sum, mean (avg), count and nunique calculations over a
// given number of float64 values.
package rolling

import "math"

const (
	// By default nan values are ignored
	ignoreNanValuesDefault bool = true
	// By default infinite values (both positive and negative) are ignored
	ignoreInfValuesDefault bool = true
	// By default zero values are treated as a per any other number (ie not ignored)
	ignoreZeroValuesDefault bool = false
)

// RollingObject - the struct that holds the 'settings' and current values to be used in any
// calculations.
type RollingObject struct {
	window           int
	values           []float64
	ignoreNanValues  bool
	ignoreInfValues  bool
	ignoreZeroValues bool
}

// SetIgnoreInfValues - controls if we want to ignore non number values when producing the outputs
// of any calculations
func (ro *RollingObject) SetIgnoreNanValues(ignoreNanValues bool) {
	ro.ignoreNanValues = ignoreNanValues
}

// SetIgnoreInfValues - controls if we want to ignore infinites (both positive and negative values)
// when producing the outputs of any calculations
func (ro *RollingObject) SetIgnoreInfValues(ignoreInfValues bool) {
	ro.ignoreInfValues = ignoreInfValues
}

// SetIgnoreInfValues - controls if we want to ignore zero values when producing the outputs of
// any calculations
func (ro *RollingObject) SetIgnoreZeroValues(ignoreZeroValues bool) {
	ro.ignoreZeroValues = ignoreZeroValues
}

// Add - if given value meets the given conditions, append to the values used in the calculation,
// adjusting this so it it relevant for the supplied window
func (ro *RollingObject) Add(value float64) {
	if ro.ignoreNanValues && math.IsNaN(value) {
		return
	}
	if ro.ignoreInfValues && (math.IsInf(value, 1) || math.IsInf(value, -1)) {
		return
	}
	if ro.ignoreZeroValues && (value == 0) {
		return
	}

	if len(ro.values) >= ro.window {
		ro.values = ro.values[1:len(ro.values)]
	}
	ro.values = append(ro.values, value)
}

// Calc - calculate the value of the supplied calculation based from the values stored within the
// rolling object values. Options are:
// - sum: find the total of all the values
// - avg: find the arithmetic mean of the values
// - count: find the number of values
// - nunique: find the number of distinct values
// - std: find the standard deviation of the values
func (ro *RollingObject) Calc(calc string) float64 {
	if calc == "sum" {
		return Sum(ro.values)
	} else if calc == "avg" {
		return Avg(ro.values)
	} else if calc == "count" {
		return Count(ro.values)
	} else if calc == "max" {
		return Max(ro.values)
	} else if calc == "min" {
		return Min(ro.values)
	} else if calc == "nunique" {
		return NUnique(ro.values)
	} else if calc == "std" {
		return Std(ro.values)
	} else if calc == "var" {
		return Var(ro.values)
	}
	panic("Supplied `calc` argument is not valid - must be one of: 'sum', 'avg', 'min', 'max', 'count', 'nunique', 'std' or 'var' received value: " + calc)
}

func (ro *RollingObject) Values() []float64 {
	return ro.values
}

// NewRollingObject - set up a new rolling object with a supplied window with the default settings
func NewRollingObject(window int) *RollingObject {
	return &RollingObject{
		window:           window,
		values:           []float64{},
		ignoreNanValues:  ignoreNanValuesDefault,
		ignoreInfValues:  ignoreInfValuesDefault,
		ignoreZeroValues: ignoreZeroValuesDefault,
	}
}
