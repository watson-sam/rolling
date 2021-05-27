package rolling

import (
	"math"
	"testing"
)

var samples = [100]float64{
	94, 61, 64, 38, 87, 93, 98, 51, 60, 41, 30, 94, 7, 30, 75, 69, 9, 93, 65, 34,
	89, 45, 22, 42, 17, 55, 99, 10, 90, 79, 33, 62, 32, 50, 33, 65, 33, 0, 25, 38,
	13, 65, 94, 51, 61, 40, 56, 100, 25, 33, 25, 3, 24, 63, 10, 71, 40, 36, 23, 49,
	18, 3, 20, 10, 75, 94, 72, 72, 61, 48, 66, 92, 17, 89, 88, 76, 94, 69, 10, 42,
	56, 57, 33, 67, 48, 96, 21, 87, 79, 97, 72, 4, 89, 41, 75, 55, 66, 89, 73, 7,
}

var suspectSamples = [20]float64{
	94, 60, math.NaN(), 38, 65, 93, 97, 51, math.NaN(), 41, 22, 94, 7, math.Inf(0), 75, 0, 0, 0, 62, 34,
}

func TestSum(t *testing.T) {
	ro := NewRollingObject(100)
	for _, f := range samples {
		ro.Add(f)
	}
	result := ro.Calc("sum")
	expected := 5322.0
	if result != expected {
		t.Errorf("e.Value() is %v, wanted %v", result, expected)
	}
}

func TestSum2(t *testing.T) {
	ro := NewRollingObject(10)
	for _, f := range samples {
		ro.Add(f)
	}
	result := ro.Calc("sum")
	expected := 571.0
	if result != expected {
		t.Errorf("e.Value() is %v, wanted %v", result, expected)
	}
}

func TestSum3(t *testing.T) {
	ro := NewRollingObject(20)
	ro.SetIgnoreZeroValues(true)
	for _, f := range suspectSamples {
		ro.Add(f)
	}
	result := ro.Calc("sum")
	expected := 833.0
	if result != expected {
		t.Errorf("e.Value() is %v, wanted %v", result, expected)
	}
}

func TestAvg(t *testing.T) {
	ro := NewRollingObject(100)
	for _, f := range samples {
		ro.Add(f)
	}
	result := ro.Calc("avg")
	expected := 53.22
	if result != expected {
		t.Errorf("e.Value() is %v, wanted %v", result, expected)
	}
}

func TestAvg2(t *testing.T) {
	ro := NewRollingObject(10)
	for _, f := range samples {
		ro.Add(f)
	}
	result := ro.Calc("avg")
	expected := 57.1
	if result != expected {
		t.Errorf("e.Value() is %v, wanted %v", result, expected)
	}
}

func TestAvg3(t *testing.T) {
	ro := NewRollingObject(20)
	ro.SetIgnoreZeroValues(true)
	for _, f := range suspectSamples {
		ro.Add(f)
	}
	result := ro.Calc("avg")
	expected := 59.5
	if result != expected {
		t.Errorf("e.Value() is %v, wanted %v", result, expected)
	}
}

func TestCount(t *testing.T) {
	ro := NewRollingObject(100)
	for _, f := range samples {
		ro.Add(f)
	}
	result := ro.Calc("count")
	expected := 100.0
	if result != expected {
		t.Errorf("e.Value() is %v, wanted %v", result, expected)
	}
}

func TestCount2(t *testing.T) {
	ro := NewRollingObject(10)
	for _, f := range samples {
		ro.Add(f)
	}
	result := ro.Calc("count")
	expected := 10.0
	if result != expected {
		t.Errorf("e.Value() is %v, wanted %v", result, expected)
	}
}

func TestCount3(t *testing.T) {
	ro := NewRollingObject(20)
	ro.SetIgnoreZeroValues(true)
	for _, f := range suspectSamples {
		ro.Add(f)
	}
	result := ro.Calc("count")
	expected := 14.0
	if result != expected {
		t.Errorf("e.Value() is %v, wanted %v", result, expected)
	}
}

func TestNUnique(t *testing.T) {
	ro := NewRollingObject(100)
	for _, f := range samples {
		ro.Add(f)
	}
	result := ro.Calc("nunique")
	expected := 59.0
	if result != expected {
		t.Errorf("e.Value() is %v, wanted %v", result, expected)
	}
}

func TestNUnique2(t *testing.T) {
	ro := NewRollingObject(10)
	for _, f := range samples {
		ro.Add(f)
	}
	result := ro.Calc("nunique")
	expected := 9.0
	if result != expected {
		t.Errorf("e.Value() is %v, wanted %v", result, expected)
	}
}

func TestNUnique3(t *testing.T) {
	ro := NewRollingObject(20)
	ro.SetIgnoreZeroValues(true)
	for _, f := range suspectSamples {
		ro.Add(f)
	}
	result := ro.Calc("nunique")
	expected := 13.0
	if result != expected {
		t.Errorf("e.Value() is %v, wanted %v", result, expected)
	}
}
