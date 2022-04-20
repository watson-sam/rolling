// Package rolling implements rolling sum, mean (avg), count and nunique calculations over a
// given number of float64 values.
package rolling

import (
	"math"
	"time"
)

type valueAndTime struct {
	value    float64
	date     time.Time
	weighted float64
}

type valueAndTimes []valueAndTime

func (vat *valueAndTimes) GetValues() (values []float64) {
	for _, v := range *vat {
		values = append(values, v.value)
	}
	return values
}

func (vat *valueAndTimes) GetWeighted() (values []float64) {
	for _, v := range *vat {
		values = append(values, v.weighted)
	}
	return values
}

// RollingObject - the struct that holds the 'settings' and current values to be used in any
// calculations.
type RollingTimeObject struct {
	windowRange      string
	window           int
	values           valueAndTimes
	ignoreNanValues  bool
	ignoreInfValues  bool
	ignoreZeroValues bool
	weight           float64
	weighted         []float64
}

// SetIgnoreInfValues - controls if we want to ignore non number values when producing the outputs
// of any calculations
func (ro *RollingTimeObject) SetIgnoreNanValues(ignoreNanValues bool) {
	ro.ignoreNanValues = ignoreNanValues
}

// SetIgnoreInfValues - controls if we want to ignore infinite (both positive and negative values)
// when producing the outputs of any calculations
func (ro *RollingTimeObject) SetIgnoreInfValues(ignoreInfValues bool) {
	ro.ignoreInfValues = ignoreInfValues
}

// SetIgnoreInfValues - controls if we want to ignore zero values when producing the outputs of
// any calculations
func (ro *RollingTimeObject) SetIgnoreZeroValues(ignoreZeroValues bool) {
	ro.ignoreZeroValues = ignoreZeroValues
}

func (ro *RollingTimeObject) GetCutoffDate(date time.Time) (cutoffDate time.Time) {
	if ro.windowRange == "day" {
		cutoffDate = date.AddDate(0, 0, -ro.window)
	} else if ro.windowRange == "month" {
		cutoffDate = date.AddDate(0, -ro.window, 0)
	} else if ro.windowRange == "year" {
		cutoffDate = date.AddDate(-ro.window, 0, 0)
	} else {
		panic("windowRange variable not recognized, given value: " + ro.windowRange)
	}
	return cutoffDate
}

func (ro *RollingTimeObject) FilterValues(date time.Time) {
	cutoffDate := ro.GetCutoffDate(date)

	filtered := make(valueAndTimes, len(ro.values))
	k := 0
	for _, vad := range ro.values {
		if vad.date.After(cutoffDate) { // filter
			filtered[k] = vad
			k++
		}
	}
	ro.values = filtered[:k]
}

// Add - if given value meets the given conditions, append to the values used in the calculation,
// adjusting this so it it relevant for the supplied window
func (ro *RollingTimeObject) Add(value float64, date time.Time) {
	if ro.ignoreNanValues && math.IsNaN(value) {
		return
	}
	if ro.ignoreInfValues && (math.IsInf(value, 1) || math.IsInf(value, -1)) {
		return
	}
	if ro.ignoreZeroValues && (value == 0) {
		return
	}

	ro.values = append(ro.values, valueAndTime{value: value, date: date})
}

// Calc - calculate the value of the supplied calculation based from the values stored within the
// rolling object values. Options are:
// - sum: find the total of all the values
// - avg: find the arithmetic mean of the values
// - count: find the number of values
// - nunique: find the number of distinct values
// - std: find the standard deviation of the values
func (ro *RollingTimeObject) Calc(calc string) float64 {
	var values []float64
	if ro.weight > 0 {
		values = ro.weighted
	} else {
		values = ro.values.GetValues()
	}
	if calc == "sum" {
		return Sum(values)
	} else if calc == "avg" {
		return Avg(values)
	} else if calc == "count" {
		return Count(values)
	} else if calc == "max" {
		return Max(values)
	} else if calc == "min" {
		return Min(values)
	} else if calc == "nunique" {
		return NUnique(values)
	} else if calc == "std" {
		return Std(values)
	}
	panic("Supplied `calc` argument is not valid - must be one of: 'sum', 'avg', 'min', 'max', 'count', 'nunique' or 'std', recieved value: " + calc)
}

// NewRollingObject - set up a new rolling object with a supplied window with the default settings
func NewRollingTimeObject(window int, windowRange string) *RollingTimeObject {
	return &RollingTimeObject{
		window:           window,
		windowRange:      windowRange,
		values:           []valueAndTime{},
		ignoreNanValues:  ignoreNanValuesDefault,
		ignoreInfValues:  ignoreInfValuesDefault,
		ignoreZeroValues: ignoreZeroValuesDefault,
	}
}
