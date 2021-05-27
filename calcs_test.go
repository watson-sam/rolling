package rolling

import (
	"math"
	"testing"
)

var calcSamples = []float64{1, 2, 3, 1, 2, 3, 1, 2, 3}

func TestCalcsSum(t *testing.T) {
	result := Sum(calcSamples)
	expected := 18.0
	if result != expected {
		t.Errorf("e.Value() is %v, wanted %v", result, expected)
	}
}

func TestCalcsAvg(t *testing.T) {
	result := Avg(calcSamples)
	expected := 2.0
	if result != expected {
		t.Errorf("e.Value() is %v, wanted %v", result, expected)
	}
}

func TestCalcsNUnique(t *testing.T) {
	result := NUnique(calcSamples)
	expected := 3.0
	if result != expected {
		t.Errorf("e.Value() is %v, wanted %v", result, expected)
	}
}

func TestCalcsStd(t *testing.T) {
	result := Std(calcSamples)
	result = math.Round(result*(10*ACCEPTABLE_ROUNDING_ERROR)) / (10 * ACCEPTABLE_ROUNDING_ERROR)
	expected := 0.8250
	if result != expected {
		t.Errorf("e.Value() is %v, wanted %v", result, expected)
	}
}
