package rolling

import "math"

// Sum - find the product of all values
func Sum(values []float64) (result float64) {
	for _, v := range values {
		result += v
	}
	return result
}

// Avg - find the mean of all values
func Avg(values []float64) (result float64) {
	return Sum(values) / float64(len(values))
}

// NUnique - find the number of distinct values
func NUnique(values []float64) float64 {
	dist := []float64{}
	for _, vi := range values {
		contained := false
		for _, vj := range dist {
			if vi == vj {
				contained = true
			}
		}
		if !contained {
			dist = append(dist, vi)
		}
	}
	return float64(len(dist))
}

// Std - find the standard deviation of the values
func Std(values []float64) (result float64) {
	mean := Avg(values)
	for _, v := range values {
		result += math.Pow(v-mean, 2)
	}
	return math.Sqrt(result / float64(len(values)))
}
